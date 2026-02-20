# =============================================================================
# Stage 1: Build Hugo static site
# =============================================================================
FROM hugomods/hugo:exts AS hugo-builder

WORKDIR /src
COPY . .

RUN hugo --gc --minify --environment production

# =============================================================================
# Stage 2: Build Go API binary
# =============================================================================
FROM golang:1.21-alpine AS go-builder

WORKDIR /build
COPY api/go.mod api/go.sum* ./
RUN go mod download 2>/dev/null || true

COPY api/ .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o fs-api .

# =============================================================================
# Stage 3: Final image — Caddy + static files + Go binary
# =============================================================================
FROM caddy:2-alpine

# Copy Hugo build output to Caddy's web root
COPY --from=hugo-builder /src/public /srv

# Copy Go binary
COPY --from=go-builder /build/fs-api /usr/local/bin/fs-api

# Copy Caddy configuration
COPY Caddyfile /etc/caddy/Caddyfile

# Copy and set entrypoint
COPY docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
RUN chmod +x /usr/local/bin/docker-entrypoint.sh /usr/local/bin/fs-api

# Environment variable defaults (can be overridden at runtime)
ENV PORT=80 \
    API_PORT=8080 \
    ALLOWED_ORIGIN=https://fireflysoftware.dev

EXPOSE 80

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD wget -qO- http://localhost:${PORT:-80}/api/health || exit 1

ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]
