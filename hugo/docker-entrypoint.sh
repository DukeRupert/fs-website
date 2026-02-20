#!/bin/sh
set -e

echo "[entrypoint] Starting Firefly Software API server..."
/usr/local/bin/fs-api &

echo "[entrypoint] Waiting for API to be ready..."
i=0
while [ $i -lt 10 ]; do
  if wget -q -O /dev/null "http://localhost:${API_PORT:-8080}/api/health" 2>/dev/null; then
    echo "[entrypoint] API is ready."
    break
  fi
  i=$((i + 1))
  sleep 1
done

echo "[entrypoint] Starting Caddy..."
exec caddy run --config /etc/caddy/Caddyfile --adapter caddyfile
