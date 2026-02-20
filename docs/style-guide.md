# Firefly Software — UI Style Guide
**Direction:** Montana Utility × Wilderness Wonder  
**Version:** 1.0  
**Stack:** Go · HTMX · Alpine.js · Tailwind CSS

---

## Philosophy

The visual language is Montana utility at its bones — sharp corners, honest type, no decoration for its own sake — with a layer of wilderness wonder on top. Think: the moment before sunrise when the sky is still dark but the horizon is starting to glow amber over the Sleeping Giant ridge. Structural decisions are grounded and practical. The wonder comes from light, color, and atmosphere — not from ornament.

**The logo** is a stylized firefly with a lightbulb as its abdomen. White lines on a dark navy background. This is the connective thread: a creature that lives at the intersection of darkness and unexpected warm light.

**The one thing to remember:** Every dark surface is midnight navy, not black. Every warm accent traces back to ember and amber — sunrise, not neon.

---

## Color Tokens

Define these as CSS custom properties on `:root`. Reference them everywhere — never hardcode hex values in components.

```css
:root {
  /* ── Light surfaces ─────────────────────────── */
  --granite:      #F0EDE8;   /* page background */
  --granite-dk:   #DDD8CE;   /* borders, dividers on light bg */
  --granite-lt:   #F9F7F4;   /* card hover, input focus bg */
  --snow:         #FAFAF8;   /* card surfaces, canvas areas */

  /* ── Dark surfaces — midnight navy, never pure black ── */
  --midnight:     #0A0E18;   /* hero sky, deepest dark */
  --navy:         #0F1628;   /* nav, footer, dark sections */
  --navy-mid:     #162038;   /* stat cards, dark section bg */
  --navy-edge:    #1E2D4A;   /* card surfaces within dark sections */

  /* ── Ink — warm, not cool ───────────────────── */
  --ink:          #1A1A18;   /* body text, headings on light bg */
  --ink-mid:      #3D3D38;   /* secondary text */
  --ink-faded:    #6B6860;   /* captions, placeholders */
  --stone:        #8A8478;   /* metadata, muted labels */

  /* ── Montana accents ────────────────────────── */
  --rust:         #8B3A1A;   /* primary action color */
  --rust-lt:      #A84620;   /* rust hover state */
  --ember:        #C4581A;   /* warmer rust — sunrise-adjacent, CTA hover */
  --pine:         #2C4A2E;   /* secondary action, success */
  --pine-lt:      #3D6640;   /* pine hover */
  --glacier:      #4A7C8E;   /* info, glacier blue */

  /* ── Wonder accents ─────────────────────────── */
  --amber:        #E8922A;   /* sunrise glow, warm highlights */
  --amber-lt:     #F0A840;   /* horizon light, nav underline */
  --amber-dim:    rgba(232, 146, 42, 0.15);  /* ambient glow fills */
  --amber-glow:   rgba(232, 146, 42, 0.08);  /* subtle card hover shadow */
  --firefly:      #B8E850;   /* bioluminescent green — use sparingly */
  --firefly-dim:  rgba(184, 232, 80, 0.12);  /* firefly ambient glow */
  --star:         #D8E4F0;   /* cool starlight — star dots only */
}
```

### Color Usage Rules

| Context | Token |
|---|---|
| Page background | `--granite` |
| Card / canvas surface (light) | `--snow` |
| Card hover (light) | `--granite-lt` |
| Nav, footer | `--navy` |
| Dark section background | `--navy-mid` |
| Dark card surface | `--navy-edge` |
| Primary button, rust accents | `--rust` |
| Primary button hover | `--ember` |
| Secondary / success | `--pine` |
| Info / glacier | `--glacier` |
| Warm wonder accent | `--amber` |
| Horizon / hero kicker | `--amber-lt` |
| Firefly glow (sparingly) | `--firefly` |
| Body text | `--ink` |
| Secondary text | `--ink-mid` |
| Muted / captions | `--ink-faded` |
| Labels, metadata | `--stone` |

