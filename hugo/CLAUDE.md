# CLAUDE.md — Firefly Software Hugo Site

This file provides guidance when working with the Hugo-based Firefly Software website.

## Project Overview

Firefly Software company website converted from SvelteKit to Hugo + Go API + Docker + Caddy architecture.

**Architecture:**
```
Internet → Outer Caddy (HTTPS/TLS on host) → Docker Container (HTTP on configurable port)
                                                  ├── Inner Caddy (/api/* reverse proxy + static files)
                                                  └── Go API (contact form on localhost:8080)
```

## Local Development Commands

```bash
# Hugo dev server (runs from /workspaces/fs-website/hugo/)
cd hugo
hugo server -D                   # Dev server with drafts at http://localhost:1313

# Go API (in a separate terminal)
cd hugo/api
go run main.go                   # Starts API on localhost:8080

# Full stack via Docker Compose
cd hugo
docker compose up --build        # Build and run container on port 3000

# Build Hugo for production
cd hugo
hugo --gc --minify               # Output goes to hugo/public/
```

## Directory Structure

```
hugo/
├── hugo.toml                    # All site config and business data
├── assets/css/main.css          # Single CSS file — no frameworks
├── content/                     # Page content (minimal markdown)
│   ├── _index.md                # Homepage
│   ├── pricing.md               # Pricing page
│   ├── contact-us.md            # Contact form page
│   ├── success.md               # Post-submission success page
│   ├── portfolio/               # Portfolio case studies
│   ├── posts/                   # Blog posts
│   ├── website-development/     # Service page (section)
│   ├── website-maintenance/     # Service page (section)
│   └── e-commerce/              # Service page (section)
├── data/
│   ├── portfolio.yaml           # Portfolio project data (drives templates)
│   └── pricing.yaml             # Pricing plan data (drives templates)
├── layouts/
│   ├── index.html               # Homepage template
│   ├── _default/baseof.html     # Base template (head, header, footer, scripts)
│   ├── _default/single.html     # Default single page
│   ├── _default/list.html       # Default list page
│   ├── page/                    # Special page layouts
│   │   ├── contact.html         # Contact form with Turnstile
│   │   ├── pricing.html         # Pricing with tabs
│   │   └── success.html         # Success page
│   ├── portfolio/               # Portfolio layouts
│   │   ├── list.html            # Portfolio listing
│   │   └── single.html          # Case study detail
│   ├── posts/                   # Blog layouts
│   │   ├── list.html            # Blog listing
│   │   └── single.html          # Blog post
│   ├── website-development/     # Service page layouts (list.html = single.html)
│   ├── website-maintenance/
│   ├── e-commerce/
│   └── partials/
│       ├── header.html          # Sticky nav with dropdown
│       ├── footer.html          # Dark footer with contact info
│       └── schema.html          # JSON-LD structured data
├── static/                      # Static assets (served as-is)
│   ├── favicon.ico
│   ├── android-chrome-192x192.png
│   ├── apple-touch-icon.png
│   ├── robots.txt
│   └── images/
│       ├── logo/                # Firefly Software logos
│       ├── portfolio/           # Case study images
│       └── [site images]
├── api/
│   ├── go.mod
│   └── main.go                  # Contact form handler (zero external dependencies)
├── Caddyfile                    # Inner Caddy config
├── Dockerfile                   # Three-stage: Hugo + Go + Caddy
├── docker-compose.yml
└── docker-entrypoint.sh
```

## Key Conventions

**Business data lives in hugo.toml** — Never hardcode phone numbers, addresses, emails, or hours in templates. Always use `.Site.Params.*`.

**Portfolio data** — `data/portfolio.yaml` drives the listing page and each case study's metadata (logo, splash image, client, date, location). Content goes in `content/portfolio/[slug].md`.

**Pricing data** — `data/pricing.yaml` has both `subscribe` and `buyout` arrays, used by the homepage and pricing page.

**Service pages** — These are `_index.md` (section pages). Their layout is defined in `layouts/[section]/list.html`. The `list.html` and `single.html` are identical since there's only one page per service section.

**CSS** — Single file at `assets/css/main.css`, minified via Hugo Pipes. Uses CSS custom properties for all colors, spacing, and typography. No frameworks.

**JavaScript** — Inline only, in `{{ define "scripts" }}` blocks. Zero external JS dependencies.

## Cloudflare Turnstile

For local development, use Cloudflare's official test keys:
- **Site key:** `1x00000000000000000000AA` (always passes)
- **Secret:** `1x0000000000000000000000000000000AA` (always passes)

For production, replace `turnstileSiteKey` in `hugo.toml` and set `TURNSTILE_SECRET` env var.

## Environment Variables

| Variable | Default | Description |
|---|---|---|
| `PORT` | `80` | Port Caddy listens on inside the container |
| `API_PORT` | `8080` | Port the Go API listens on (localhost only) |
| `ALLOWED_ORIGIN` | `https://fireflysoftware.dev` | CORS allowed origin |
| `TURNSTILE_SECRET` | _(empty — skip verification)_ | Cloudflare Turnstile secret key |
| `POSTMARK_TOKEN` | _(empty — log only)_ | Postmark API token for email delivery |
| `FROM_EMAIL` | `noreply@fireflysoftware.dev` | Email "From" address |
| `TO_EMAIL` | `service@fireflysoftware.dev` | Email recipient |
| `LISTEN_PORT` | `3000` | Host port (docker-compose.yml) |

## Deployment

### GitHub Actions (automated)
Push to `master` or `main` triggers the workflow in `.github/workflows/deploy.yml`.

Required GitHub Secrets:
- `DOCKERHUB_USERNAME`
- `DOCKERHUB_TOKEN`
- `VPS_HOST`
- `VPS_USER`
- `VPS_SSH_KEY`

### Manual Deploy
```bash
# Build and push image
cd hugo
docker build -t dukerupert/fs-website:latest .
docker push dukerupert/fs-website:latest

# On VPS
cd /opt/fs-website
docker compose pull
docker compose up -d
docker image prune -f
```

### VPS Caddy (outer — handles HTTPS)
```
fireflysoftware.dev {
    reverse_proxy localhost:3000
}
```

## Adding Content

### New Blog Post
```bash
cd hugo
hugo new posts/my-post-title.md
```
Edit the generated file at `content/posts/my-post-title.md`.

### New Portfolio Case Study
1. Add project entry to `data/portfolio.yaml`
2. Copy images to `static/images/portfolio/`
3. Create `content/portfolio/[slug].md` with testimonial front matter and body content

## Color Reference

| Custom Property | Value | Use |
|---|---|---|
| `--color-primary` | `#1e91d9` | Curious Blue — links, buttons, accents |
| `--color-primary-dark` | `#1072b9` | Hover state for primary |
| `--color-secondary` | `#091740` | Downriver — headings, dark backgrounds |
| `--color-tertiary` | `#00cdb0` | Bright Turquoise — highlights, CTA accents |
