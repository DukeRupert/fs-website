# Firefly Software — Coding Agent Brief
*Site Rebuild · v1.0*

---

## Overview

This brief instructs a coding agent to rebuild fireflysoftware.dev as a coastal-positioned custom development studio site. The existing site is built in Go with `html/template`. The rebuild extends the existing repo — it is not a greenfield project.

All copy is finalized in `firefly-homepage-copy.md`. All visual design decisions are finalized in the Visual Design Direction section of that same document. This brief translates both into implementation instructions.

Work in small, testable commits. Each phase below represents one logical commit boundary.

---

## Stack

- **Backend:** Go, `html/template`
- **CSS:** Tailwind CSS v4 with a custom `@layer base` block for design tokens
- **JS:** Alpine.js for interactive components (mobile nav, scroll reveals, contact form state)
- **Fonts:** Vendored locally — Bebas Neue, Instrument Sans, IBM Plex Mono
- **No external CSS frameworks. No jQuery. No component libraries.**

---

## File & Asset Assumptions

The following assets exist in the repo and should be referenced as-is:

- `/static/images/logo/logo_two_full.png` — primary logo (light version)
- `/static/images/logo/Firefly_Logo_Blue.png` — alternate logo
- `/static/images/sleeping-giant.svg` — Sleeping Giant SVG for hero background texture
- `/static/fonts/` — vendored font files (Bebas Neue, Instrument Sans, IBM Plex Mono)
- `/static/images/portfolio/` — existing client logo and photo assets

A new dark-mode logo variant may be needed. If `logo_two_full.png` does not render legibly on `#0E0F0D`, flag it — do not substitute.

---

## Route Map

Implement the following routes. Existing routes not listed here should be left untouched until explicitly instructed.

| Route | Template | Notes |
|---|---|---|
| `/` | `home.html` | Full rebuild |
| `/work` | `work.html` | Replaces `/portfolio/` |
| `/services/websites` | `services-websites.html` | Replaces `/website-development/` |
| `/services/software` | `services-software.html` | Replaces `/custom-software/` |
| `/process` | `process.html` | New page — does not exist yet |
| `/about` | `about.html` | Full rebuild |
| `/contact` | `contact.html` | Rebuild of `/contact-us/` |

Existing routes (`/website-maintenance/`, `/site-rescue/`, `/outfitters/`, `/pricing/`) should return HTTP 301 redirects to their nearest replacement:

| Old route | Redirects to |
|---|---|
| `/website-maintenance/` | `/services/websites` |
| `/site-rescue/` | `/services/websites` |
| `/outfitters/` | `/work` |
| `/pricing/` | `/services/websites` |
| `/portfolio/` | `/work` |
| `/contact-us/` | `/contact` |

---

## Phase 1 — Design Tokens & Base CSS

**Commit scope:** CSS foundation only. No templates touched.

Create `/static/css/tokens.css`. This file is imported before Tailwind's base layer.

