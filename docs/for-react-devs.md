# Blank for React developers

Blank is a **server-rendered BFF frontend**: HTML is composed on the Go server with [templ](https://templ.guide/), styled with Tailwind, and served with static assets under `/static/`. It is not a client-side SPA, but the dev workflow is intentionally close to **Vite + SSR**.

## Mental model

| React + Vite | Blank |
|--------------|-------|
| `vite.config.ts` | [`fastygo.config.mjs`](../fastygo.config.mjs) |
| `npm run dev` | `bun run dev` |
| `npm run build` | `bun run build` |
| `npm run preview` | `bun run preview` (one-shot server, no watch) |
| `public/` | [`web/static/`](../web/static/) |
| `src/components/` | [`internal/ui/components/`](../internal/ui/components/) |
| `src/pages/` or `app/routes` | [`internal/views/`](../internal/views/) + routes in [`internal/site/feature.go`](../internal/site/feature.go) |
| JSX / TSX | `.templ` templates |
| React props | Go structs passed into `@Component(props)` |
| i18n JSON / CMS strings (site) | [`internal/fixtures/locale/`](../internal/fixtures/locale/) |
| Dev overlay i18n | [`internal/devoverlay/fixtures/locale/`](../internal/devoverlay/fixtures/locale/) |
| Client interactivity (dialog, sheet) | [`@ui8kit/aria`](../web/static/js/ui8kit.js) + `data-ui8kit` hooks |
| React DevTools-like status panel | Dev overlay (`APP_DEV_OVERLAY=1`) with Health / Assets / Request tabs |

## Routes and layouts

Blank uses **two layout layers**. When docs say “layout”, check which one is meant.

| Next App Router | Blank | Role |
|-----------------|-------|------|
| Root document frame | [`internal/ui/layout/shell.templ`](../internal/ui/layout/shell.templ) → `layout.Shell` | **Document shell** — `html`, `head`, `body`, header, footer, mobile sheet, assets |
| `app/(app)/layout.tsx` | [`internal/views/layout.templ`](../internal/views/layout.templ) → `views.AppShell` | **Route shell** — app zone (topnav today) |
| `app/(marketing)/layout.tsx` | `views.MarketingShell` | Route shell — landing/marketing (currently same chrome as `AppShell`) |
| `app/(docs)/layout.tsx` | `views.DocsShell` | Route shell — docs (placeholder until docs routes) |
| `app/**/page.tsx` | `internal/views/<page>.templ` | Page content only |

**Request flow:** `GET /sample` → handler in [`internal/site/feature.go`](../internal/site/feature.go) → route shell (`AppShell`, …) → page (`SamplePage`).

**Naming:** Prefer `views.AppShell` in new code. `views.SiteShell` is a **temporary alias** to `AppShell` for backward compatibility.

Architecture details: [`.project/specs/next-shadcn-architecture.md`](../.project/specs/next-shadcn-architecture.md).

## Dev loop

1. **Start once:** `bun install && bun run dev`
2. **Edit `.templ` or Tailwind classes** — CSS and templ rebuild automatically; refresh the browser (F5).
3. **Edit Go handlers** in `internal/site/feature.go` — restart dev (`Ctrl+C`, then `bun run dev`). Go auto-restart is not enabled in v1.
4. **Stop:** `Ctrl+C` in the terminal running `dev` (closing a browser tab does not stop the server).

Server URL defaults to [http://127.0.0.1:8080/](http://127.0.0.1:8080/). Override with `APP_BIND` in [`fastygo.config.mjs`](../fastygo.config.mjs) or the environment.

## What `bun run dev` does

```text
fastygo.config.mjs
       ↓
scripts/dev.mjs
  ├─ templ generate (initial)
  ├─ build CSS + JS (initial)
  ├─ tailwind --watch
  ├─ watch internal/**/*.templ → templ generate
  └─ go run ./cmd/server
```

Logs are prefixed with `[fastygo]` (similar to Vite’s `[vite]`).

## Dev overlay

With `APP_DEV_OVERLAY=1` (enabled in [`fastygo.config.mjs`](../fastygo.config.mjs) for local dev), Blank injects a floating widget at SSR time. It does not modify `internal/views/**`.

- **Health:** browser probes for `/healthz` and `/readyz`
- **Assets:** server reports static file age; stale CSS hints `bun run watch:css`
- **Request:** shows `X-Request-ID`, current path, and document locale

Click **Hide overlay** to opt out via cookie and reload. After reload, View Source and Network should show no overlay assets.

Overlay copy is maintained in [`internal/devoverlay/fixtures/locale/`](../internal/devoverlay/fixtures/locale/) — separate from site fixtures. Locale follows the same `?lang=` / cookie rules as the header switcher.

## Adding a page (checklist)

1. Extend [`fixtures.Locale`](../internal/fixtures/fixtures.go) and every file in [`internal/fixtures/locale/`](../internal/fixtures/locale/).
2. Add a nav item in [`internal/site/feature.go`](../internal/site/feature.go) (`siteNav`).
3. Create `internal/views/<page>.templ` and register a handler + route in `internal/site/feature.go`. Use `views.AppShell` (or `MarketingShell` / `DocsShell`) when rendering.

See the main [README](../README.md) for details.

## Verification

```bash
bun run verify
```

Pipeline: `templ generate` → Tailwind CSS → JS bundle → `ui8px lint` → `validate:aria` → `go test ./...`.

Run this before opening a PR or after large markup changes.

## Common commands

| Task | Command |
|------|---------|
| Dev with watch | `bun run dev` |
| One-shot run | `bun run start` |
| Production binary | `bun run build` → `./blank` |
| Regenerate templ only | `bun run templ` |
| CSS only | `bun run build:css` |
| Full check | `bun run verify` |

## SSR vs HMR

Vite HMR updates modules in the browser without a full reload. Blank uses **SSR**: the server renders full HTML documents. Template and CSS watchers rebuild artifacts on disk; you refresh the page to see changes. That is normal for this stack and matches many Vite SSR setups without a custom HMR plugin.
