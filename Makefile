.PHONY: help tidy css js generate build verify run clean deploy install

APP_NAME := blank
BIN_DIR := bin
BINARY := $(BIN_DIR)/$(APP_NAME)

# Default target shows available commands
help:
	@echo "Blank - Makefile targets:"
	@echo ""
	@echo "  tidy       Update go.mod/go.sum and refresh vendor/"
	@echo "  css        Build minified Tailwind CSS (bun)"
	@echo "  js         Build ui8kit JS bundle (bun)"
	@echo "  generate   Run templ code generation"
	@echo "  build      Production build: generate + css + js + go build -mod=vendor"
	@echo "  verify     Full check: templ, css, js, ui8px lint, go test"
	@echo "  run        Start dev server with watch (bun run dev)"
	@echo "  install    Install bun dependencies"
	@echo "  clean      Remove build artifacts (bin/)"
	@echo "  deploy     Full production pipeline (tidy + generate + css + js + build)"
	@echo ""
	@echo "Server deploy example:"
	@echo "  git pull && make deploy"
	@echo "  sudo systemctl restart $(APP_NAME)"

# Ensure Go dependencies and vendor directory are up to date.
# Run this after git pull on the server or when adding new imports.
tidy:
	go mod tidy
	go mod vendor

# Build frontend assets.
css:
	bun run build:css

js:
	bun run build:js

# Generate Go code from .templ files.
# Must run before building if templates changed.
generate:
	bun run templ

# Production binary build.
# Uses vendor/ for reproducible offline builds.
# Strips debug info with -ldflags="-s -w".
build: generate css js
	@mkdir -p $(BIN_DIR)
	go build -mod=vendor -ldflags="-s -w" -o $(BINARY) ./cmd/server

# Run the full verification pipeline (matches package.json "verify").
verify:
	bun run verify

# Start development server (CSS + templ watch, signal handling).
run:
	bun run dev

# Install bun dependencies (run once after clone or package.json changes).
install:
	bun install

# Remove compiled binary and other build outputs.
clean:
	rm -rf $(BIN_DIR)/*

# Production deploy pipeline.
deploy: tidy generate css js build
	@echo ""
	@echo "Deploy build finished."
	@echo "   Binary ready: $(BINARY)"
	@echo ""
	@echo "Next steps on server:"
	@echo "   sudo systemctl restart $(APP_NAME)   # if using systemd"
	@echo "   # or: ./$(BINARY)                   # for manual start"
