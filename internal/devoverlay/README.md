# Dev overlay

Dev-only overlay for local FastyGo apps: **Health**, **Assets**, and **Request** tabs. The HTTP layer injects only a viewport-gated loader script; desktop clients load an isolated Web Component with Shadow DOM, while mobile clients do not fetch the overlay bundle — no edits to `internal/views` or site routes.

## What ships in this folder

| Path | Role |
|------|------|
| `config.go`, `locale.go` | Enable gate, loopback check, locale negotiation |
| `inject.go`, `middleware.go` | HTML inject before `</body>` |
| `routes.go`, `assets.go` | `/__fastygo/dev/*` routes + `status.json` |
| `fixtures/` | **Own** embedded `locale/en.json`, `locale/ru.json` |
| `ui/` | Templ loader script (`widget.templ`) |
| `frontend/` | TypeScript web component, shadow styles, panel logic |
| `static/overlay.js` | Bundled client (`go:embed`) |

## Host wiring (Blank)

[`internal/serverapp/app.go`](../serverapp/app.go):

```go
overlayCfg := devoverlay.Load(cfg.Config)
handler := application.Handler()
if overlayCfg.Enabled {
    handler = devoverlay.Wrap(handler, overlayCfg)
}
```

Requires `APP_DEV_OVERLAY=1` and loopback bind (`127.0.0.1`, `localhost`, …).

## Build

```bash
go tool templ generate ./internal/devoverlay/...
bun run build:dev-overlay
```

Overlay CSS is bundled into the web component shadow root. Host Tailwind does not scan overlay sources.

## Locale

Overlay strings live in **`fixtures/locale/`** — separate from app [`internal/fixtures`](../fixtures/).

Locale code is resolved per request in [`locale.go`](locale.go) (`?lang=`, `lang` cookie, `Accept-Language`) because `devoverlay.Wrap` sits outside app locale middleware. Rules mirror the site; JSON files do not.

Client panels read serialized copy from `data-i18n` on the loader script; the loader copies it to the desktop-only overlay bundle script.

## Config knobs

| Field | Default | Notes |
|-------|---------|-------|
| `DefaultLocale` | from `app.Config` | Fallback when locale unknown |
| `AvailableLocales` | from `app.Config` | Negotiation allow-list |
| `LangCookieName` | `lang` | Must match site cookie |
| `HealthPaths` | `/healthz`, `/readyz` | Health tab probes (TS defaults match) |
| `StaleCSSSeconds` | `300` | Stale hint for `app.css` |
| `Assets` | app.css, ui8kit.js, theme.js | Customize per host app |

Use `devoverlay.Load(app.Config)` in Blank or construct `devoverlay.Config{...}` directly in other apps.

## JS DOM contract

Do not rename without updating `frontend/main.ts`:

- `#fastygo-dev-overlay-loader` with `data-request-id`, `data-i18n`
- `/__fastygo/dev/overlay.js` is dynamically loaded only when `(max-width: 767px)` does not match
- `<fastygo-dev-overlay>` custom element is created by `overlay.js` after the desktop gate
- `#fastygo-dev-launcher`, `#fastygo-dev-panel`, `[data-dev-tab]`, and panel mount nodes live inside the custom element shadow root

## Portability checklist

Copy to another FastyGo app:

1. Copy entire `internal/devoverlay/` tree.
2. Copy `scripts/dev-overlay-entry.ts`, `scripts/build-dev-overlay.mjs`; add `devOverlay` to `fastygo.config.mjs`.
3. Add `"build:dev-overlay"` to `package.json`; run from `dev.mjs` when `APP_DEV_OVERLAY=1`.
4. Wire `devoverlay.Wrap(handler, devoverlay.Load(appConfig))` in the composition root.
5. Set loopback bind + `APP_DEV_OVERLAY=1`.
6. Align `DefaultLocale`, `AvailableLocales`, `LangCookieName` with site locale config.
7. Customize `Config.Assets`, `HealthPaths`, `StaleCSSSeconds` for the target project.
8. Run `go tool templ generate ./internal/devoverlay/...` and `bun run build:dev-overlay`; commit `static/overlay.js`.

**Do not copy** app `internal/fixtures/locale/*` for overlay strings — this package brings its own JSON.

## Dependencies

- `github.com/fastygo/templ` — loader script markup
- Shadow DOM CSS bundled in `overlay.js`; no host Tailwind dependency
- `github.com/fastygo/framework` — locale negotiation in `locale.go` (optional to replace)

## Opt out

POST `/__fastygo/dev/disable` or cookie `fastygo_dev=off`. Next HTML response has no overlay markup or script.