```css
@layer base {
  :root {
    /* Dark palette */
    --color-dark-bg:        #0E0F0D;
    --color-dark-surface:   #161810;
    --color-dark-border:    #2A2D26;
    --color-dark-text:      #F0EDE6;
    --color-dark-muted:     #8A8C82;

    /* Light palette */
    --color-light-bg:       #F7F5F0;
    --color-light-surface:  #FFFFFF;
    --color-light-border:   #E2DDD6;
    --color-light-text:     #1A1C18;
    --color-light-muted:    #6B6D63;

    /* Accent — same in both modes */
    --color-accent:         #C8F060;

    /* Typography */
    --font-display:         'Bebas Neue', sans-serif;
    --font-body:            'Instrument Sans', sans-serif;
    --font-mono:            'IBM Plex Mono', monospace;

    /* Fluid type scale */
    --text-display: clamp(4.5rem, 8vw, 7rem);
    --text-h1:      clamp(2.5rem, 4vw, 3.5rem);
    --text-h2:      clamp(1.5rem, 2.5vw, 2rem);
    --text-body:    clamp(1rem, 1.2vw, 1.125rem);
    --text-small:   0.8125rem;

    /* Spacing */
    --section-padding-y:    7.5rem;   /* 120px desktop */
    --content-max-width:    1180px;
    --border-radius-card:   4px;
  }
}

/* Font face declarations */
@font-face {
  font-family: 'Bebas Neue';
  src: url('/static/fonts/BebasNeue-Regular.woff2') format('woff2');
  font-weight: 400;
  font-display: swap;
}

@font-face {
  font-family: 'Instrument Sans';
  src: url('/static/fonts/InstrumentSans-Regular.woff2') format('woff2');
  font-weight: 400;
  font-display: swap;
}

@font-face {
  font-family: 'Instrument Sans';
  src: url('/static/fonts/InstrumentSans-Italic.woff2') format('woff2');
  font-weight: 400;
  font-style: italic;
  font-display: swap;
}

@font-face {
  font-family: 'IBM Plex Mono';
  src: url('/static/fonts/IBMPlexMono-Regular.woff2') format('woff2');
  font-weight: 400;
  font-display: swap;
}

/* Base resets */
body {
  font-family: var(--font-body);
  font-size: var(--text-body);
  color: var(--color-light-text);
  background-color: var(--color-light-bg);
  line-height: 1.7;
  -webkit-font-smoothing: antialiased;
}

/* Scroll reveal base state — Alpine.js toggles x-reveal class */
[x-cloak] { display: none !important; }

.reveal {
  opacity: 0;
  transform: translateY(16px);
  transition: opacity 0.4s ease-out, transform 0.4s ease-out;
}

.reveal.is-visible {
  opacity: 1;
  transform: translateY(0);
}
```

**Test:** Confirm tokens load, fonts render, no Tailwind conflicts.

---

## Phase 2 — Base Layout Templates

**Commit scope:** Shared layout templates only. No page content yet.

### `layouts/base.html`

The base layout wraps all pages. Structure:

```
<html>
  <head>
    <!-- meta, title, CSS -->
  </head>
  <body>
    {{ template "nav" . }}
    <main>{{ block "content" . }}{{ end }}</main>
    {{ template "footer" . }}
  </body>
</html>
```

### `layouts/nav.html`

**Structure:** Dark background (`var(--color-dark-bg)`). Logo left, nav links right. Mobile hamburger right on small screens.

**Nav links:** Work · Services · Process · About · Contact

Services is a simple dropdown with two items: Websites and Software. No mega-menu. Dropdown appears on hover (desktop) and tap (mobile).

**Active state:** 2px underline in `var(--color-accent)` on the active nav item. No background highlight.

**Mobile nav:** Full-screen overlay. Background `var(--color-dark-bg)`. Nav items in `var(--font-display)` at large size (`2.5rem`). Active item accent colored. Alpine.js controls open/close state.

```html
<!-- Alpine.js mobile nav controller -->
<nav x-data="{ open: false }">
  <!-- hamburger button: @click="open = !open" -->
  <!-- overlay: x-show="open" x-transition -->
</nav>
```

**Sticky:** Nav sticks to top. On scroll past 60px, add a `1px` bottom border in `var(--color-dark-border)` and slight background opacity increase. No box shadow.

### `layouts/footer.html`

**Structure:** Dark background matching nav. Three columns on desktop, stacked on mobile.

- Col 1: Logo + tagline — *"Websites, fixes, and software — built for small businesses."*
- Col 2: Nav links (Work, Services, Process, About, Contact) + Posts
- Col 3: Contact info (phone, email)

**Bottom bar:** Single line. Left: `© 2026 Firefly Software, LLC.` Right: `✦ Made in Montana by Firefly Software`

**Remove:** "Service Areas: Helena, MT / All of Montana" — replace with: *"Based in Helena, Montana. Working with clients across the US."*

