package handlers

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

	"github.com/getsentry/sentry-go"
)

// ContactRequest mirrors the JSON body sent by the frontend.
type ContactRequest struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	Service           string `json:"service"`
	Message           string `json:"message"`
	Website           string `json:"website"`               // Honeypot field
	TurnstileResponse string `json:"cf-turnstile-response"` // Cloudflare Turnstile token
}

// captureError sends an error to Sentry if it has been initialized.
func captureError(err error) {
	sentry.CaptureException(err)
}

// jsonResponse writes a JSON body with the given status code.
func jsonResponse(w http.ResponseWriter, status int, key, value string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, `{%q:%q}`, key, value)
}

// verifyTurnstile calls Cloudflare's siteverify endpoint.
// Returns (success bool, err error). Skips if TURNSTILE_SECRET is unset.
func verifyTurnstile(token string) (bool, error) {
	secret := os.Getenv("TURNSTILE_SECRET")
	if secret == "" {
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
		wrappedErr := fmt.Errorf("turnstile request failed: %w", err)
		captureError(wrappedErr)
		return false, wrappedErr
	}
	defer resp.Body.Close()

	var result struct {
		Success bool     `json:"success"`
		Errors  []string `json:"error-codes"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		wrappedErr := fmt.Errorf("turnstile decode failed: %w", err)
		captureError(wrappedErr)
		return false, wrappedErr
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
		wrappedErr := fmt.Errorf("postmark request: %w", err)
		captureError(wrappedErr)
		return wrappedErr
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		wrappedErr := fmt.Errorf("postmark error %d: %s", resp.StatusCode, string(body))
		captureError(wrappedErr)
		return wrappedErr
	}
	return nil
}

// HandleContact processes POST /api/contact.
func HandleContact(w http.ResponseWriter, r *http.Request) {
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
		toEmail = "logan@fireflysoftware.dev"
	}

	req.Service = strings.TrimSpace(req.Service)

	subject := fmt.Sprintf("New contact from %s via fireflysoftware.dev", req.Name)
	if req.Service != "" {
		subject = fmt.Sprintf("[%s] New contact from %s via fireflysoftware.dev", req.Service, req.Name)
	}

	serviceLine := ""
	if req.Service != "" {
		serviceLine = fmt.Sprintf("Service: %s\n", req.Service)
	}
	textBody := fmt.Sprintf(
		"New contact form submission\n\nName: %s\nEmail: %s\nPhone: %s\n%s\nMessage:\n%s",
		req.Name, req.Email, req.Phone, serviceLine, req.Message,
	)

	serviceRow := ""
	if req.Service != "" {
		serviceRow = fmt.Sprintf(`<tr><th style="text-align:left;padding-right:16px;">Service</th><td>%s</td></tr>`, req.Service)
	}
	htmlBody := fmt.Sprintf(`<h2>New contact form submission</h2>
<table>
  <tr><th style="text-align:left;padding-right:16px;">Name</th><td>%s</td></tr>
  <tr><th style="text-align:left;padding-right:16px;">Email</th><td><a href="mailto:%s">%s</a></td></tr>
  <tr><th style="text-align:left;padding-right:16px;">Phone</th><td>%s</td></tr>
  %s
</table>
<h3>Message</h3>
<p style="white-space:pre-wrap;">%s</p>`,
		req.Name, req.Email, req.Email, req.Phone, serviceRow, req.Message,
	)

	if err := sendEmail(fromEmail, toEmail, subject, textBody, htmlBody); err != nil {
		log.Printf("[email] send failed: %v", err)
		jsonResponse(w, http.StatusInternalServerError, "error", "Failed to send message. Please call us directly.")
		return
	}

	log.Printf("[contact] sent email from %s <%s>", req.Name, req.Email)
	jsonResponse(w, http.StatusOK, "message", "Thank you! Your message has been sent. We'll be in touch shortly.")
}

// HandleHealth responds with 200 OK for readiness/liveness probes.
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"status":"ok"}`)
}
