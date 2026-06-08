# `registry:blocks`

Section-level wireframe organisms (hero, feature grid, dashboard sections, …).
**shadcn-style blocks**: self-contained templ + **default English copy** in-package (`defaults.go` / `placeholders.go`).

## Groups (reserved)

| Package | Showcase focus |
|---------|----------------|
| `dashboard/` | Dashboard / home sections |
| `marketing/` | Landing, hero, CTA scaffolds |
| `docs/` | Documentation-style sections |

Add new top-level folders only when a **new showcase group** is needed (e.g. `storefront/`, `editorial/`).

## Rules

- Compose with `github.com/fastygo/templ/ui` and `templ/components` — **no raw HTML tags**.
- Tailwind + semantic tokens only; pass **ui8px** policy.
- **Do not** `require github.com/fastygo/blocks` during active staging in this repo.
- Interactive patterns: `data-ui8kit` + static ARIA + manifest (see `blank-aria.mdc`).

## Extraction

When a block is stable, move the package to **`github.com/fastygo/blocks/<name>`** and `require` it from the app.
Keep default data inside the block package; `internal/fixtures` only for i18n overlay.
