# `registry:blocks`

Section-level and **layout organism** wireframes (hero, dashboard shell, sidebar app shell, docs toc shell, …).
**shadcn-style blocks**: self-contained templ + **default English copy** in-package (`defaults.go` / `placeholders.go`).

Blocks are **copy-paste registry artifacts**, not runtime routers. `internal/site/router.go` chooses which adapter renders a block; the block itself does not register routes.

## Domain folders (frozen convention)

Use **`blocks/<domain>/<organism>/`** — not `blocks/layout/` (conflicts with app chrome in `internal/ui/layout/`).

| Package | Example organisms | Showcase focus |
|---------|-------------------|----------------|
| `dashboard/` | `app_shell`, `sidebar_app` | App/dashboard shell wireframes |
| `marketing/` | `topnav_shell`, `landing_shell` | Public/landing layouts |
| `docs/` | `toc_shell` | Docs toc + content column |

Add new **domain** top-level folders only when a new showcase group is needed (e.g. `storefront/`, `editorial/`). Add **organism** subfolders under an existing domain for new layout or section scaffolds.

## Layout organisms vs app chrome

| Concern | Location |
|---------|----------|
| Document frame, shared header/footer, mobile sheet host | `internal/ui/layout/` (stays in app) |
| Sidebar grid, dashboard shell, docs toc layout | `blocks/<domain>/<organism>/` |
| Shell with API fetch / live nav | `widgets/` |

Three sidebar wireframe images (`sidebars_full`, `sidebars_main`, `sidebars_header`) are **showcase ids** for block packages — not runtime names or a layout engine.

## Rules

- Compose with `github.com/fastygo/templ/ui` and `templ/components` — **no raw HTML tags**.
- Tailwind + semantic tokens only; pass **ui8px** policy.
- **Do not** `require github.com/fastygo/blocks` during active staging in this repo.
- Interactive patterns: `data-ui8kit` + static ARIA + manifest (see `blank-aria.mdc`); prefer `templ/components` Sheet with `Behavior: "ui8kit"`.
- **No custom JS** for covered W3C patterns.

## Extraction

When a block is stable, move the package to **`github.com/fastygo/blocks/<name>`** and `require` it from the app.
Keep default data inside the block package; `internal/fixtures` only for i18n overlay.

## Related

- Registry index: [`../README.md`](../README.md)
- Architecture: [`.project/specs/next-shadcn-architecture.md`](../../.project/specs/next-shadcn-architecture.md)
