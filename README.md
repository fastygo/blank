# Blank

Official **Go + templ** starter for the [FastyGo](https://github.com/fastygo) stack.
Clone it as a **universal** base — local kit atoms, FastyGoUI-aligned registry, retag policy — with no product brand.

**Module:** `github.com/fastygo/blank`

Demo routes: topnav home (`/`) and sidebar dashboard sample (`/sample`), with mobile sheet, dark theme, and En/Ru locale switching. Each page composes its own layout layers in `internal/views/`; routing lives in `internal/site/router.go`.

## What this repo is

| Layer | Location | Role |
|-------|----------|------|
| Framework wiring | [`internal/serverapp/`](internal/serverapp/), [`cmd/server/`](cmd/server/) | Locales, security, site feature, dev overlay |
| Kit (atoms) | [`internal/kit/`](internal/kit/) | Codegen primitives + utils (stand-in for `fastygo/templ`) |
| UI registry | [`internal/ui/`](internal/ui/) | Layout shells, components, blocks (FastyGoUI tree) |
| Pages | [`internal/views/`](internal/views/) | Route bodies — compose layout + blocks only |
| Copy & i18n | [`internal/fixtures/locale/`](internal/fixtures/locale/) | Embedded En/Ru strings |
| Project docs | [`.project/`](.project/) | Architecture, retag, ADRs |
| Dev tooling | [`fastygo.config.mjs`](fastygo.config.mjs), [`scripts/`](scripts/) | CSS/retag/JS — not routes |

## Prerequisites

- Go 1.25+
- [Bun](https://bun.sh) (CSS build + config helpers)
- `templ` CLI (`go tool templ` / `templ generate`)

## Quick start

```bash
bun install
go mod download
bun run dev
```

Open [http://0.0.0.0:3000/](http://127.0.0.1:3000/) — hero welcome (`RootLayout` + `TopnavLayout`). Sample: [http://127.0.0.1:3000/sample](http://127.0.0.1:3000/sample) (`RootLayout` + `DashboardLayout`).

`bun run dev` builds CSS/JS/overlay, then [`scripts/dev.sh`](scripts/dev.sh) (`templ generate` + `go run ./cmd/server`). After `.templ` edits run `bun run templ` (or restart). After Tailwind class changes run `bun run build:css`. **Ctrl+C** stops the server.

Tooling SoT: [`fastygo.config.mjs`](fastygo.config.mjs) (`registry`, `retag`, `css`). Go does not load it.

## Scripts

| Command | Purpose |
|---------|---------|
| `bun run dev` / `g` | CSS + JS + overlay + Go server |
| `bun run build:css` | Tailwind via config paths |
| `bun run retag` | Zoned raw-HTML ratchet |
| `bun run test` | Retag + `go test ./...` |
| `bun run verify` | templ + CSS + JS + retag + tests |
| `bun run config` | Dump tooling config JSON |

## For React developers

See [`docs/for-react-devs.md`](docs/for-react-devs.md). Maintainer docs: [`.project/`](.project/) and [`docs/architecture.md`](docs/architecture.md).

## Environment

| Variable | Default | Purpose |
|----------|---------|---------|
| `APP_BIND` | `0.0.0.0:3000` | HTTP listen address |
| `APP_DEV_OVERLAY` | `1` in local dev via config | Dev status overlay (loopback only) |
| `APP_STATIC_DIR` | `web/static` | Static files under `/static/` |
| `APP_DEFAULT_LOCALE` | `en` | Default locale |
| `APP_AVAILABLE_LOCALES` | `en,ru` | Header locale switcher |

Probes: `GET /healthz` and `GET /readyz`.

## Project layout

| Path | Role |
|------|------|
| [`fastygo.config.mjs`](fastygo.config.mjs) | Tooling SoT: server, css, registry, retag |
| [`internal/kit/`](internal/kit/) | Atoms (`ui/*`) + `utils` |
| [`internal/ui/`](internal/ui/) | Registry — layout / components / blocks / widgets |
| [`internal/views/`](internal/views/) | Thin page composition |
| [`internal/site/`](internal/site/) | Runtime route manifest |
| [`internal/fixtures/locale/`](internal/fixtures/locale/) | UI copy per locale |
| [`web/static/css/`](web/static/css/) | `tokens.css` + `theme-blank.css` |
| [`.project/`](.project/) | Architecture, retag, ADRs |

## Retag

Composition layers must not own raw HTML. See [`.project/retag.md`](.project/retag.md) and run `bun run retag`.

## Dev overlay

When `APP_DEV_OVERLAY=1` on loopback, Blank injects a viewport-gated status overlay. Details: [`internal/devoverlay/README.md`](internal/devoverlay/README.md).

## Adding a page

1. Add copy to fixtures + every [`internal/fixtures/locale/*.json`](internal/fixtures/locale/) file.
2. Add `internal/views/<page>.templ` composing `@layout.RootLayout` + a shell layout + blocks.
3. Add one **`PageSpec`** in [`internal/site/router.go`](internal/site/router.go).
4. Run `bun run verify` before landing.

## License

MIT — see [LICENSE](LICENSE). Copyright (c) 2026 FastyGo.

Depends on [a-h/templ](https://github.com/a-h/templ), [github.com/fastygo/framework](https://github.com/fastygo/framework); kit atoms live in-repo. Optional Sheet still uses [github.com/fastygo/templ](https://github.com/fastygo/templ) `components` until Sheet lands in kit.