**Never use `#000000` or `#ffffff` directly.** Use `--midnight`/`--navy` for darks, `--snow`/`--granite` for lights.

---

## Typography

### Font Stack

```css
--font-display: 'Playfair Display', serif;   /* hero titles, section headings */
--font-body:    'Libre Baskerville', serif;   /* all body copy */
--font-ui:      'Oswald', sans-serif;         /* buttons, labels, nav, badges */
--font-mono:    'Source Code Pro', monospace; /* metadata, coordinates, code */
```

Load from Google Fonts:
```html
<link href="https://fonts.googleapis.com/css2?family=Playfair+Display:ital,wght@0,700;0,900;1,700&family=Libre+Baskerville:ital,wght@0,400;0,700;1,400&family=Oswald:wght@400;500;600;700&family=Source+Code+Pro:wght@400;600&display=swap" rel="stylesheet">
```

### Type Scale

| Role | Font | Size | Weight | Transform | Spacing |
|---|---|---|---|---|---|
| Hero title | Playfair Display | `clamp(42px, 5.5vw, 78px)` | 900 | — | `-0.01em` |
| Section title | Playfair Display | `clamp(28px, 3.5vw, 46px)` | 900 | — | `-0.01em` |
| Card title | Oswald | `16px` | 700 | uppercase | `0.06em` |
| Nav links | Oswald | `12px` | 500 | uppercase | `0.20em` |
| Buttons | Oswald | `11–13px` | 700 | uppercase | `0.25em` |
| Section labels | Oswald | `10px` | 600 | uppercase | `0.40em` |
| Badge / tag | Oswald | `10px` | 600 | uppercase | `0.25em` |
| Body copy | Libre Baskerville | `15–16px` | 400 | — | — |
| Captions | Libre Baskerville | `14px` | 400 | — | — |
| Italic sub-copy | Libre Baskerville | `15px` | 400 | italic | — |
| Metadata / coords | Source Code Pro | `10–11px` | 400 | — | `0.10–0.15em` |

### Typography Rules

- Line height: `1.75` for body, `1.0–1.1` for display titles
- **No Inter, Roboto, Arial, or system-ui** anywhere in the UI
- Playfair italic (`<em>`) in hero titles can receive a warm gradient: `linear-gradient(135deg, var(--amber-lt), var(--snow))`
- Hero kicker lines use `--font-ui`, `10–11px`, `0.42em` tracking, `var(--amber-lt)` color, with a `20px` rule prefix

---

## Spacing & Layout

```
Max content width:   1100px (centered, auto margins)
Section padding:     88px 40px (desktop) / 60px 20px (mobile)
Section gap:         72px between major sections
Card grid gap:       2px (rendered as background color of grid container)
```

### Grid System

Use CSS Grid. The `2px` gap pattern is a signature detail — set `gap: 2px` on the grid container and match `background` to `--granite-dk` (light) or `--midnight` (dark). Cards fill the rest.

```css
.card-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 2px;
  background: var(--granite-dk); /* the "gap" is the background showing through */
}
.card {
  background: var(--snow);
}
```

---

## Shape & Decoration

### Border Radius
**Zero everywhere.** No `border-radius` on any UI element — cards, buttons, inputs, badges, modals, dropdowns. The only exception is circular elements (logo glow dot, progress pips, firefly dots).

### Borders
- Light section dividers: `1px solid var(--granite-dk)`
- Dark section dividers: `1px solid rgba(255,255,255,0.06)`
- Accent left-borders on cards/alerts: `2–3px solid` using the semantic color
- Nav bottom border: `2px solid var(--ember)`
- Dark section top border: `1px solid rgba(232,146,42,0.2)`

### Grain Texture
Apply a subtle noise grain to the entire page via `body::after`:

