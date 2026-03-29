# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Firefly Software company website — Go server with `html/template` + Tailwind CSS v4 + Alpine.js + Docker + Caddy.

**Architecture:**
```
Internet → Outer Caddy (HTTPS/TLS on VPS) → Docker Container (HTTP on configurable port)
                                                ├── Caddy (port 80, reverse proxy + compression + headers)
                                                └── Go Server (port 8080)
                                                    ├── html/template page rendering
                                                    ├── Static file serving (/static/*)
                                                    ├── Contact form API (POST /api/contact)
                                                    ├── Blog post rendering (goldmark markdown)
                                                    └── 301 redirects for old routes
```

The Docker image is a three-stage build: Tailwind CSS (Node) → Go (server binary) → Caddy (runtime). The entrypoint script starts the Go server in the background, waits for its health check, then runs Caddy.

## Commands

```bash
# Dev server (Tailwind watcher + Go server)
make dev                                  # http://localhost:8080

# Go server only (no CSS watcher)
make dev-server                           # localhost:8080

# Tailwind CSS watcher only
make dev-css

# Production CSS build
make build-css

# Production binary
make build                                # Outputs: ./fs-website

# Full stack via Docker
docker compose up --build                 # localhost:3000
```

## Key Conventions

### Business Data

Phone number, email, and site metadata are defined in `handlers/middleware.go` via `NewSiteData()`. Access via `.Site.*` or `.Phone` / `.Email` in templates.

### Design System — "Precision Instrument"

The visual language is restrained, technically confident, and warm without being rustic. Think "well-made tool" — everything is exactly where it needs to be and nothing is there for decoration. The one decorative exception: the firefly glow accent (`#C8F060`), which should feel earned when it appears.

Key rules:

- **Max `4px` border-radius** — only on cards and form fields. No large rounded corners.
- **No drop shadows** — shadows make things look like web templates.
- **No stock photography** — typography does the work.
- **Accent color is never a background fill** — only text, borders, underlines, hover states.
- **One accent moment per section maximum** — restraint makes the glow impactful.
- **Inline SVG only** — no icon font libraries, no external icon CDNs.

### Color Tokens

All colors are CSS custom properties on `:root` in `static/css/tokens.css`:

**Dark palette (hero, nav, footer, CTA sections):**

| Token | Hex | Use |
|---|---|---|
| `--color-dark-bg` | `#0E0F0D` | Near-black with warm undertone |
| `--color-dark-surface` | `#161810` | Cards, nav, elevated elements |
| `--color-dark-border` | `#2A2D26` | Subtle dividers |
| `--color-dark-text` | `#F0EDE6` | Warm off-white primary text |
| `--color-dark-muted` | `#A3A59B` | Secondary/muted text (AAA contrast) |
| `--color-sky-top` | `#0A0C14` | Hero sky gradient top (home page only) |
| `--color-sky-horizon` | `#1C1A0F` | Hero sky gradient horizon (home page only) |

**Light palette (body sections):**

| Token | Hex | Use |
|---|---|---|
| `--color-light-bg` | `#F7F5F0` | Warm parchment background |
| `--color-light-surface` | `#FFFFFF` | Cards and elevated elements |
| `--color-light-border` | `#E2DDD6` | Warm grey borders |
| `--color-light-text` | `#1A1C18` | Near-black body text |
| `--color-light-muted` | `#6B6D63` | Secondary/muted text |

**Accent — two tokens, surface-dependent:**

| Token | Hex | Use |
|---|---|---|
| `--color-accent` | `#C8F060` | Firefly glow on dark surfaces — hero, nav, footer, dark cards |
| `--color-accent-light` | `#4D6B00` | Accent text on light surfaces — labels, links, step numbers |
| `--color-accent-light-border` | `#6B8F0A` | Accent borders on light surfaces — card hovers, testimonials, asides |

The two-accent system keeps the brand thread on both surfaces while solving contrast. `#C8F060` on dark is the firefly glow. `#4D6B00` on light is the same hue darkened for AAA contrast on parchment. See `component-guide.html` for the full reference.

### Typography

Three fonts only — no Inter, Roboto, Arial, or system-ui anywhere:

| Variable | Font | Use |
|---|---|---|
| `--font-display` | Bebas Neue | Hero titles, section headings, section labels (all caps) |
| `--font-body` | Instrument Sans | All body copy, buttons, nav links |
| `--font-mono` | IBM Plex Mono | Coordinates, metadata, code snippets |

Fonts are self-hosted in `static/fonts/` as woff2.

### CSS

Tailwind CSS v4 with design tokens in `static/css/tokens.css`. Entry point is `static/css/input.css` which imports Tailwind and tokens. Output goes to `static/css/output.css` (gitignored, built by `make build-css` or the watcher).