**Test:** Nav and footer render correctly at mobile (375px), tablet (768px), desktop (1280px). Mobile nav opens and closes. Dropdown works.

---

## Phase 3 — Home Page

**Commit scope:** `/` route and `home.html` template only.

Build sections in order. Each section is a separate `<section>` element with a consistent wrapper div: `max-width: var(--content-max-width); margin: 0 auto; padding: 0 1.5rem;`

### Section 1 — Hero

**Background:** `var(--color-dark-bg)`. Full viewport height minimum (`min-height: 100vh`).

**Background texture:** Sleeping Giant SVG positioned absolutely, bottom-right, `opacity: 0.06`, `pointer-events: none`. Do not scale beyond its natural aspect ratio.

**Content — centered column, max-width 780px:**

- Eyebrow label: `FIREFLY SOFTWARE` in `var(--font-mono)`, `var(--text-small)`, `var(--color-accent)`, `letter-spacing: 0.2em`
- Headline: copy from document, `var(--font-display)`, `var(--text-display)`, `var(--color-dark-text)`, `line-height: 0.95`
- Subhead: copy from document, `var(--font-body)`, `1.2rem`, `var(--color-dark-muted)`, `max-width: 560px`
- CTAs: two buttons — primary and secondary styles (see Component Spec below)
- Coordinates: `46.5958°N · 112.0270°W` in `var(--font-mono)`, `var(--text-small)`, `var(--color-dark-muted)`, bottom of section

**Animation:** Hero headline fades up on load — `opacity: 0 → 1`, `translateY: 20px → 0`, `0.6s ease-out`. Apply via CSS class added by a small inline script after DOM ready. Nothing else animates on load.

### Transition

Between hero and first light section: gradient `var(--color-dark-bg)` → `var(--color-light-bg)` over `80px`. Implement as a `<div>` with `height: 80px` and the gradient as background.

### Section 2 — Credentialing

**Background:** `var(--color-light-bg)`.

**Layout:** Single centered column, max-width `680px`.

**Section label:** `WHO WE ARE` — label style (see Component Spec).

**Content:** Three paragraphs from copy document. Body text size, `var(--color-light-text)`.

### Section 3 — What We Build

**Background:** `var(--color-light-bg)`.

**Section label:** `WHAT WE BUILD`

**Layout:** Three cards in a row on desktop, stacked on mobile.

Each card: white surface (`var(--color-light-surface)`), `1px` border (`var(--color-light-border)`), `var(--border-radius-card)`. Headline in `var(--font-display)` at `1.6rem`. Body in `var(--font-body)`. Secondary CTA link at bottom.

Hover state: border color transitions to `var(--color-accent)`, `translateY(-2px)`. Transition `0.15s`.

**Scroll reveal:** Cards stagger — delays of `0s`, `0.1s`, `0.2s`.

### Section 4 — What Clients Say

**Background:** `var(--color-light-bg)`.

**Section label:** `WHAT CLIENTS SAY`

**Layout:** Three testimonial cards. On desktop: three columns. On mobile: stacked.

Each card: `var(--color-light-surface)`, `3px` left border `var(--color-accent)`, `var(--border-radius-card)`, `1.5rem` padding. Quote in `Instrument Sans` italic. Attribution in small caps, `var(--color-light-muted)`.

No carousel. All three visible simultaneously.

**Logo strip:** Below testimonials. Logos in a flex row, `gap: 2.5rem`, centered. Images in grayscale (`filter: grayscale(1)`), `opacity: 0.5`. On hover: `opacity: 1`, grayscale removed. Transition `0.2s`.

**Guarantee block:** Below logos. Centered. Four sentences from copy document. `var(--font-body)`, slightly larger than body (`1.1rem`), `var(--color-light-muted)`. No border, no card — just text with generous padding above and below.

### Section 5 — CTA / Footer Close