```css
body::after {
  content: '';
  position: fixed;
  inset: 0;
  pointer-events: none;
  z-index: 9999;
  opacity: 0.025;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='300' height='300'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.85' numOctaves='4' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='300' height='300' filter='url(%23n)'/%3E%3C/svg%3E");
}
```

---

## Components

### Navigation

```
Height:          66px
Background:      --navy
Bottom border:   2px solid --ember
Brand font:      --font-display, 15px, weight 700
Link font:       --font-ui, 12px, weight 500, uppercase, 0.20em tracking
Link color:      rgba(250,250,248,0.45) → --snow on hover/active
Link hover:      bottom 2px gradient bar: linear-gradient(90deg, --ember, --amber)
CTA background:  --rust → --ember on hover
Aurora shimmer:  subtle linear-gradient across nav (see implementation below)
```

```css
/* Aurora shimmer — apply as nav::before */
background: linear-gradient(
  90deg,
  transparent 0%,
  rgba(232,146,42,0.04) 30%,
  rgba(184,232,80,0.03) 60%,
  transparent 100%
);
```

The brand mark includes an inline SVG firefly icon with a pulsing `radial-gradient` glow behind it:

```css
@keyframes firefly-pulse {
  0%, 100% { opacity: 1; transform: scale(1); }
  50%       { opacity: 0.4; transform: scale(0.85); }
}
/* animation: firefly-pulse 2.8s ease-in-out infinite */
```

---

### Buttons

All buttons: `font-family: --font-ui`, `font-weight: 700`, `uppercase`, `letter-spacing: 0.25em`, `border-radius: 0`.

| Variant | Background | Color | Border | Hover |
|---|---|---|---|---|
| Primary | `--rust` | `--snow` | `2px solid --rust` | bg → `--ember`, `translateY(-1px)`, rust shadow |
| Secondary (pine) | `--pine` | `--snow` | `2px solid --pine` | bg → `--pine-lt`, `translateY(-1px)` |
| Amber outline | transparent | `--amber` | `2px solid rgba(232,146,42,0.4)` | border → `--amber`, amber glow shadow |
| Ghost | transparent | `--ink-faded` | `2px solid --granite-dk` | border → `--stone` |
| Outline ink | transparent | `--ink` | `2px solid --ink` | bg → `--ink`, color → `--snow` |
| Outline snow | transparent | `--snow` | `2px solid rgba(250,250,248,0.3)` | border → `--snow` |

```
Sizes:
  Small:   padding 9px 18px,  font-size 11px
  Default: padding 14px 28px, font-size 12px
  Large:   padding 17px 34px, font-size 13px

Hover transition: all 0.15s ease
Hover lift:       translateY(-1px)
Primary shadow:   0 4px 20px rgba(196,88,26,0.35)
```

---

### Cards (Light Mode)

```
Background:        --snow
Hover background:  --granite-lt
Transition:        background 0.2s
Top bar on hover:  2px gradient, scaleX(0→1) from left, 0.25s ease
Top bar gradient:  linear-gradient(90deg, --rust, --amber)
```

Card number: `--font-mono`, `10px`, `--stone`  
Card title: `--font-ui`, `16px`, weight 700, uppercase, `0.06em` tracking, `--ink`  
Card body: `--font-body`, `14px`, `--ink-faded`, line-height `1.65`  
Card link: `--font-ui`, `11px`, weight 600, uppercase, `--rust`, border-bottom on hover

```css
.card::after {
  content: '';
  position: absolute;
  top: 0; left: 0; right: 0;
  height: 2px;
  background: linear-gradient(90deg, var(--rust), var(--amber));
  transform: scaleX(0);
  transform-origin: left;
  transition: transform 0.25s ease;
}
.card:hover::after { transform: scaleX(1); }
```

---

### Stat Cards (Dark Mode)

