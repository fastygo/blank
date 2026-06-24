# `registry:layout`

Application-owned **document/chrome shells** — **stays in the app** after blocks/widgets are extracted.

This folder hosts the **named document shells** that route pages compose directly:

| Shell | Composition | Used by |
|-------|-------------|---------|
| `layout.Shell` | `html + head + body + Header + main{children} + Footer` (mobile sheet host) | `views.HomePage` (topnav) |
| `layout.SidebarShell` | `Shell + flex row (aside slot + main{children})` | `views.SamplePage` (sidebar) |

Both shells render the **full HTML document** — pages do `@layout.Shell(data.Shell) { ... }` or `@layout.SidebarShell(data.Shell, appsidebar.AppSidebar(data.Sidebar)) { ... }` as the outermost expression in their `.templ`.

## Blank packages

| Area | Files | Role |
|------|-------|------|
| Topnav shell | `shell.templ`, `shell_templ.go` | Document frame + main slot; hosts `components/navigation` mobile sheet |
| Sidebar shell | `sidebar_shell.templ`, `sidebar_shell_templ.go` | Wraps `Shell` with desktop aside slot beside main column |
| Header | `header.templ` | Top bar; desktop nav + mobile trigger via `components/navigation` |
| Footer | `footer.templ` | App footer |
| Glue | `props.go`, `helpers.go` | `ShellProps`, theme/brand helpers; nav types alias `components/navigation` |

## What belongs here permanently

- `DOCTYPE`, `html`, `head`, `body` document shell
- Shared header/footer chrome used by all routes
- Mobile sheet **host** placement in `Shell`; markup lives in `components/navigation`
- Props/helpers passed from `internal/site/layout_data.go`
- **Named shells** (`Shell`, `SidebarShell`, …) composed directly from pages

## What does **not** belong here

- Page-specific aside markup → `internal/ui/components/appsidebar/`
- Sidebar geometry variants forked from a product need → `blocks/<domain>/<organism>/` as a **full scaffold**, not as an adapter wrapper around `Shell`
- Docs toc + content column → `blocks/docs/toc_shell/` (full scaffold)
- Marketing landing shell variants → `blocks/marketing/*` (full scaffold)
- Fetch/stateful shells → `widgets/`

Do **not** add `internal/ui/blocks/layout/` — use domain folders under `blocks/` instead.

## Sidebar slot contract

`layout.SidebarShell(shell, sidebar)` accepts the aside as `templ.Component`. The supported populator is `components/appsidebar.AppSidebar(props)`. Pass `nil` to omit the aside (rare; matches Next "no sidebar variant"). Mobile navigation is served by the inherited mobile sheet — no separate mobile slot.

## Rules

- Compose with `github.com/fastygo/templ/ui` and `templ/components` where appropriate.
- Preserve `data-ui8kit-*` hooks for `theme.js` / `ui8kit.js` / `@ui8kit/aria` (see `blank-aria.mdc`).
- Routing and locale resolution live in `internal/site/` — layout receives resolved props only.
- **No** `views.*Shell` route adapter functions. Pages compose shells directly.

For sidebar history, see git branch **`sidebar`**.