**Background:** `var(--color-dark-bg)`.

**Layout:** Centered column, max-width `620px`, generous vertical padding.

**Section label:** `LET'S TALK` in accent color.

**Headline:** `var(--font-display)`, `var(--text-h1)`, `var(--color-dark-text)`.

**Body:** Two paragraphs from copy document. `var(--color-dark-muted)`.

**CTAs:** Primary button + phone number as plain text link.

**Test:** Home page renders correctly at all three breakpoints. Hero animation fires once on load. Scroll reveals trigger correctly. Card hover states work. Logo strip grayscale toggles.

---

## Phase 4 — Process Page

**Commit scope:** `/process` route and `process.html` only.

**Hero:** Shared dark hero partial — same nav, same dark background, but shorter (`min-height: 50vh`). Page title in `var(--font-display)` at `var(--text-h1)`. Subhead from copy document.

**Framing block:** Below hero transition. Single centered column, max-width `680px`. The philosophy paragraph from copy document. Slightly larger text (`1.15rem`), `var(--color-light-muted)`. Italic. This is the only italic body block on the page — use it to signal that this is a framing statement, not a step.

**Steps:** Six steps — Discover, Agree, Research, Build, Refine, Launch.

Each step:
- Step number: `var(--font-mono)`, `var(--text-small)`, `var(--color-accent)`, e.g. `01`
- Step name: `var(--font-display)`, `var(--text-h2)`, `var(--color-light-text)`
- Body: copy from document, `var(--font-body)`
- Separated by a `40px` vertical gap — no horizontal rules

On desktop: two-column layout — step number/name left, body right. On mobile: stacked.

**"What happens if something goes wrong?" block:**

Visually distinct. Dark surface (`var(--color-dark-surface)`), `1px` border `var(--color-dark-border)`, `var(--border-radius-card)`, generous padding. Header in `var(--font-display)`. Body from copy document. This is the one dark card in an otherwise light page — use the contrast deliberately.

**CTA:** Standard page-close CTA from copy document. Primary + phone.

**Test:** Steps render correctly at mobile. The "what goes wrong" block renders on its dark surface. Page reads top to bottom without visual noise.

---

## Phase 5 — Services Pages

**Commit scope:** `/services/websites` and `/services/software` routes and templates.

Both pages share a structure: dark short hero → light body sections → dark CTA close.

### Services / Websites (`services-websites.html`)

Sections in order:
1. Dark short hero — headline and subhead from copy
2. Opening copy — two paragraphs, single centered column
3. What's Included — single column, body text
4. Turnaround — single column, body text
5. Pricing block — slightly elevated card (`var(--color-light-surface)`), border, the two pricing tiers and monthly support rates from copy. This is the one place on the page where numbers get visual emphasis — use `var(--font-display)` for the dollar amounts at `2rem`
6. "Not sure if you need a full rebuild?" — subtle aside. `var(--color-light-bg)` background, `3px` left border `var(--color-accent)`, the two paragraphs from copy, a secondary CTA
7. Dark CTA close — headline, body, primary button + phone

### Services / Software (`services-software.html`)

Sections in order:
1. Dark short hero
2. Opening copy — two paragraphs
3. What We Build — prose section from copy. No bullet list. Body text, single column
4. How We Price It — single column. The three sentences from copy. Straightforward
5. The Short Version — visually distinct block. Same dark card treatment as the "what goes wrong" block on the Process page. The philosophy statement from copy. Use `var(--font-display)` for the opening sentence at slightly larger size
6. Dark CTA close

**Test:** Both pages render correctly. Pricing numbers are visually prominent on the Websites page. The dark philosophy card renders correctly on Software.

---

## Phase 6 — Work Page

**Commit scope:** `/work` route and `work.html` template.

**Hero:** Short dark hero. "Work" as display headline. Subhead from copy document.

**Websites section:**

Section label: `WEBSITES`

