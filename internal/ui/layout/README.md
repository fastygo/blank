# `registry:layout`

Application-owned **document/chrome shells** — **stays in the app** after blocks/widgets are extracted.

This folder hosts the **two-layer layout model** that route pages compose directly, plus request-scoped layout data and head/trailing partials.

| Layer | Component | Composition | Used by |
|-------|-----------|-------------|---------|
| Document | `layout.RootLayout` | `html + head + body{children}` only | Every page |
| Topnav chrome | `layout.TopnavLayout` | `Header + main{children} + Footer + MobileSheet` | `views.HomePage` |
| Dashboard chrome | `layout.DashboardLayout` | `TopnavLayout + aside + main column{children}` | `views.SamplePage` |

Pages compose **two layers**: `@layout.RootLayout(d.Document()) { @layout.TopnavLayout(d.Topnav()) { ... } }` or `@layout.RootLayout { @layout.DashboardLayout(d.Dashboard(title)) { ... } }`.

## Blank packages

| Area | Files | Role |
|------|-------|------|
| Document frame | `root_layout.templ` | `html`, `head`, `body` — no app chrome |
| Topnav chrome | `topnav_layout.templ` | Header, main slot, footer, mobile sheet |
| Dashboard chrome | `dashboard_layout.templ` | TopnavLayout + desktop aside slot |
| Header / footer | `header.templ`, `footer.templ` | Shared chrome used by TopnavLayout |
| Head / trailing | `shell_head.templ`, `header_trailing.templ` | Favicons, CSS/JS tags, locale switch slot |
| Request data | `data.go` | `Data`, `Document()`, `Topnav()`, `Dashboard(title)` |
| Build | `build.go`, `build_document.go`, `build_topnav.go`, `build_dashboard.go` | `BuildData(BuildParams)` and per-layer helpers |
| Types / helpers | `props.go`, `helpers.go` | `DocumentProps`, `TopnavLayoutProps`, `DashboardLayoutProps` |

## What belongs here permanently

- `DOCTYPE`, `html`, `head`, `body` document frame (`RootLayout`)
- Shared header/footer chrome used by all routes (`TopnavLayout`)
- Mobile sheet placement in `TopnavLayout`; markup lives in `components/navigation`
- Request `layout.Data` and prop builders (`Document()`, `Topnav()`, `Dashboard()`)
- **Named layout layers** (`RootLayout`, `TopnavLayout`, `DashboardLayout`) composed directly from pages

## What does **not** belong here

- Page-specific aside markup → `internal/ui/components/appsidebar/`
- Section scaffolds with default copy → `blocks/<domain>/<organism>/` (e.g. `marketing/hero`)
- Fetch/stateful shells → `widgets/`

Do **not** add `internal/ui/blocks/layout/` — use domain folders under `blocks/` instead.

## Dashboard aside contract

`layout.DashboardLayout(props)` renders `appsidebar.AppSidebar(props.Sidebar)` as the desktop aside. Mobile navigation is inherited from `TopnavLayout` via the mobile sheet.

## Rules

- Compose with `github.com/fastygo/templ/ui` and `templ/components` where appropriate.
- Preserve `data-ui8kit-*` hooks for `theme.js` / `ui8kit.js` / `@ui8kit/aria` (see `blank-aria.mdc`).
- `internal/site/layout_data.go` calls `layout.BuildData` — routing stays in `internal/site/`.
- **No** `views.*From` wrapper functions. Pages are `views.<Page>(d, f)` registered directly in the router.

For sidebar history, see git branch **`sidebar`**.
