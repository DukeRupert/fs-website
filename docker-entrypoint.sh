#!/bin/sh
set -e

# Log in the same JSON-per-line format (sub-second precision is padded with
# zeros: busybox date has no %N) as the Go server and Caddy, so the
# container emits exactly one format on stdout (fleet logging standard §5.2).
log() {
  level="$1"
  msg="$2"
  shift 2
  printf '{"time":"%s","level":"%s","msg":"%s"%s}\n' \
    "$(date -u +%Y-%m-%dT%H:%M:%S).000000000Z" "$level" "$msg" "$*"
}

log INFO "entrypoint starting"
/usr/local/bin/fs-website &

i=0
api_ready=false
while [ $i -lt 10 ]; do
  if wget -q -O /dev/null "http://127.0.0.1:${API_PORT:-8080}/api/health" 2>/dev/null; then
    api_ready=true
    break
  fi
  i=$((i + 1))
  sleep 1
done

if [ "$api_ready" = false ]; then
  log ERROR "server failed to start" ',"attempts":10'
  exit 1
fi

log INFO "server ready" ",\"attempts\":$i"
log INFO "caddy starting"
exec caddy run --config /etc/caddy/Caddyfile --adapter caddyfile
