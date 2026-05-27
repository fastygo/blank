# Blank

Minimal **Go + templ** app shell on [FastyGo Framework](https://github.com/fastygo/framework) and [github.com/fastygo/templ](https://github.com/fastygo/templ). **This branch (`sidebar`)** ships a classic sidebar + header layout with mobile sheet, dark theme, and En/Ru locale switching — nothing else. Use it as a neutral starting point for a sidebar app. For a hero layout without sidebar, use **`main`**.

## Prerequisites

- Go 1.25+
- [Bun](https://bun.sh) (for CSS build and `ui8px`)

## Quick start

```bash
bun install
go mod download
bun run build:css
go tool templ generate ./...
bun run go
```

Static assets (Tailwind CSS, theme script, `@ui8kit/aria` dialog bundle) live under [`web/static/`](web/static/).

`bun run go` runs [`scripts/run-server.mjs`](scripts/run-server.mjs): the server starts with the **repository root as cwd** and **Ctrl+C** is forwarded to the Go process.

Open [http://127.0.0.1:8080/](http://127.0.0.1:8080/) — home page with sidebar. Second demo route: [http://127.0.0.1:8080/sample](http://127.0.0.1:8080/sample).

## Environment

| Variable | Default | Purpose |
|----------|---------|---------|
| `APP_BIND` | `127.0.0.1:8080` | HTTP listen address |
| `APP_STATIC_DIR` | `web/static` when env omitted | Static files under `/static/` |
| `APP_DEFAULT_LOCALE` | `en` | Default locale |
| `APP_AVAILABLE_LOCALES` | `en,ru` | Locales for the header switcher (query + cookie) |

Probes: `GET /healthz` and `GET /readyz` in [`cmd/server/main.go`](cmd/server/main.go).

## Project layout

| Path | Role |
|------|------|
| [`cmd/server/main.go`](cmd/server/main.go) | Composition root: config, locales, health, site feature |
| [`internal/site/`](internal/site/) | HTTP routes: `/`, `/sample` |
| [`internal/fixtures/locale/`](internal/fixtures/locale/) | Embedded JSON copy per locale |
| [`internal/ui/layout/`](internal/ui/layout/) | Shell, sidebar, header |
| [`internal/ui/components/`](internal/ui/components/) | Icon, language switch |
| [`internal/views/`](internal/views/) | `templ` pages and shell glue |
| [`web/static/`](web/static/) | `app.css`, tokens, fonts, `theme.js`, `ui8kit.js` |

## Verification

```bash
bun run verify
```

Runs: `templ generate` → Tailwind build → `build:js` (dialog-only `@ui8kit/aria`) → `ui8px lint` → `validate:aria` → `go test ./...`.

## Adding a page

1. Add copy to **`fixtures.Locale`** and every **`locale/*.json`** file.
2. Add a nav item in [`internal/site/feature.go`](internal/site/feature.go) (`siteNav`).
3. Add `internal/views/<page>.templ` and a route handler in `internal/site/feature.go`.

For auth, panel navigation, and cabinet routes, use the **`dashboard`** branch as a reference.
