.PHONY: dev dev-css dev-server build build-css clean

# Run Tailwind watcher and Go server in parallel
dev:
	@echo "Starting dev server..."
	@make dev-css &
	@DEV_MODE=true API_PORT=8080 go run .

# Tailwind CSS watcher
dev-css:
	npx @tailwindcss/cli -i static/css/input.css -o static/css/output.css --watch

# Go server only (no CSS watcher)
dev-server:
	DEV_MODE=true API_PORT=8080 go run .

# Production CSS build
build-css:
	npx @tailwindcss/cli -i static/css/input.css -o static/css/output.css --minify

# Production Go binary
build: build-css
	CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o fs-website .

# Clean build artifacts
clean:
	rm -f fs-website static/css/output.css
