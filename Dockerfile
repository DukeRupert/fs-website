# =============================================================================
# Stage 1: Build Tailwind CSS
# =============================================================================
FROM node:22-alpine AS css-builder

WORKDIR /build
COPY package.json ./
RUN npm install

COPY static/css/ static/css/
COPY templates/ templates/
RUN npx @tailwindcss/cli -i static/css/input.css -o static/css/output.css --minify

# =============================================================================
# Stage 2: Build Go server binary
# =============================================================================
FROM golang:1.23-alpine AS go-builder

WORKDIR /build
COPY go.mod go.sum* ./
RUN go mod download 2>/dev/null || true

COPY main.go .
COPY handlers/ handlers/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o fs-website .

# =============================================================================
# Stage 3: Final image — Caddy + Go binary + templates + static assets
# =============================================================================
FROM caddy:2-alpine

# Copy Go binary
COPY --from=go-builder /build/fs-website /usr/local/bin/fs-website

# Copy templates
COPY templates/ /app/templates/

# Copy static assets
COPY static/ /app/static/

# Copy built CSS (overwrite the input.css with output.css)
COPY --from=css-builder /build/static/css/output.css /app/static/css/output.css

# Copy content (blog posts)
COPY content/ /app/content/

# Copy Caddy configuration
COPY Caddyfile /etc/caddy/Caddyfile

# Copy and set entrypoint
COPY docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
RUN chmod +x /usr/local/bin/docker-entrypoint.sh /usr/local/bin/fs-website

WORKDIR /app

# Environment variable defaults (can be overridden at runtime)
ENV PORT=80 \
    API_PORT=8080 \
    ALLOWED_ORIGIN=https://fireflysoftware.dev

EXPOSE 80

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD wget -qO- http://127.0.0.1:${PORT:-80}/api/health || exit 1

ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]
