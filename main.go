package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"

	"fireflysoftware.dev/website/handlers"
)

func main() {
	// Initialize Sentry (Bugsink-compatible via DSN)
	sentryDSN := os.Getenv("SENTRY_DSN")
	if sentryDSN != "" {
		env := os.Getenv("SENTRY_ENVIRONMENT")
		if env == "" {
			env = "production"
		}
		err := sentry.Init(sentry.ClientOptions{
			Dsn:              sentryDSN,
			Environment:      env,
			EnableTracing:    false,
			TracesSampleRate: 0,
		})
		if err != nil {
			log.Printf("[sentry] init failed: %v", err)
		} else {
			log.Printf("[sentry] initialized (env=%s)", env)
		}
		defer sentry.Flush(2 * time.Second)
	} else {
		log.Println("[sentry] SENTRY_DSN not set — error reporting disabled")
	}

	// Configuration
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}
	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
	if allowedOrigin == "" {
		allowedOrigin = "http://localhost:8080"
	}
	devMode := os.Getenv("DEV_MODE") == "true"

	// Template renderer
	tr, err := handlers.NewTemplateRenderer("templates", devMode)
	if err != nil {
		log.Fatalf("[server] template parse error: %v", err)
	}

	// Site data
	site := handlers.NewSiteData()

	// Blog post handler
	postHandler := handlers.NewPostHandler("content/posts", tr, site)

	// Router
	mux := http.NewServeMux()

	// Static files
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Page routes
	mux.HandleFunc("GET /{$}", handlers.PageHandler(tr, "home", handlers.NewPageData(
		"Firefly Software — Custom Websites & Software for Small Business",
		"Firefly Software is a small custom development studio based in Helena, Montana. We build websites and software for small businesses.",
		site,
	)))
	mux.HandleFunc("GET /work", handlers.PageHandler(tr, "work", handlers.NewPageData(
		"Work — Firefly Software",
		"A selection of recent projects. Websites, software, and the problems behind both.",
		site,
	).WithNav("work")))
	mux.HandleFunc("GET /process", handlers.PageHandler(tr, "process", handlers.NewPageData(
		"How We Work — Firefly Software",
		"Every project is different. The process isn't. Here's exactly what working with Firefly looks like.",
		site,
	).WithNav("process")))
	mux.HandleFunc("GET /about", handlers.PageHandler(tr, "about", handlers.NewPageData(
		"About — Firefly Software",
		"A small, focused team building custom websites and software for small businesses. Based in Helena, Montana.",
		site,
	).WithNav("about")))
	mux.HandleFunc("GET /pricing", handlers.PageHandler(tr, "pricing", handlers.NewPageData(
		"Pricing — Firefly Software",
		"One relationship, one monthly number. Custom websites built, hosted, and maintained under one agreement.",
		site,
	).WithNav("pricing")))
	mux.HandleFunc("GET /contact", handlers.PageHandler(tr, "contact", handlers.NewPageData(
		"Contact — Firefly Software",
		"Tell us what you're working on. We'll tell you honestly whether we're the right fit.",
		site,
	).WithNav("contact")))
	mux.HandleFunc("GET /privacy", handlers.PageHandler(tr, "privacy", handlers.NewPageData(
		"Privacy Policy — Firefly Software",
		"How Firefly Software handles your data. We use Google Analytics and Plausible Analytics.",
		site,
	)))
	mux.HandleFunc("GET /terms", handlers.PageHandler(tr, "terms", handlers.NewPageData(
		"Terms of Service — Firefly Software",
		"Terms governing Firefly Software's website development and maintenance services.",
		site,
	)))
	mux.HandleFunc("GET /services/websites", handlers.PageHandler(tr, "services-websites", handlers.NewPageData(
		"Websites — Firefly Software",
		"Custom-built from scratch. No templates, no page builders, no WordPress.",
		site,
	)))
	mux.HandleFunc("GET /services/software", handlers.PageHandler(tr, "services-software", handlers.NewPageData(
		"Custom Software — Firefly Software",
		"If you're doing it in a spreadsheet, there's a better way.",
		site,
	)))

	// Blog routes
	mux.HandleFunc("GET /posts/", postHandler.HandlePostList)
	mux.HandleFunc("GET /posts/{slug}", postHandler.HandlePost)

	// API routes
	mux.HandleFunc("/api/contact", handlers.CORSMiddleware(allowedOrigin, handlers.HandleContact))
	mux.HandleFunc("GET /api/health", handlers.HandleHealth)

	// 301 Redirects from old routes
	mux.HandleFunc("GET /website-maintenance/", handlers.Redirect301("/services/websites"))
	mux.HandleFunc("GET /website-maintenance", handlers.Redirect301("/services/websites"))
	mux.HandleFunc("GET /site-rescue/", handlers.Redirect301("/services/websites"))
	mux.HandleFunc("GET /site-rescue", handlers.Redirect301("/services/websites"))
	mux.HandleFunc("GET /outfitters/", handlers.Redirect301("/work"))
	mux.HandleFunc("GET /outfitters", handlers.Redirect301("/work"))
	mux.HandleFunc("GET /pricing/", handlers.Redirect301("/pricing"))
	mux.HandleFunc("GET /portfolio/", handlers.Redirect301("/work"))
	mux.HandleFunc("GET /portfolio", handlers.Redirect301("/work"))
	mux.HandleFunc("GET /contact-us/", handlers.Redirect301("/contact"))
	mux.HandleFunc("GET /contact-us", handlers.Redirect301("/contact"))
	mux.HandleFunc("GET /website-development/", handlers.Redirect301("/services/websites"))
	mux.HandleFunc("GET /website-development", handlers.Redirect301("/services/websites"))
	mux.HandleFunc("GET /custom-software/", handlers.Redirect301("/services/software"))
	mux.HandleFunc("GET /custom-software", handlers.Redirect301("/services/software"))
	mux.HandleFunc("GET /helena-web-developer/", handlers.Redirect301("/"))
	mux.HandleFunc("GET /helena-web-developer", handlers.Redirect301("/"))
	mux.HandleFunc("GET /montana-web-design/", handlers.Redirect301("/"))
	mux.HandleFunc("GET /montana-web-design", handlers.Redirect301("/"))

	// Middleware stack: logging → recovery → mux
	var handler http.Handler = mux
	handler = handlers.RecoveryMiddleware(handler)
	handler = handlers.LoggingMiddleware(handler)

	addr := "0.0.0.0:" + port
	log.Printf("[server] listening on %s (dev=%v)", addr, devMode)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("[server] fatal: %v", err)
	}
}
