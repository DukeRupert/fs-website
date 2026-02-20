# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Firefly Software company website built with Hugo + Go API + Docker + Caddy. All site code lives in the `hugo/` directory. See `hugo/CLAUDE.md` for comprehensive documentation.

**Architecture:**
```
Internet → Outer Caddy (HTTPS/TLS on host) → Docker Container (HTTP on configurable port)
                                                  ├── Inner Caddy (/api/* reverse proxy + static files)
                                                  └── Go API (contact form on localhost:8080)
```

## Quick Commands

```bash
# Hugo dev server
cd hugo && hugo server -D             # http://localhost:1313

# Go API (separate terminal)
cd hugo/api && go run main.go         # localhost:8080

# Full stack via Docker
cd hugo && docker compose up --build  # localhost:3000

# Production build
cd hugo && hugo --gc --minify         # Output: hugo/public/
```

## Key Directories

```
hugo/
├── hugo.toml           # Site config and business data
├── assets/css/main.css # Single CSS file (no frameworks)
├── content/            # Page content (markdown)
├── data/               # Portfolio and pricing YAML data
├── layouts/            # Hugo templates
├── static/             # Static assets (images, favicon)
├── api/                # Go contact form handler
├── Caddyfile           # Inner Caddy config
├── Dockerfile          # Three-stage: Hugo + Go + Caddy
└── docker-compose.yml
```

## Key Conventions

- **Business data in hugo.toml** — Never hardcode phone, address, email, or hours in templates. Use `.Site.Params.*`.
- **CSS** — Single file, CSS custom properties for theming. No frameworks.
- **JavaScript** — Inline only, in `{{ define "scripts" }}` blocks. Zero external JS dependencies.
- **Portfolio data** — `data/portfolio.yaml` drives listing and case study metadata.
- **Pricing data** — `data/pricing.yaml` has `subscribe` and `buyout` arrays.

## Color Reference

| Custom Property | Value | Use |
|---|---|---|
| `--color-primary` | `#1e91d9` | Curious Blue — links, buttons, accents |
| `--color-secondary` | `#091740` | Downriver — headings, dark backgrounds |
| `--color-tertiary` | `#00cdb0` | Bright Turquoise — highlights, CTA accents |

## Deployment

Push to `master` or `main` triggers GitHub Actions (`.github/workflows/deploy.yml` inside `hugo/`). See `hugo/CLAUDE.md` for full deployment docs.
