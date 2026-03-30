# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What This Is

Firefly Software company website — a Go server-rendered site with Tailwind CSS v4 and Alpine.js. Deployed as a Docker image with Caddy reverse-proxying to the Go binary.

## Commands

```bash
make dev          # Tailwind watcher + Go server (hot reload, DEV_MODE=true)
make dev-server   # Go server only (no CSS watcher)
make dev-css      # Tailwind watcher only
make build        # Production binary + minified CSS
make clean        # Remove build artifacts
```

There are no tests in this project.

## Architecture

**Go server** (`main.go` + `handlers/`): Standard `net/http` with Go 1.22+ method routing (`"GET /path"`). No framework. Middleware chain: LoggingMiddleware → RecoveryMiddleware → mux.

**Template rendering** (`handlers/pages.go`): Each page template is parsed with shared partials (`templates/base.html`, `templates/nav.html`, `templates/footer.html`). All templates execute the `"base"` block. In dev mode, templates reparse on every request. The template map in `TemplateRenderer.parseAll()` must be updated when adding a new page.

**Adding a new page** requires three changes:
1. Create `templates/{name}.html` defining `{{define "content"}}...{{end}}`
2. Register the template name in `handlers/pages.go` → `parseAll()` pages map
3. Add the route in `main.go` with `PageHandler(tr, "name", NewPageData(...))`

**Blog posts** (`handlers/posts.go`): Markdown files in `content/posts/` with YAML front matter (`title`, `description`, `date`, `author`, `tags`). Rendered via goldmark. Served at `/posts/{slug}` where slug = filename without `.md`.

**Contact form** (`handlers/contact.go`): POST `/api/contact` → Cloudflare Turnstile verification → Postmark email. Honeypot field (`website`). Gracefully degrades when `TURNSTILE_SECRET` or `POSTMARK_TOKEN` are unset.

**CSS** (`static/css/input.css` + `static/css/tokens.css`): Tailwind v4 with `@import "tailwindcss"` and `@source "../../templates"` for class scanning. Design tokens (colors, fonts, spacing) are CSS custom properties in `tokens.css`. Component classes (`.btn-primary`, `.card-dark`, `.section`, `.hero`, etc.) are also in `tokens.css`, not Tailwind utilities.

**Frontend interactivity**: Alpine.js (vendored at `static/js/alpine.min.js`). Used for nav dropdown/hamburger, scroll reveal (`.reveal` → `.is-visible` via IntersectionObserver in `base.html`), and contact form state.

## Deployment

Multi-stage Dockerfile: Node (Tailwind build) → Go (binary build) → Caddy (final image). Caddy handles compression, security headers, caching, and reverse-proxies to the Go server on `API_PORT`. TLS is terminated by an outer Caddy on the host. Error reporting via Sentry SDK pointed at a Bugsink instance (tracing disabled).

## Environment Variables

- `DEV_MODE` — enables template hot reload (dev only)
- `API_PORT` — Go server port (default `8080`)
- `ALLOWED_ORIGIN` — CORS origin for `/api/contact`
- `TURNSTILE_SECRET`, `POSTMARK_TOKEN` — contact form integrations
- `SENTRY_DSN`, `SENTRY_ENVIRONMENT` — error reporting
- `FROM_EMAIL`, `TO_EMAIL` — contact form email addresses