Reusable component classes are defined in `tokens.css`: `.btn-primary`, `.btn-secondary`, `.section-label`, `.card-dark`, `.card-light`, `.card-testimonial`, `.aside-accent`, `.hero`, `.hero--short`, `.hero-headline`, `.gradient-transition`, `.logo-strip`, `.form-field`, `.form-label`, `.content-wrap`, `.content-narrow`, `.section`, `.reveal`.

### JavaScript

Alpine.js vendored at `static/js/alpine.min.js`. Used for:
- Mobile nav open/close and scroll detection
- Services dropdown
- Scroll reveal (IntersectionObserver on body element)
- Contact form state management (idle → loading → success/error)

Page-specific scripts go in `{{ define "scripts" }}` blocks.

### Templates & Layouts

- **Base template**: `templates/base.html` — defines `head`, `nav`, `content`, `footer`, and `scripts` blocks
- **Partials**: `templates/nav.html`, `templates/footer.html`
- **Page templates**: `templates/home.html`, `templates/work.html`, `templates/process.html`, `templates/about.html`, `templates/contact.html`, `templates/services-websites.html`, `templates/services-software.html`
- **Blog**: `templates/post.html` (single), `templates/posts.html` (listing)
- **Go template tags**: preserve `{{ }}` blocks exactly including interior whitespace — never alter them

### Go Server

`main.go` at repo root — routes, template parsing, Sentry init, middleware stack. `handlers/` package contains:
- `contact.go` — POST /api/contact with honeypot, Turnstile verification, Postmark email
- `middleware.go` — CORS, recovery, logging, redirect helpers, PageData/SiteData types
- `pages.go` — TemplateRenderer with dev-mode hot reload
- `posts.go` — Blog post markdown rendering with goldmark

### Blog Posts

Markdown files in `content/posts/` with YAML front matter. Rendered via goldmark at request time. All internal links point to `/contact-us/` which 301-redirects to `/contact`.

### Motion

Restrained. The site should feel fast and deliberate, not animated.

- **Hero headline:** Single fade-up on load, `0.6s ease-out`. Nothing else animates on load.
- **Scroll reveals:** `.reveal` → `.is-visible` via IntersectionObserver. `opacity 0→1`, `translateY 16px→0`, `0.4s ease-out`. Staggered for card grids.
- **Hover states:** `0.15s` transition. Fast enough to feel responsive.
- **No parallax. No looping animations. No cursor effects.**

## Environment Variables

| Variable | Default | Description |
|---|---|---|
| `PORT` | `80` | Caddy listen port inside container |
| `API_PORT` | `8080` | Go server port |
| `DEV_MODE` | `false` | Enable template hot reload |
| `ALLOWED_ORIGIN` | `https://fireflysoftware.dev` | CORS allowed origin |
| `TURNSTILE_SECRET` | _(empty — skip)_ | Cloudflare Turnstile secret |
| `POSTMARK_TOKEN` | _(empty — log only)_ | Postmark API token |
| `SENTRY_DSN` | _(empty — skip)_ | Bugsink/Sentry DSN for error reporting |
| `SENTRY_ENVIRONMENT` | `production` | Environment tag for error reports |
| `FROM_EMAIL` | `noreply@fireflysoftware.dev` | Email sender |
| `TO_EMAIL` | `logan@fireflysoftware.dev` | Email recipient |
| `LISTEN_PORT` | `3000` | Host port (docker-compose.yml) |

## Deployment

Push to `master` triggers GitHub Actions (`.github/workflows/deploy-prod.yml`): builds Docker image, pushes to Docker Hub (`dukerupert/fs-website`), SSHes to VPS, pulls and restarts.

Required GitHub Secrets: `DOCKERHUB_USERNAME`, `DOCKERHUB_TOKEN`, `VPS_HOST`, `PROD_VPS_USER`, `PROD_VPS_SSH_KEY`, `PROD_VPS_DEPLOY_PATH`

## Design Context

Read alongside `visual-design-direction.md` for full visual specification.

### Users

Small business owners evaluating a custom development studio — often arriving cold from search or a referral. Many are coastal (California, Florida, Pacific Northwest) and are comparing Firefly against local agencies or freelancers. They're pragmatic buyers: they want to know what they'll get, what it costs, and whether the people behind the studio can be trusted. They don't want a sales pitch. They want a straight answer.

### Brand Personality

**Precise. Warm. Direct.**

Firefly speaks with quiet confidence — technically excellent but never showy. The tone is first-person plural ("we"), conversational but never casual, and specific rather than aspirational. It reads like a letter from someone who builds things for a living, not a marketing team.

The emotional reference is *Yellowstone* — not the props (no barbed wire, no weathered wood), but the feeling: unhurried confidence, landscape as character, craft that's been tested.

### Aesthetic Direction

**"Precision instrument, Montana made."** The visual language of a well-made tool. Dark above the fold (late Montana night sky), warm parchment below.

