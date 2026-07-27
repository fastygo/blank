# `registry:blocks`

**Full scaffolds** — section-level wireframes and self-contained product pages
that ship with their own default copy. **shadcn-style blocks:** copy-paste
templ + in-package `defaults.go` / `placeholders.go`.

Blocks are **copy-paste registry artifacts**, not runtime routers and not
adapter shells. `internal/site/router.go` chooses which page renders; pages
compose their own layout via `internal/ui/layout/` shells; blocks are only used
when a complete pre-built section is dropped in.

## What belongs here

A block is a **stable, complete artifact** — equivalent to a shadcn block under
`apps/v4/registry/new-york-v4/blocks/`. Examples:

- A full hero section with its own copy + variants.
- A complete dashboard view (charts, table, filters) with default demo data.
- A docs toc shell that owns its own grid + sample structure.

## What does **not** belong here

- **Empty adapters.** Packages whose only job is `@layout.TopnavLayout { @body }` belong nowhere — delete them. Use `internal/ui/layout/` layers directly from the page.
- **Aside / sidebar markup** consumed by a layout layer — put it in `internal/ui/components/appsidebar/` (or another component package).
- **Anything renaming a layout layer.** Pages compose `layout.RootLayout` + `layout.TopnavLayout` / `layout.DashboardLayout` directly.

The previous `blocks/dashboard/app_shell/` and `blocks/dashboard/sidebar_app/`
adapter packages were removed — they were indirection, not scaffolds. Their
roles moved to:

- Document frame → `internal/ui/layout/root_layout.templ`
- Topnav chrome → `internal/ui/layout/topnav_layout.templ`
- Dashboard chrome → `internal/ui/layout/dashboard_layout.templ`
- Aside markup → `internal/ui/components/appsidebar/`

## Domain folders (frozen convention)

Use **`blocks/<domain>/<organism>/`** — not `blocks/layout/` (conflicts with app
chrome in `internal/ui/layout/`).

| Package | Example organisms | Showcase focus |
|---------|-------------------|----------------|
| `marketing/` | [`hero`](marketing/hero/) | Landing hero section (welcome + title + description) |

Add new **domain** top-level folders only when a new showcase group is needed
(e.g. `storefront/`, `editorial/`). Add **organism** subfolders under an
existing domain for new full scaffolds.

## Rules

- Compose with `github.com/fastygo/blank/internal/kit/ui` and `templ/components` — **no raw HTML tags**.
- Tailwind + semantic tokens only; pass **ui8px** policy.
- **Do not** `require github.com/fastygo/blocks` during active staging in this repo.
- Interactive patterns: `data-ui8kit` + static ARIA + manifest (see `blank-aria.mdc`); prefer `templ/components` Sheet with `Behavior: "ui8kit"`.
- **No custom JS** for covered W3C patterns.

## Extraction

When a block is stable, move the package to **`github.com/fastygo/blocks/<name>`** and `require` it from the app.
Keep default data inside the block package; `internal/fixtures` only for i18n overlay.

## Related

- Registry index: [`../README.md`](../README.md)
- Layout shells: [`../layout/README.md`](../layout/README.md)
- Local aside component: [`../components/appsidebar/README.md`](../components/appsidebar/README.md)
- Architecture: [`docs/architecture.md`](../../docs/architecture.md)
