# Blank

Minimal **Go + templ** app shell on [FastyGo Framework](https://github.com/fastygo/framework) and [github.com/fastygo/templ](https://github.com/fastygo/templ). Top navigation, centered hero welcome, mobile sheet, dark theme, and En/Ru locale switching — nothing else. Use it as a neutral starting point for a new app.

## Prerequisites

- Go 1.25+
- [Bun](https://bun.sh) (for CSS build and `ui8px`)

## Quick start

```bash
bun install
go mod download
bun run dev
```

Open [http://127.0.0.1:8080/](http://127.0.0.1:8080/) — hero welcome page. Second demo route: [http://127.0.0.1:8080/sample](http://127.0.0.1:8080/sample).

`bun run dev` runs [`scripts/dev.mjs`](scripts/dev.mjs): one-shot templ/CSS/JS (+ dev overlay when enabled), then the Go server. For CSS watch, use a second terminal: `bun run watch:css`. **Ctrl+C** stops the server.

Static assets (Tailwind CSS, theme script, `@ui8kit/aria` dialog bundle) live under [`web/static/`](web/static/).

Dev tooling is configured in [`fastygo.config.mjs`](fastygo.config.mjs) (Vite-like central config).

## Scripts

| Command | Purpose |
|---------|---------|
| `bun run dev` | Dev server with CSS + templ watch |
| `bun run start` | One-shot build + `go run` (no watch) |
| `bun run preview` | Same as `start` |
| `bun run build` | Production assets + `go build -o blank` |
| `bun run build:dev-overlay` | Dev-only overlay bundle (`APP_DEV_OVERLAY=1`) |
| `bun run verify` | Full CI-style check |
| `bun run go` | Alias for `dev` |

## For React developers

See [`docs/for-react-devs.md`](docs/for-react-devs.md) for a Vite-to-Blank mental model and dev workflow.

## Environment

| Variable | Default | Purpose |
|----------|---------|---------|
| `APP_BIND` | `127.0.0.1:8080` | HTTP listen address |
| `APP_DEV_OVERLAY` | `1` in local dev via `fastygo.config.mjs` | SSR dev status widget (loopback only) |
| `APP_STATIC_DIR` | `web/static` when env omitted | Static files under `/static/` |
| `APP_DEFAULT_LOCALE` | `en` | Default locale |
| `APP_AVAILABLE_LOCALES` | `en,ru` | Locales for the header switcher (query + cookie) |

Probes: `GET /healthz` and `GET /readyz`.

## Project layout

| Path | Role |
|------|------|
| [`fastygo.config.mjs`](fastygo.config.mjs) | Dev/build tooling config (server, templ, css, js, ui8px) |
| [`cmd/server/main.go`](cmd/server/main.go) | Composition root entry |
| [`internal/serverapp/`](internal/serverapp/) | Framework wiring (locales, security, site feature, dev overlay) |
| [`internal/devoverlay/`](internal/devoverlay/) | Dev-only SSR overlay ([`README`](internal/devoverlay/README.md)) |
| [`internal/devoverlay/fixtures/locale/`](internal/devoverlay/fixtures/locale/) | Dev overlay copy (separate from site fixtures) |
| [`internal/site/`](internal/site/) | HTTP routes: `/`, `/sample` |
| [`internal/fixtures/locale/`](internal/fixtures/locale/) | Site shell and page copy per locale |
| [`internal/ui/`](internal/ui/) | **UI registry** — layout, components, blocks, widgets, variants, utils ([`README`](internal/ui/README.md)) |
| [`internal/ui/layout/`](internal/ui/layout/) | Shell, header nav, footer, mobile sheet |
| [`internal/ui/components/`](internal/ui/components/) | Icon, language switch |
| [`internal/views/`](internal/views/) | `templ` pages and shell glue |
| [`web/static/`](web/static/) | `app.css`, tokens, fonts, `theme.js`, `ui8kit.js` |

## Dev overlay

When `APP_DEV_OVERLAY=1` on loopback, Blank injects a small dev widget with three tabs:

- **Health** — `/healthz` and `/readyz` probe status and latency
- **Assets** — `app.css`, `ui8kit.js`, `theme.js` freshness
- **Request** — page `X-Request-ID`, path, and `<html lang>`

Use **Hide overlay** to set an opt-out cookie and reload. The next SSR response contains no overlay markup or script tags.

Build the overlay bundle with `bun run build:dev-overlay` (also runs during `bun run dev` when overlay is enabled).

Overlay strings live in [`internal/devoverlay/fixtures/locale/`](internal/devoverlay/fixtures/locale/). Portability notes: [`internal/devoverlay/README.md`](internal/devoverlay/README.md).

## Verification

```bash
bun run verify
```

Runs: `templ generate` → Tailwind build → `build:js` → `ui8px lint` → `validate:aria` → `go test ./...`.

## Adding a page

1. Add copy to **`fixtures.Locale`** and every **`locale/*.json`** file.
2. Add a nav item in [`internal/site/feature.go`](internal/site/feature.go) (`siteNav`).
3. Add `internal/views/<page>.templ` and a route handler in `internal/site/feature.go`.

For the previous sidebar layout, use the **`sidebar`** branch as a reference.
