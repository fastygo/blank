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

`bun run dev` runs [`scripts/dev.mjs`](scripts/dev.mjs): initial templ/CSS/JS build, Tailwind and templ watchers, then the Go server. **Ctrl+C** stops all child processes.

Static assets (Tailwind CSS, theme script, `@ui8kit/aria` dialog bundle) live under [`web/static/`](web/static/).

Dev tooling is configured in [`fastygo.config.mjs`](fastygo.config.mjs) (Vite-like central config).

## Scripts

| Command | Purpose |
|---------|---------|
| `bun run dev` | Dev server with CSS + templ watch |
| `bun run start` | One-shot build + `go run` (no watch) |
| `bun run preview` | Same as `start` |
| `bun run build` | Production assets + `go build -o blank` |
| `bun run verify` | Full CI-style check |
| `bun run go` | Alias for `dev` |

## For React developers

See [`docs/for-react-devs.md`](docs/for-react-devs.md) for a Vite-to-Blank mental model and dev workflow.

## Environment

| Variable | Default | Purpose |
|----------|---------|---------|
| `APP_BIND` | `127.0.0.1:8080` | HTTP listen address |
| `APP_STATIC_DIR` | `web/static` when env omitted | Static files under `/static/` |
| `APP_DEFAULT_LOCALE` | `en` | Default locale |
| `APP_AVAILABLE_LOCALES` | `en,ru` | Locales for the header switcher (query + cookie) |

Probes: `GET /healthz` and `GET /readyz`.

## Project layout

| Path | Role |
|------|------|
| [`fastygo.config.mjs`](fastygo.config.mjs) | Dev/build tooling config (server, templ, css, js, ui8px) |
| [`cmd/server/main.go`](cmd/server/main.go) | Composition root entry |
| [`internal/serverapp/`](internal/serverapp/) | Framework wiring (locales, security, site feature) |
| [`internal/site/`](internal/site/) | HTTP routes: `/`, `/sample` |
| [`internal/fixtures/locale/`](internal/fixtures/locale/) | Embedded JSON copy per locale |
| [`internal/ui/`](internal/ui/) | **UI registry** — layout, components, blocks, widgets, variants, utils ([`README`](internal/ui/README.md)) |
| [`internal/ui/layout/`](internal/ui/layout/) | Shell, header nav, footer, mobile sheet |
| [`internal/ui/components/`](internal/ui/components/) | Icon, language switch |
| [`internal/views/`](internal/views/) | `templ` pages and shell glue |
| [`web/static/`](web/static/) | `app.css`, tokens, fonts, `theme.js`, `ui8kit.js` |

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