Three project cards — Nautilus, A-Team Asphalt, A-Team Gutters. Each card:
- Client name in `var(--font-display)` at `1.4rem`
- Industry + location in `var(--font-mono)`, `var(--text-small)`, `var(--color-light-muted)`, italic
- Body copy from document (Problem → Approach → Result)
- "View project →" secondary CTA link — only if a live URL exists. If no URL, omit the link entirely rather than link to a placeholder

Card layout: on desktop, cards in a 3-column grid. On mobile, stacked.

**Software section:**

Section label: `SOFTWARE`

Two project entries — Western Skies and Rockabilly Roasting. Same card structure as above. No "View project →" link on either — these are internal/NDA projects. The copy stands alone.

**"More Coming" block:**

Below the software cards. Single centered column, max-width `560px`. The three sentences from copy document. `var(--color-light-muted)`. Secondary CTA → `/contact`.

**Test:** All five project cards render. Software cards have no broken CTA links. Page reads as two distinct sections — Websites and Software — without feeling like two separate pages.

---

## Phase 7 — About Page

**Commit scope:** `/about` route and `about.html` template.

**Hero:** Short dark hero. "About" as display headline. No subhead needed.

**Opening block:** Two paragraphs from copy. Single centered column, max-width `680px`.

**Team section:**

Section label: `THE TEAM`

Two-column layout on desktop (Logan left, Jaden right), stacked on mobile.

Each team member:
- Photo (existing assets for Logan — `/static/images/about/logan_and_elliot_williams.webp`). Jaden: placeholder until asset is provided. Do not use a stock photo — use a simple monogram placeholder: dark surface card, initials `JM` in `var(--font-display)` centered, `var(--color-accent)` text.
- Name in `var(--font-display)` at `1.8rem`
- Title: `Co-Founder` in `var(--font-mono)`, `var(--text-small)`, `var(--color-accent)`
- Bio copy from document. Logan's bio is three paragraphs. Jaden's is one paragraph with an italicized `[Expand with Jaden's input when available.]` note — styled as `var(--color-light-muted)`, smaller size, so it reads as a placeholder rather than finished copy.

**"Built in Montana. Built for America." block:**

Full-width section with `var(--color-dark-bg)` background. Centered column. Section label `WHERE WE'RE FROM` in accent color. Headline in `var(--font-display)`. Body from copy document.

**"What We Won't Do" block:**

Back to light background. Single centered column. Section label `WHAT WE WON'T DO` in accent color — this label doing the work of signaling that this is a deliberate standards statement, not fine print. Two paragraphs from copy.

**CTA:** Standard dark close. Headline and body from copy. Primary button + phone.

**Test:** Team section renders at mobile with stacked layout. Jaden placeholder is visually consistent with the page — not broken-looking. Dark Montana block renders correctly between two light sections.

---

## Phase 8 — Contact Page

**Commit scope:** `/contact` route and `contact.html` template.

**Hero:** Short dark hero. "Let's talk." as display headline — note the period, it belongs.

**Opening block:** Two paragraphs from copy. Single centered column, max-width `620px`.

**"One Thing Worth Knowing" block:**

Visually distinct aside — same treatment as the aside on the Websites service page. `3px` left border `var(--color-accent)`, slightly inset from the main column. The paragraph from copy document.

**Form:**

Three fields only:
- Name — `<input type="text">`
- Email — `<input type="email">`
- Tell us what you're working on — `<textarea rows="5">`

Field styling: `var(--color-light-surface)` background, `1px` border `var(--color-light-border)`, `var(--border-radius-card)`. On focus: border transitions to `var(--color-accent)`. No floating labels — standard `<label>` above each field.

Submit button: primary button style. Label: `Send it →`

Alpine.js handles form state — loading, success, error. On success: replace form with a confirmation message. Do not redirect.

**Below the form:**

Plain text: *"Prefer to call? We pick up."* followed by the phone number as a `tel:` link. `var(--color-light-muted)` text. No styling beyond that.

