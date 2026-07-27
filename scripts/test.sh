#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
./scripts/retag-check.sh check
go test ./...
