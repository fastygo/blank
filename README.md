# Blank

Official **Go + templ** starter shell for the [FastyGo](https://github.com/fastygo) stack: [FastyGo Framework](https://github.com/fastygo/framework) + [`github.com/fastygo/templ`](https://github.com/fastygo/templ). Clone it as a neutral base for a new server-rendered app — no legacy UI8Kit/Elements/Blocks stack.

**Module:** `github.com/fastygo/blank`

Demo routes: topnav home (`/`) and sidebar dashboard sample (`/sample`), with mobile sheet, dark theme, and En/Ru locale switching. Each page composes its own layout layers in `internal/views/`; routing lives in `internal/site/router.go`.

## What this repo is

| Layer | Location | Role |
|-------|----------|------|
| Framework wiring | [`internal/serverapp/`](internal/serverapp/), [`cmd/server/`](cmd/server/) | Locales, security, site feature, dev overlay |
| UI registry | [`internal/ui/`](internal/ui/) | Layout shells, components, blocks (staging tree aligned with FastyGoUI) |
| Pages | [`internal/views/`](internal/views/) | Route bodies — `views.<Page>(d, f)` with chosen layout layers |
| Copy & i18n | [`internal/fixtures/locale/`](internal/fixtures/locale/) | Embedded En/Ru strings |
| Dev tooling | [`fastygo.config.mjs`](fastygo.config.mjs), [`scripts/`](scripts/) | templ/CSS/JS build, ui8px — not routes or layouts |

## Prerequisites

- Go 1.25+
- [Bun](https://bun.sh) (for CSS build and `ui8px`)

## Quick start

```bash
bun install
go mod download
bun run dev
```

Open [http://127.0.0.1:8080/](http://127.0.0.1:8080/) — hero welcome page (composes `layout.RootLayout` + `layout.TopnavLayout`). Second demo route: [http://127.0.0.1:8080/sample](http://127.0.0.1:8080/sample) composes `layout.RootLayout` + `layout.DashboardLayout` directly in the page template.

`bun run dev` runs [`scripts/dev.mjs`](scripts/dev.mjs): one-shot templ/CSS/JS (+ dev overlay when enabled), then the Go server. After `.templ` edits run `bun run templ`. After Tailwind class changes run `bun run build:css`. After Go edits, restart dev (`Ctrl+C`, then `bun run dev`). **Ctrl+C** stops the server.

Static assets (Tailwind CSS, theme script, `@ui8kit/aria` dialog bundle, and the dev overlay bundle) live under [`web/static/`](web/static/) and [`internal/devoverlay/static/`](internal/devoverlay/static/).

Dev tooling is configured in [`fastygo.config.mjs`](fastygo.config.mjs) — **tooling only** (server env, templ/CSS/JS build, ui8px). Routes live in [`internal/site/router.go`](internal/site/router.go); each page composes its own layout layers inside [`internal/views/<page>.templ`](internal/views/) — there are no layout adapters.

## Scripts

| Command | Purpose |
|---------|---------|
| `bun run dev` | One-shot templ/CSS/JS build + Go server |
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
| `APP_DEV_OVERLAY` | `1` in local dev via `fastygo.config.mjs` | Dev status overlay loader (loopback only) |
| `APP_STATIC_DIR` | `web/static` when env omitted | Static files under `/static/` |
| `APP_DEFAULT_LOCALE` | `en` | Default locale |
| `APP_AVAILABLE_LOCALES` | `en,ru` | Locales for the header switcher (query + cookie) |

Probes: `GET /healthz` and `GET /readyz`.

## Project layout

| Path | Role |
|------|------|
| [`fastygo.config.mjs`](fastygo.config.mjs) | Dev/build tooling only (server, templ, css, js, ui8px) — not routes or layouts |
| [`cmd/server/main.go`](cmd/server/main.go) | Composition root entry |
| [`internal/serverapp/`](internal/serverapp/) | Framework wiring (locales, security, site feature, dev overlay) |
| [`internal/devoverlay/`](internal/devoverlay/) | Dev-only overlay loader, routes, embedded Shadow DOM web component ([`README`](internal/devoverlay/README.md)) |
| [`internal/devoverlay/fixtures/locale/`](internal/devoverlay/fixtures/locale/) | Dev overlay copy (separate from site fixtures) |
| [`internal/site/`](internal/site/) | Runtime route manifest (`router.go`), render helpers, feature wiring |
| [`internal/fixtures/locale/`](internal/fixtures/locale/) | Site shell and page copy per locale |
| [`internal/ui/`](internal/ui/) | **UI registry** — layout shells, components, blocks, widgets, variants, utils ([`README`](internal/ui/README.md)) |
| [`internal/ui/layout/`](internal/ui/layout/) | Layout layers (`RootLayout`, `TopnavLayout`, `DashboardLayout`), header, footer, mobile sheet |
| [`internal/ui/components/`](internal/ui/components/) | Icon, language switch, navigation, **appsidebar** (local aside) |
| [`internal/views/`](internal/views/) | `templ` pages — each composes `views.<Page>(d, f)` with its own layout layers |
| [`web/static/`](web/static/) | `app.css`, tokens, fonts, `theme.js`, `ui8kit.js` |

## Dev overlay

When `APP_DEV_OVERLAY=1` on loopback, Blank injects a viewport-gated loader script. On desktop it loads an isolated `<fastygo-dev-overlay>` web component with Shadow DOM and three tabs:

- **Health** — `/healthz` and `/readyz` probe status and latency
- **Assets** — `app.css`, `ui8kit.js`, `theme.js` freshness
- **Request** — page `X-Request-ID`, path, and `<html lang>`

On mobile viewport the overlay bundle is not loaded, so it cannot interfere with the mobile navigation sheet. If you resize from desktop to mobile after the overlay has already loaded, the widget shows a small reload hint above the launcher.

Use **Hide overlay** to set an opt-out cookie and reload. The next SSR response contains no overlay loader or bundle script.

Build the overlay bundle with `bun run build:dev-overlay` (also runs during `bun run dev` when overlay is enabled). The bundle is built as an IIFE and embedded from [`internal/devoverlay/static/overlay.js`](internal/devoverlay/static/overlay.js).

Overlay strings live in [`internal/devoverlay/fixtures/locale/`](internal/devoverlay/fixtures/locale/). Portability notes: [`internal/devoverlay/README.md`](internal/devoverlay/README.md).

## Verification

```bash
bun run verify
```

Runs: `templ generate` → Tailwind build → `build:js` → `ui8px lint` → `validate:aria` → `go test ./...`.

## Adding a page

1. Add copy to [`internal/fixtures/fixtures.go`](internal/fixtures/fixtures.go) and every [`internal/fixtures/locale/*.json`](internal/fixtures/locale/) file.
2. Add `internal/views/<page>.templ` — `templ <Page>(d layout.Data, f fixtures.Locale)` composing `@layout.RootLayout` + `@layout.TopnavLayout` or `@layout.DashboardLayout`.
3. Add one **`PageSpec`** in [`internal/site/router.go`](internal/site/router.go) with `Body: views.<Page>`.
4. Run `bun run verify` before landing the change.

See [`docs/for-react-devs.md`](docs/for-react-devs.md) for the full cookbook (Next/shadcn mental model, request flow, dev loop).

## License

MIT — see [LICENSE](LICENSE). Copyright (c) 2026 FastyGo.

Depends on [a-h/templ](https://github.com/a-h/templ), [github.com/fastygo/framework](https://github.com/fastygo/framework), and [github.com/fastygo/templ](https://github.com/fastygo/templ); those projects remain under their respective licenses.
