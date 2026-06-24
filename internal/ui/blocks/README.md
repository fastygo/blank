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

- **Empty adapters.** Packages whose only job is `@layout.Shell { @body }` belong nowhere — delete them. Use `internal/ui/layout/` shells directly from the page.
- **Aside / sidebar markup** consumed by a layout shell — put it in `internal/ui/components/appsidebar/` (or another component package).
- **Anything renaming a layout shell.** Pages compose `layout.Shell` / `layout.SidebarShell` directly.

The previous `blocks/dashboard/app_shell/` (4 lines wrapping `Shell`) and
`blocks/dashboard/sidebar_app/` (aside + main wrapper around `Shell`) were
removed in the **page-composes-layout** refactor — they were adapter
indirection, not scaffolds. Their content moved to:

- `Shell` chrome → `internal/ui/layout/shell.templ` (unchanged)
- Sidebar geometry → `internal/ui/layout/sidebar_shell.templ`
- Aside markup → `internal/ui/components/appsidebar/`

## Domain folders (frozen convention)

Use **`blocks/<domain>/<organism>/`** — not `blocks/layout/` (conflicts with app
chrome in `internal/ui/layout/`).

| Package | Example organisms | Showcase focus |
|---------|-------------------|----------------|
| `dashboard/` | (currently empty — stub) | Future dashboard scaffolds |
| `marketing/` | `topnav_shell`, `landing_shell` | Public/landing layouts |
| `docs/` | `toc_shell` | Docs toc + content column |

Add new **domain** top-level folders only when a new showcase group is needed
(e.g. `storefront/`, `editorial/`). Add **organism** subfolders under an
existing domain for new full scaffolds.

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
- Layout shells: [`../layout/README.md`](../layout/README.md)
- Local aside component: [`../components/appsidebar/README.md`](../components/appsidebar/README.md)
- Architecture: [`.project/specs/next-shadcn-architecture.md`](../../.project/specs/next-shadcn-architecture.md)