```
Grid background:   --midnight
Card background:   --navy-mid
Card hover:        --navy-edge
Top bar:           24px wide, 2px tall, linear-gradient(90deg, --ember, --amber)
Value font:        --font-display, 50px, weight 900, --snow
Value sup:         0.4em, superscript, --amber
Label font:        --font-ui, 10px, weight 600, uppercase, 0.3em tracking, rgba(250,250,248,0.4)
Note font:         --font-body, 11px, italic, rgba(250,250,248,0.25)
```

---

### Section Labels (Eyebrows)

Used above every major section. Pattern: `01 — Label Name`

```css
.section-label {
  font-family: var(--font-ui);
  font-size: 10px;
  font-weight: 600;
  letter-spacing: 0.4em;
  text-transform: uppercase;
  color: var(--stone);
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 18px;
}
.section-label::after {
  content: '';
  flex: 1;
  height: 1px;
  background: var(--granite-dk);
}
```

---

### Form Elements

```
Input background:   --granite
Input border:       1px solid --granite-dk  +  2px bottom solid --stone
Input focus:        border-bottom → --ember, background → --granite-lt
Input font:         --font-body, 15px, --ink
Placeholder:        --stone, italic, 14px
Label font:         --font-ui, 10px, weight 600, uppercase, 0.3em tracking, --ink-mid
Textarea:           min-height 90px, resize vertical
```

Never use `border-radius` on inputs. No custom checkbox/radio styling beyond what's needed for clarity.

---

### Badges / Tags

```
Font:        --font-ui, 10px, weight 600, uppercase, 0.25em tracking
Border:      1px solid currentColor
Padding:     4px 10px
Radius:      0
```

| Variant | Color |
|---|---|
| `badge-rust` | `--rust` outline |
| `badge-pine` | `--pine` outline |
| `badge-amber` | `--amber` outline |
| `badge-glacier` | `--glacier` outline |
| `badge-stone` | `--stone` outline |
| `badge-solid-rust` | `--rust` bg, `--snow` text |
| `badge-solid-pine` | `--pine` bg, `--snow` text |
| `badge-solid-amber` | `--amber` bg, `--ink` text |

---

### Alerts

```
Structure:    left border (3px, semantic color) + 3-side border (1px --granite-dk) + icon + content
Background:   --snow
Title font:   --font-ui, 12px, weight 600, uppercase, 0.2em tracking
Body font:    --font-body, 14px, --ink-faded, line-height 1.55
```

| Variant | Color token |
|---|---|
| Info | `--glacier` |
| Success | `--pine` |
| Warning | `--amber` |
| Error | `--rust` |

---

### Progress Bars

```
Track:       3px height, --granite-dk background
Fill:        height 100%, gradient or solid per variant
Pip marker:  8px circle at fill endpoint, colored glow (box-shadow)
```

```css
/* Amber fill */
background: linear-gradient(90deg, var(--rust), var(--amber));

/* Pip glow */
.progress-fill::after {
  content: '';
  position: absolute;
  right: -1px; top: 50%;
  transform: translateY(-50%);
  width: 8px; height: 8px;
  border-radius: 50%;
  background: var(--amber-lt);
  box-shadow: 0 0 6px 2px var(--amber-dim), 0 0 2px var(--amber-lt);
}
```

---

### Tagline Strip

Full-width dark band used between sections.

```
Background:   --navy
Left accent:  3px vertical bar, linear-gradient(180deg, --amber, --rust)
Top/bottom:   1px solid rgba(232,146,42,0.2)
Horizon glow: linear-gradient(90deg, rgba(232,146,42,0.06), transparent) on left side
Text font:    --font-display, 17px, weight 700, italic, --snow
Meta font:    --font-mono, 10px, rgba(250,250,248,0.25)
```

---

## Hero SVG Landscape

The hero background is a hand-crafted SVG. No photography. No external image CDNs.

