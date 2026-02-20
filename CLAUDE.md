# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Firefly Software company website — Hugo static site + Go contact form API + Docker + Caddy.

**Architecture:**
```
Internet → Outer Caddy (HTTPS/TLS on VPS) → Docker Container (HTTP on configurable port)
                                                ├── Inner Caddy (/api/* reverse proxy + static files)
                                                └── Go API (contact form on localhost:8080)
```

The Docker image is a three-stage build: Hugo (static site) → Go (API binary) → Caddy (runtime). The entrypoint script starts the Go API in the background, waits for its health check, then runs Caddy.

## Commands

```bash
# Hugo dev server
hugo server -D                            # http://localhost:1313

# Go API (separate terminal, for contact form testing)
cd api && go run main.go                  # localhost:8080

# Full stack via Docker
docker compose up --build                 # localhost:3000

# Production build
hugo --gc --minify                        # Output: public/
```

## Key Conventions

### Business Data

Never hardcode phone numbers, addresses, emails, or hours in templates. All business data lives in `hugo.toml` under `[params]`. Access via `.Site.Params.*` in templates.

### Design System — "Montana Utility × Wilderness Wonder"

The visual language uses sharp corners, warm earth tones, and night-sky atmosphere. Key rules:

- **Zero `border-radius` everywhere** — no rounded corners on cards, buttons, inputs, badges, or any UI element. The only exception is circles (firefly dots, progress pips).
- **Never use `#000000` or `#ffffff`** — use `--midnight`/`--navy` for darks, `--snow`/`--granite` for lights.
- **2px gap grid pattern** — signature detail. Set `gap: 2px` on grid containers with background color matching the divider color (`--granite-dk` on light, `--midnight` on dark). Cards fill the cells.
- **Inline SVG only** — no icon font libraries, no external icon CDNs. Icons are stroke-only (`stroke-width: 1.5`, `fill: none`).
- **No heavy animations** — motion is limited to fades, 1-2px translates, and scaleX for hover bars. No bounce, spring, or slide-in effects.

### Color Tokens

All colors are CSS custom properties on `:root` in `assets/css/main.css`. Primary accent colors:

| Token | Hex | Use |
|---|---|---|
| `--rust` | `#8B3A1A` | Primary actions, buttons, links |
| `--ember` | `#C4581A` | Hover states, warm sunrise accent |
| `--amber` | `#E8922A` | Highlights, warm glow, dark section accents |
| `--pine` | `#2C4A2E` | Secondary actions, success states |
| `--midnight` | `#0A0E18` | Deepest dark backgrounds |
| `--navy` | `#0F1628` | Nav, footer, dark sections |
| `--granite` | `#F0EDE8` | Page background |
| `--snow` | `#FAFAF8` | Card surfaces |
| `--ink` | `#1A1A18` | Body text, headings |

### Typography

Four fonts only — no Inter, Roboto, Arial, or system-ui anywhere:

| Variable | Font | Use |
|---|---|---|
| `--font-display` | Playfair Display | Hero titles, section headings |
| `--font-body` | Libre Baskerville | All body copy |
| `--font-ui` | Oswald | Buttons, nav, labels, badges |
| `--font-mono` | Source Code Pro | Metadata, coordinates |

### CSS

Single file at `assets/css/main.css`, processed via Hugo Pipes (`resources.Get | minify`). No CSS frameworks.

### JavaScript

Inline only, placed in `{{ define "scripts" }}` blocks at the bottom of page templates. Zero external JS dependencies.

### Templates & Layouts

- **Base template**: `layouts/_default/baseof.html` — defines `head`, `main`, and `scripts` blocks
- **Page type layouts**: selected via front matter `layout` field (e.g., `layout: contact` maps to `layouts/page/contact.html`)
- **Service pages**: section pages using `_index.md` with layout in `layouts/[section]/list.html`
- **Go template tags**: preserve `{{ }}` blocks exactly including interior whitespace — never alter them

### Data Files

- `data/portfolio.yaml` — portfolio project entries with `slug`, `name`, `category`, `splash`, `logo`, etc. Content body in `content/portfolio/[slug].md`
- `data/pricing.yaml` — `tiers` array with `name`, `buildPrice`, `monthlyPrice`, `features`, `featured`, `cta`

### Cloudflare Turnstile

Contact form uses Turnstile for bot protection. For local dev, use Cloudflare's test keys (already set in `hugo.toml`):
- Site key: `1x00000000000000000000AA` (always passes)
- Secret: `1x0000000000000000000000000000000AA` (always passes)

### Go API

`api/main.go` — zero external Go dependencies (stdlib only). Handles `POST /api/contact` with honeypot field, Turnstile verification, and Postmark email delivery. Gracefully degrades when `TURNSTILE_SECRET` or `POSTMARK_TOKEN` are unset (logs instead).

## Environment Variables

| Variable | Default | Description |
|---|---|---|
| `PORT` | `80` | Caddy listen port inside container |
| `API_PORT` | `8080` | Go API port (localhost only) |
| `ALLOWED_ORIGIN` | `https://fireflysoftware.dev` | CORS allowed origin |
| `TURNSTILE_SECRET` | _(empty — skip)_ | Cloudflare Turnstile secret |
| `POSTMARK_TOKEN` | _(empty — log only)_ | Postmark API token |
| `FROM_EMAIL` | `noreply@fireflysoftware.dev` | Email sender |
| `TO_EMAIL` | `service@fireflysoftware.dev` | Email recipient |
| `LISTEN_PORT` | `3000` | Host port (docker-compose.yml) |

## Deployment

Push to `master` or `main` triggers GitHub Actions (`.github/workflows/deploy.yml`): builds Docker image, pushes to Docker Hub (`dukerupert/fs-website`), SSHes to VPS, pulls and restarts.

Required GitHub Secrets: `DOCKERHUB_USERNAME`, `DOCKERHUB_TOKEN`, `PROD_VPS_HOST`, `PROD_VPS_USER`, `PROD_VPS_SSH_KEY`, `PROD_VPS_DEPLOY_PATH`, `PROD_PORT`
