# `registry:layout`

Application-owned **document/chrome infrastructure** — **stays in the app** after blocks/widgets are extracted.

This folder is **not** a catalog for every sidebar, dashboard, or docs layout variant. Those are **layout organisms** under `internal/ui/blocks/<domain>/<organism>/`.

## Blank packages

| Area | Files | Role |
|------|-------|------|
| Shell | `shell.templ` | Document frame, main slot, mobile sheet host |
| Header | `header.templ`, `nav.templ` | Top bar, desktop nav, mobile trigger |
| Footer | `footer.templ` | App footer |
| Glue | `props.go`, `helpers.go` | Shell props, nav helpers |

## What belongs here permanently

- `DOCTYPE`, `html`, `head`, `body` document shell
- Shared header/footer/nav chrome used by all routes
- Mobile sheet **host** markup and `@ui8kit/aria` hooks
- Props/helpers passed from `internal/site/layout_data.go`

## What does **not** belong here

- Sidebar app shell geometry → `blocks/dashboard/sidebar_app` (or similar)
- Docs toc + content column → `blocks/docs/toc_shell`
- Marketing landing shell variants → `blocks/marketing/*`
- Fetch/stateful shells → `widgets/`

Do **not** add `internal/ui/blocks/layout/` — use domain folders under `blocks/` instead.

## Rules

- Compose with `github.com/fastygo/templ/ui` and `templ/components` where appropriate.
- Preserve `data-ui8kit-*` hooks for `theme.js` / `ui8kit.js` / `@ui8kit/aria` (see `blank-aria.mdc`).
- Routing and locale resolution live in `internal/site/` — layout receives resolved props only.
- Route shell adapters (`views.AppShell`, …) may delegate into registry blocks later; they stay in `internal/views/`, not here.

For sidebar-style shell wireframes, see Block 07+ in [`.project/next-shadcn-refactor-progress.md`](../../.project/next-shadcn-refactor-progress.md) and git branch **`sidebar`** as historical reference.