**Test:** Form renders correctly. Focus states work. Alpine.js state transitions (loading → success / error) work without page reload. Mobile layout is comfortable — fields are full width.

---

## Phase 9 — Redirects & Cleanup

**Commit scope:** Routing cleanup only.

1. Implement 301 redirects per the Route Map table in this brief
2. Remove "Service Areas" section from footer
3. Add "Based in Helena, Montana. Working with clients across the US." to footer
4. Confirm all old internal links in existing blog posts (`/posts/`) still resolve or redirect correctly — blog posts are not being rebuilt, only their internal links need to stay valid
5. Update `<title>` tags and meta descriptions on all new pages — copy from the page headlines and subheads is sufficient for now, SEO optimization is a separate pass

---

## Component Spec Reference

### Primary Button
```
border: 1px solid var(--color-accent)
color: var(--color-accent)
background: transparent
padding: 0.75rem 1.75rem
font-family: var(--font-body)
font-size: 0.9375rem
letter-spacing: 0.02em
border-radius: var(--border-radius-card)
transition: background 0.15s, color 0.15s

:hover {
  background: var(--color-accent)
  color: var(--color-dark-bg)
}
```

### Secondary Button / Link
```
color: var(--color-light-text) [or var(--color-dark-text) on dark bg]
background: none
border: none
font-family: var(--font-body)
font-size: var(--text-body)

:hover { color: var(--color-accent) }

Append " →" as text — not a pseudo-element, not an icon
```

### Section Label
```
font-family: var(--font-display)
font-size: 0.8125rem
color: var(--color-accent)
letter-spacing: 0.15em
text-transform: uppercase
display: block
margin-bottom: 0.75rem
```

### Dark Card (Process "what goes wrong", Software philosophy)
```
background: var(--color-dark-surface)
border: 1px solid var(--color-dark-border)
border-radius: var(--border-radius-card)
padding: 2.5rem
color: var(--color-dark-text)
```

### Accent Aside (Websites "not sure", Contact "one thing worth knowing")
```
border-left: 3px solid var(--color-accent)
padding-left: 1.5rem
margin: 2.5rem 0
```

---

## Scroll Reveal Implementation

Use Alpine.js with an `IntersectionObserver` to add `is-visible` to `.reveal` elements as they enter the viewport.

```html
<!-- On the body or a wrapper element -->
<div x-data x-init="
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(e => {
      if (e.isIntersecting) {
        e.target.classList.add('is-visible');
        observer.unobserve(e.target);
      }
    });
  }, { threshold: 0.1 });
  document.querySelectorAll('.reveal').forEach(el => observer.observe(el));
">
```

Apply `.reveal` class to: section headers, card grids (with stagger via inline `style="transition-delay: 0.1s"`), team bios, and the two dark callout cards. Do not apply to nav, footer, or the hero headline (which has its own load animation).

---

## Notes for the Agent

- **Commit after each phase.** Each phase is independently testable. Do not bundle phases.
- **Copy is final.** Do not rewrite, summarize, or paraphrase any copy. Use the text from `firefly-homepage-copy.md` verbatim.
- **Do not add placeholder content.** If an asset is missing, leave a clearly commented `<!-- TODO: asset needed -->` rather than inserting stock imagery or dummy text.
- **No new dependencies** without flagging first. The stack is Go + html/template + Tailwind v4 + Alpine.js. That's it.
- **Mobile first.** Build each template at 375px first, then layer up with breakpoints at 768px and 1280px.
- **Accessibility baseline.** Semantic HTML throughout. All images get descriptive `alt` attributes. Form fields get associated `<label>` elements. Nav landmark is `<nav>`. Main content is `<main>`. Footer is `<footer>`.

---

*Firefly Software · Helena, Montana · fireflysoftware.dev*
*Brief prepared alongside `firefly-homepage-copy.md` — read both documents before beginning.*