**The One Memorable Thing:** The first `#C8F060` firefly glow against the near-black hero — the ridgeline below it, the coordinates in mono type beside it. Quiet, specific, earned. Everything else on the site exists to not compete with that moment.

**References:** Pared-back product pages (Stripe's developer docs, Linear's marketing site) — confident typography, generous whitespace, zero decoration.

**Anti-references:** Creative agency sites with parallax, animated counters, gradient buttons, purple/teal accents, stock photography hero sections, or mega-menu navigation. No drop shadows, no gradient fills on buttons, no rounded corners above 6px, no full-bleed background images, no animated stat counters.

**Theme:** Dark hero → gradient transition → warm light body. Both halves share the same accent color as a thread.

### Hero Composition (Home Page Only)

The home page hero is the one place on the site where the Montana landscape is shown directly. Three elements:

1. **Sky gradient** — CSS gradient from `#0A0C14` (top) through `#0E0F0D` (55%) to `#1C1A0F` (bottom). Two to three stops. The kind of sky where fireflies come out. Home page hero only; all other dark sections use flat `#0E0F0D`.

2. **Ridgeline** — SVG silhouette of the Sleeping Giant profile. Fill `#1A1C18`. Positioned `absolute; bottom: 0`. Full viewport width. Dark enough to read as shape, warm enough to not disappear.

3. **Firefly glows** — Two to three static radial gradient glows (`#C8F060` center → transparent). Positioned above the ridgeline. No animation. They are simply present, the way fireflies are present at dusk if you happen to look.

Type sits in the upper two-thirds. The landscape owns the bottom third. They do not compete.

### The Firefly Magic — Rules of Use

The `#C8F060` glow is the single decorative thread. Its power comes entirely from restraint.

**Where it appears (decorative):**
- Hero glows — static, home page only
- Logo mark — single small dot near the ridgeline in the SVG
- Scroll-transition pulse — one firefly glow pulses once at the dark-to-light transition, triggered by IntersectionObserver, fires once per page load
- Process page step glows — brief `4px` glow next to step numbers on scroll-into-view, `0.3s` then gone

**Where it appears (functional):**
- `--color-accent` (#C8F060): section labels on dark bg, active nav underlines, CTA buttons on dark bg, focus-visible outlines, dark card labels
- `--color-accent-light` (#4D6B00): section labels on light bg, buttons/links on light bg, step numbers, team roles, card eyebrows
- `--color-accent-light-border` (#6B8F0A): card hover borders on light bg, testimonial left borders, accent asides on light bg

**Where it never appears:**
- As a background fill (exception: `.btn-primary:hover` fill)
- In any looping animation
- More than once per section in decorative form
- `#C8F060` directly on light surfaces (use `--color-accent-light` instead)

**Governing principle:** A firefly never performs. It appears when the conditions are right and then it's gone. If it feels like a design decision, it's too much.

### Interior Pages

Interior heroes (Process, Services, Work, About, Contact) use flat `#0E0F0D` — no sky gradient, no ridgeline, no firefly glows. The landscape composition belongs to the home page alone.

Light-background sections on interior pages get topographic contour lines as subtle texture: thin strokes (`0.5px`), `--color-light-border`, `opacity: 0.4`, repeating SVG background. Barely perceptible — texture, not decoration. (Asset: `static/images/topo-pattern.svg` — needs creation.)

### Emotional Goals

A cold visitor should feel **curiosity and respect** ("this isn't what I expected from a Montana studio") layered with **confidence and calm** ("these people know what they're doing and I can trust them"). The design earns trust through restraint, not through volume.

### Accessibility

Target **WCAG AAA where feasible**. Enhanced contrast ratios, `prefers-reduced-motion` support (all animations/transitions disabled), `:focus-visible` outlines on all interactive elements, keyboard navigation, semantic HTML landmarks, descriptive alt text, associated form labels. Color is never the sole indicator of state.

### Design Principles

1. **Restraint is the design.** Every element earns its place. If it doesn't serve clarity or trust, it doesn't belong. One accent moment per section maximum.

2. **Typography does the work.** No hero images, no stock photography, no illustration beyond the Sleeping Giant ridgeline. Bebas Neue at display size, generous whitespace, and the firefly glow are the entire visual vocabulary.

3. **Warm, not sterile.** Colors have undertones (warm off-white, not clinical white; warm near-black, not blue-black). The site should feel handmade in the best sense — considered, not generated.

4. **Motion is earned.** One fade-up on load. Scroll reveals at 0.4s. Hover states at 0.15s. One scroll-transition pulse. Nothing bounces, springs, loops, or parallaxes. Speed signals competence.

5. **Accessibility is structural, not decorative.** Semantic HTML, landmark roles, AAA contrast where feasible, reduced-motion support. These aren't features — they're how the site is built.

6. **The firefly never performs.** Every decorative use of the accent color should feel like you caught the scene at exactly the right moment. If it feels like a design decision, it's too much.