### Sky Gradient (Night/Dusk)
```
0%   #050810  (near-black midnight)
35%  #080D1C
65%  #0E1628
80%  #1A2030
90%  #2A2A18  (starts warming)
100% #3A2A10  (amber at the ridge)
```

### Horizon Glow
Two overlapping `radialGradient` elements centered at `cy="100%"`:
- Primary: `#D05818` → transparent, centered at 50%, radius 65%
- Secondary: `#E07020` → transparent, centered at 38%, radius 45%

### The Sleeping Giant Silhouette
A single SVG `<path>` tracing the Beartooth Mountain ridgeline as seen from the Helena Valley. Head to the right (east), feet to the left (west). Key landmarks reading left to right:

| Feature | Approx X | Approx Y |
|---|---|---|
| Feet (left edge) | 30–318 | 340–410 |
| Belly (flat stretch) | 595–706 | 237–247 |
| Chest peak | 765 | 184 |
| Sternum dip | 792–815 | 220–246 |
| Nose tip *(most prominent)* | 883 | 153 |
| Brow saddle | 901–927 | 194–210 |
| Forehead crown | 959 | 175 |
| Back of head | 987–1081 | 190–281 |

Apply a rim-light stroke along the mountain's top edge using the sunrise gradient — this is the sunrise catching the ridge:
```svg
<path stroke="url(#sunrise)" stroke-width="2.5" fill="none" opacity="0.6" d="[ridge top path]"/>
```

### Stars
Hand-placed `<circle>` elements, three layers:
- **Bright:** `r="1.4–1.8"`, `opacity="0.85–0.95"`
- **Medium:** `r="1.0–1.2"`, `opacity="0.70–0.80"`
- **Dim (depth):** `r="0.6–0.8"`, `opacity="0.45–0.60"`

Fill: `#D8E4F0` (cool starlight, not white)

### Moon
`<circle>` at upper-right sky. `r="22"`, fill `#F0E8D0`. Soft glow ring at `r="40"`, opacity `0.06`, with `feGaussianBlur`.

### Firefly Glows (near treeline)
Small `<circle>` elements, `r="2–3.5"`, fill `#B8E850`, with `feGaussianBlur` filter for soft glow. Keep to 6–10 total. These sit just above the valley floor near the treeline.

### Pine Silhouettes
Simple `<polygon>` triangles. Fill `#080C08`. Clustered in groups of 3–6 at left edge, scattered center, clustered at right edge.

### Hero Overlay Gradient
```css
background: linear-gradient(
  to top,
  rgba(10,14,24,0.97) 0%,   /* near-opaque at bottom for text */
  rgba(10,14,24,0.60) 35%,
  rgba(10,14,24,0.15) 65%,
  rgba(10,14,24,0.02) 100%  /* transparent at top — show the sky */
);
```

---

## Dark Sections

Used for "Why Us" and similar high-emphasis sections. Layered construction:

```
Background:   --midnight
Stars:        CSS radial-gradient star field (two layers, bright + dim)
Fireflies:    Absolutely positioned divs with blink animation, --firefly color
Horizon glow: Bottom 160px, linear-gradient(to top, rgba(196,88,26,0.35), transparent)
Ridge:        SVG silhouette at bottom of section, fill #080C08
Content:      z-index 2, above all atmospheric layers
```

### Star Field (CSS Only)
```css
.star-layer {
  position: absolute;
  inset: 0;
  background-image:
    radial-gradient(1.5px 1.5px at 27% 8%, #D8E4F0 0%, transparent 100%),
    radial-gradient(1px 1px at 8% 12%, #D8E4F0 0%, transparent 100%),
    /* ... 20–30 total points ... */;
  opacity: 0.75;
}
/* Second layer at opacity 0.45 for depth */
```

