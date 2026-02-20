# Portfolio Image Crop Guide

How to prepare screenshots for case study pages. All commands use ImageMagick (`magick`).

## Source Screenshots

Take full-page screenshots of each site at these viewport widths:

| Type    | Viewport Width | Tool                                          |
|---------|---------------|-----------------------------------------------|
| Desktop | ~2195px (2x)  | Chrome DevTools > Capture full size screenshot |
| Mobile  | iPhone 14 Pro Max (1290px logical) | Chrome DevTools > Device toolbar > Capture full size screenshot |

## Crop Specs

### Desktop Splash (`*_splash.png`)

Crop from the top to capture the hero/above-the-fold area, then resize to 800px wide.

- **Source crop**: `2195x1370` (roughly 16:10 ratio)
- **Final size**: `800x499`

```bash
magick input.png -crop 2195x1370+0+0 +repage -resize 800x static/images/portfolio/PREFIX_splash.png
```

### Mobile Screenshot (`*_mobile.png`)

Crop from the top to one full iPhone screen height (19.5:9 ratio), then resize to 400px wide.

- **Source crop**: `1290x2796` (iPhone 14 Pro Max ratio)
- **Final size**: `400x867`

```bash
magick input.png -crop 1290x2796+0+0 +repage -resize 400x static/images/portfolio/PREFIX_mobile.png
```

## File Naming

Use a short lowercase prefix per project:

| Project                        | Prefix     |
|-------------------------------|------------|
| Traver Hardwood Floors        | `thf`      |
| Christian Brother's Lining Co.| `cbl`      |
| 406 Records                   | `406`      |
| Momentum Business Solutions   | `mbs`      |
| Nautilus Group Cleaning       | `nautilus`  |

Each project gets three image files:

```
static/images/portfolio/PREFIX_splash.png    # Desktop hero crop
static/images/portfolio/PREFIX_mobile.png    # Mobile screen crop
static/images/portfolio/PREFIX_logo.png      # Client logo (as-is from client)
```

## CSS Display Sizes

These are the max display sizes in the template, for reference:

| Image   | Max Display Width | Context                          |
|---------|------------------|----------------------------------|
| Mobile  | 260px (200px on small screens) | Hero right column, rounded corners (32px) |
| Desktop | 100% of narrow container       | Below case study content         |
| Logo    | 60px height max  | Hero left column, granite background, 4px radius |

## Quick Reference: Full Project Setup

```bash
# 1. Crop desktop
magick full-page-desktop.png -crop 2195x1370+0+0 +repage -resize 800x static/images/portfolio/PREFIX_splash.png

# 2. Crop mobile
magick full-page-mobile.png -crop 1290x2796+0+0 +repage -resize 400x static/images/portfolio/PREFIX_mobile.png

# 3. Copy logo
cp logo.png static/images/portfolio/PREFIX_logo.png

# 4. Add entry to data/portfolio.yaml
# 5. Create content/portfolio/SLUG.md
```
