package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// ContactRequest mirrors the JSON body sent by the frontend.
type ContactRequest struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	Message           string `json:"message"`
	Website           string `json:"website"`           // Honeypot field
	TurnstileResponse string `json:"cf-turnstile-response"` // Cloudflare Turnstile token
}

// jsonResponse writes a JSON body with the given status code.
func jsonResponse(w http.ResponseWriter, status int, key, value string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, `{%q:%q}`, key, value)
}

// corsMiddleware validates the Origin header and injects CORS headers.
func corsMiddleware(allowedOrigin string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if allowedOrigin != "" && origin != allowedOrigin {
			// Strict check when env is set
			jsonResponse(w, http.StatusForbidden, "error", "forbidden")
			return
		}

		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		}
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "86400")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

// verifyTurnstile calls Cloudflare's siteverify endpoint.
// Returns (success bool, err error). Skips if TURNSTILE_SECRET is unset.
func verifyTurnstile(token string) (bool, error) {
	secret := os.Getenv("TURNSTILE_SECRET")
	if secret == "" {
		// Local dev — skip Turnstile verification
		log.Println("[turnstile] TURNSTILE_SECRET not set — skipping verification")
		return true, nil
	}
	if token == "" {
		return false, nil
	}

	body := strings.NewReader(fmt.Sprintf(
		"secret=%s&response=%s", secret, token,
	))
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Post(
		"https://challenges.cloudflare.com/turnstile/v0/siteverify",
		"application/x-www-form-urlencoded",
		body,
	)
	if err != nil {
		return false, fmt.Errorf("turnstile request failed: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Success bool     `json:"success"`
		Errors  []string `json:"error-codes"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, fmt.Errorf("turnstile decode failed: %w", err)
	}
	return result.Success, nil
}

// sendEmail sends an email via the Postmark API.
func sendEmail(from, to, subject, textBody, htmlBody string) error {
	token := os.Getenv("POSTMARK_TOKEN")
	if token == "" {
		log.Printf("[email] POSTMARK_TOKEN not set — would have sent: To=%s Subject=%s", to, subject)
		return nil
	}

	payload := map[string]string{
		"From":     from,
		"To":       to,
		"Subject":  subject,
		"TextBody": textBody,
		"HtmlBody": htmlBody,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.postmarkapp.com/email", bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Postmark-Server-Token", token)

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("postmark request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("postmark error %d: %s", resp.StatusCode, string(body))
	}
	return nil
}

// handleContact processes POST /api/contact.
func handleContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusMethodNotAllowed, "error", "method not allowed")
		return
	}

	// Parse body — limit to 64 KB to prevent abuse
	r.Body = http.MaxBytesReader(w, r.Body, 64*1024)
	var req ContactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonResponse(w, http.StatusBadRequest, "error", "invalid JSON body")
		return
	}

	// Honeypot — bot filled the hidden field; return fake success
	if strings.TrimSpace(req.Website) != "" {
		log.Printf("[honeypot] blocked submission from %s", r.RemoteAddr)
		jsonResponse(w, http.StatusOK, "message", "Thank you! We'll be in touch shortly.")
		return
	}

	// Validate required fields
	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)
	req.Message = strings.TrimSpace(req.Message)

	if req.Name == "" {
		jsonResponse(w, http.StatusBadRequest, "error", "Name is required.")
		return
	}
	if req.Email == "" || !strings.Contains(req.Email, "@") {
		jsonResponse(w, http.StatusBadRequest, "error", "A valid email address is required.")
		return
	}
	if len(req.Message) < 40 {
		jsonResponse(w, http.StatusBadRequest, "error", "Message must be at least 40 characters.")
		return
	}
	if len(req.Message) > 500 {
		jsonResponse(w, http.StatusBadRequest, "error", "Message must be 500 characters or fewer.")
		return
	}

	// Turnstile verification
	ok, err := verifyTurnstile(req.TurnstileResponse)
	if err != nil {
		log.Printf("[turnstile] error: %v", err)
		// On Turnstile service error, allow through (degrade gracefully)
	} else if !ok {
		jsonResponse(w, http.StatusBadRequest, "error", "CAPTCHA verification failed. Please try again.")
		return
	}

	// Build and send email
	fromEmail := os.Getenv("FROM_EMAIL")
	toEmail := os.Getenv("TO_EMAIL")
	if fromEmail == "" {
		fromEmail = "noreply@fireflysoftware.dev"
	}
	if toEmail == "" {
		toEmail = "service@fireflysoftware.dev"
	}

	subject := fmt.Sprintf("New contact from %s via fireflysoftware.dev", req.Name)

	textBody := fmt.Sprintf(
		"New contact form submission\n\nName: %s\nEmail: %s\nPhone: %s\n\nMessage:\n%s",
		req.Name, req.Email, req.Phone, req.Message,
	)

	htmlBody := fmt.Sprintf(`<h2>New contact form submission</h2>
<table>
  <tr><th style="text-align:left;padding-right:16px;">Name</th><td>%s</td></tr>
  <tr><th style="text-align:left;padding-right:16px;">Email</th><td><a href="mailto:%s">%s</a></td></tr>
  <tr><th style="text-align:left;padding-right:16px;">Phone</th><td>%s</td></tr>
</table>
<h3>Message</h3>
<p style="white-space:pre-wrap;">%s</p>`,
		req.Name, req.Email, req.Email, req.Phone, req.Message,
	)

	if err := sendEmail(fromEmail, toEmail, subject, textBody, htmlBody); err != nil {
		log.Printf("[email] send failed: %v", err)
		jsonResponse(w, http.StatusInternalServerError, "error", "Failed to send message. Please call us directly.")
		return
	}

	log.Printf("[contact] sent email from %s <%s>", req.Name, req.Email)
	jsonResponse(w, http.StatusOK, "message", "Thank you! Your message has been sent. We'll be in touch shortly.")
}

// handleHealth responds with 200 OK for readiness/liveness probes.
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"status":"ok"}`)
}

func main() {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}
	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
	if allowedOrigin == "" {
		allowedOrigin = "http://localhost:1313"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/contact", corsMiddleware(allowedOrigin, handleContact))
	mux.HandleFunc("/api/health", handleHealth)

	addr := "127.0.0.1:" + port
	log.Printf("[api] listening on %s (ALLOWED_ORIGIN=%s)", addr, allowedOrigin)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("[api] fatal: %v", err)
	}
}