### Firefly Animation
```css
@keyframes blink {
  0%, 100% { opacity: 0; transform: scale(0.6); }
  40%, 60%  { opacity: 1; transform: scale(1); }
}
.firefly {
  position: absolute;
  border-radius: 50%;
  background: var(--firefly);
  box-shadow: 0 0 8px 3px var(--firefly-dim);
  animation: blink var(--d, 3s) ease-in-out var(--delay, 0s) infinite;
}
```

Place 6–10 fireflies at varied positions across the lower 40% of the section. Vary `--d` (duration) between `2.4s–4.1s` and `--delay` between `0s–2.5s` for organic feel.

### Dark Section Cards (Night Items)
```
Background:    rgba(255,255,255,0.03)
Border:        1px solid rgba(255,255,255,0.06)
Left accent:   2px solid --ember → --amber on hover
Hover bg:      rgba(255,255,255,0.05)
Title font:    --font-ui, 13px, weight 700, uppercase, 0.12em tracking, --snow
Body font:     --font-body, 13px, rgba(250,250,248,0.5), line-height 1.55
Icon color:    --amber
```

---

## Animations & Motion

Keep motion purposeful and subtle. The wonder is in stillness with occasional light.

| Element | Animation | Duration | Easing |
|---|---|---|---|
| Firefly logo glow | `firefly-pulse` (scale + opacity) | 2.8s | ease-in-out, infinite |
| Card top bar | `scaleX(0→1)` on hover | 0.25s | ease |
| Firefly field dots | `blink` (opacity + scale) | 2.4–4.1s | ease-in-out, infinite, staggered |
| Button hover | `translateY(-1px)` | 0.15s | ease |
| Nav link underline | `scaleX(0→1)` | 0.22s | ease |
| All color transitions | `background`, `color`, `border-color` | 0.15–0.20s | ease |

**No bounce, no spring, no slide-in-from-left.** Motion is either a fade, a scale, or a translate of 1–2px maximum.

---

## Icons

Use inline SVG only. No icon font libraries. No external icon CDNs.

```
Stroke width:  1.5 (default) or 2 (emphasis)
Fill:          none (stroke-only)
Size:          24×24px viewBox, scale via width/height attributes
Color:         inherit via currentColor or explicit token
```

---

## What Never to Do

| Rule | Detail |
|---|---|
| No `border-radius` | Zero on all UI elements except circles |
| No pure black | Use `--midnight` or `--navy` for darks |
| No pure white | Use `--snow` or `--granite-lt` |
| No Inter, Roboto, Arial | Use the four defined fonts only |
| No external background images | Inline SVG only for atmospheric elements |
| No purple or blue gradients | Warm amber/ember/rust only for gradients |
| No heavy shadows | Max: `0 4px 20px rgba(196,88,26,0.35)` |
| No content changes | Copy, prices, links, structure — preserve exactly |
| No icon packs | Inline SVG, hand-drawn, stroke-only |
| No scattered micro-animations | Limit to: nav hover, button hover, card top bar, fireflies |
| No bright/neon colors | `--firefly` is the limit — and only for the logo glow and firefly dots |

---

## Go Template Preservation

This codebase uses Go's `html/template`. When converting pages, **never modify template tags.** Treat `{{ }}` blocks as opaque — preserve them exactly including whitespace inside the delimiters.

```html
<!-- ✅ Correct — template tag untouched -->
<a href="{{ .URL }}">{{ .Title }}</a>

<!-- ❌ Wrong — template tag altered -->
<a href="{{.URL}}">{{.Title}}</a>
```

---

## File Reference

| File | Purpose |
|---|---|
| `firefly-wonder.html` | Full component showcase — canonical reference for all patterns |
| `firefly-landing.html` | Full landing page with real business content |
| `firefly-sleeping-giant.html` | Isolated hero SVG test |

The component showcase (`firefly-wonder.html`) is the ground truth. When in doubt about a spacing, color, or type decision, defer to what's rendered there.

---

*Firefly Software · Helena, Montana · 46.8°N 111.9°W*