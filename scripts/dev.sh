#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."

if [[ -f .env ]]; then
  set -a
  # shellcheck disable=SC1091
  source .env
  set +a
fi

export APP_BIND="${APP_BIND:-0.0.0.0:3000}"
export APP_STATIC_DIR="${APP_STATIC_DIR:-web/static}"
export APP_DEFAULT_LOCALE="${APP_DEFAULT_LOCALE:-en}"
export APP_AVAILABLE_LOCALES="${APP_AVAILABLE_LOCALES:-en,ru}"
export APP_DEV_OVERLAY="${APP_DEV_OVERLAY:-1}"
export HEALTH_LIVE_PATH="${HEALTH_LIVE_PATH:-/healthz}"
export HEALTH_READY_PATH="${HEALTH_READY_PATH:-/readyz}"
export SESSION_KEY="${SESSION_KEY:-dev-session-key-change-me-32b-min}"

templ generate ./...
go run ./cmd/server
