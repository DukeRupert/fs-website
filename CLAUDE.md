# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Firefly Software company website built with SvelteKit 1.x, TypeScript, and TailwindCSS. The site uses Sanity CMS for blog content and Web3Forms for contact form submissions.

## Commands

```bash
# Development
pnpm dev              # Start dev server (http://localhost:5173)

# Build & Preview
pnpm build            # Production build (outputs to /build)
pnpm preview          # Preview production build

# Type Checking
pnpm check            # Run svelte-check once
pnpm check:watch      # Run svelte-check in watch mode
```

## Architecture

### Tech Stack
- **Framework**: SvelteKit 1.x with adapter-node (SSR deployment)
- **Styling**: TailwindCSS with custom color palette (primary/secondary/tertiary)
- **CMS**: Sanity (projectId: vpzagt04, dataset: production) for blog posts
- **Images**: @zerodevx/svelte-img for optimized image handling with `?as=run` imports
- **Forms**: Web3Forms (external service, no backend needed)

### Route Structure
```
src/routes/
├── +layout.svelte          # Root layout (Navigation + Footer)
├── +page.svelte            # Homepage
├── api/posts/              # API endpoint for paginated blog posts
├── contact-us/             # Contact form page
├── e-commerce/             # E-commerce service page
├── portfolio/              # Portfolio pages with case studies
│   ├── projects.ts         # Portfolio data (CaseStudy[])
│   └── [project]/          # Individual project pages
├── posts/                  # Blog listing and [slug] pages
├── pricing/                # Pricing page
├── website-development/    # Service page
└── website-maintenance/    # Service page
```

### Key Patterns

**Sanity Integration** (`src/lib/db.ts`):
- `client` - Configured Sanity client
- `urlFor(source)` - Image URL builder for Sanity assets

**Portable Text**: Custom components in `src/lib/components/portableText/` for rendering Sanity rich text blocks.

**Image Optimization**: Import images with `?as=run` suffix for svelte-img processing:
```typescript
import image from "$lib/assets/images/example.jpg?as=run";
```

**SEO Component** (`src/lib/components/SEO.svelte`): Pass seoData object for meta tags and Open Graph.

### Tailwind Color Palette
- `primary` - Curious Blue (#1e91d9)
- `secondary` - Downriver (#091740)
- `tertiary` - Bright Turquoise (#00cdb0)

## Deployment

Docker-based deployment using adapter-node:
```bash
docker build -t dukerupert/fs-website:latest .
docker compose up -d
```

Container exposes port 3000. Requires `ORIGIN` environment variable set to production domain.
