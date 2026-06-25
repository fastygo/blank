# `registry:layout`

Application-owned **document/chrome shells** — **stays in the app** after blocks/widgets are extracted.

This folder hosts the **named document shells** that route pages compose directly, plus request-scoped layout data and head/trailing partials.

| Shell | Composition | Used by |
|-------|-------------|---------|
| `layout.Shell` | `html + head + body + Header + main{children} + Footer` (mobile sheet host) | `views.HomePage` (topnav) |
| `layout.SidebarShell` | `Shell + flex row (aside slot + main{children})` | `views.SamplePage` (sidebar) |

Both shells render the **full HTML document** — pages pass `d.ShellProps()` from `layout.Data`.

## Blank packages

| Area | Files | Role |
|------|-------|------|
| Topnav shell | `shell.templ` | Document frame + main slot; hosts `components/navigation` mobile sheet |
| Sidebar shell | `sidebar_shell.templ` | Wraps `Shell` with desktop aside slot beside main column |
| Header / footer | `header.templ`, `footer.templ` | Shared chrome |
| Head / trailing | `shell_head.templ`, `header_trailing.templ` | Favicons, CSS/JS tags, locale switch slot |
| Request data | `data.go` | `Data`, `AssetPaths`, `ShellProps()`, `SidebarProps()` |
| Build | `build.go` | `BuildData(BuildParams)` — assets, nav props, language switch |
| Types / helpers | `props.go`, `helpers.go` | `ShellProps`, theme/brand helpers |

## What belongs here permanently

- `DOCTYPE`, `html`, `head`, `body` document shell
- Shared header/footer chrome used by all routes
- Mobile sheet **host** placement in `Shell`; markup lives in `components/navigation`
- Request `layout.Data` and prop builders (`ShellProps`, `SidebarProps`)
- **Named shells** (`Shell`, `SidebarShell`, …) composed directly from pages

## What does **not** belong here

- Page-specific aside markup → `internal/ui/components/appsidebar/`
- Section scaffolds with default copy → `blocks/<domain>/<organism>/` (e.g. `marketing/hero`)
- Fetch/stateful shells → `widgets/`

Do **not** add `internal/ui/blocks/layout/` — use domain folders under `blocks/` instead.

## Sidebar slot contract

`layout.SidebarShell(shell, sidebar)` accepts the aside as `templ.Component`. The supported populator is `components/appsidebar.AppSidebar(props)`. Pass `nil` to omit the aside (rare). Mobile navigation is served by the inherited mobile sheet.

## Rules

- Compose with `github.com/fastygo/templ/ui` and `templ/components` where appropriate.
- Preserve `data-ui8kit-*` hooks for `theme.js` / `ui8kit.js` / `@ui8kit/aria` (see `blank-aria.mdc`).
- `internal/site/layout_data.go` calls `layout.BuildData` — routing stays in `internal/site/`.
- **No** `views.*Shell` route adapter functions. Pages compose shells directly.

For sidebar history, see git branch **`sidebar`**.
