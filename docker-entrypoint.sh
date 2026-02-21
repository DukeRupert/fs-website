#!/bin/sh
set -e

echo "[entrypoint] Starting Firefly Software API server..."
/usr/local/bin/fs-api &

echo "[entrypoint] Waiting for API to be ready..."
i=0
api_ready=false
while [ $i -lt 10 ]; do
  if wget -q -O /dev/null "http://127.0.0.1:${API_PORT:-8080}/api/health" 2>/dev/null; then
    echo "[entrypoint] API is ready."
    api_ready=true
    break
  fi
  i=$((i + 1))
  sleep 1
done

if [ "$api_ready" = false ]; then
  echo "[entrypoint] ERROR: API failed to start after 10 attempts. Exiting."
  exit 1
fi

echo "[entrypoint] Starting Caddy..."
exec caddy run --config /etc/caddy/Caddyfile --adapter caddyfile
