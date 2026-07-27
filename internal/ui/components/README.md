# `registry:components`

**Self-contained, reusable** templ composites — props in, markup out. No HTTP, domain fetch, or app fixtures inside the package.

## Contract

| Rule | Detail |
|------|--------|
| **Dependencies** | `github.com/fastygo/blank/internal/kit/ui`, `templ/components`, `templ/utils` — plus other `internal/ui/components/*` when composed |
| **Defaults** | Wireframe copy in `defaults.go` or `placeholders.go` when the component ships demo content |
| **Overrides** | Views pass structs with resolved strings; `internal/fixtures/locale` overlays i18n at the view layer |
| **Files** | Prefer single self-contained `*.templ` per package (types, helpers, markup) |

## Current packages (Blank)

| Package | Role |
|---------|------|
| `icon/` | Latty mask icons |
| `navigation/` | Reusable nav links + mobile sheet (`Nav`, `MobileSheet`, `MobileSheetTrigger`) via `templ/components` Sheet with `Behavior: "ui8kit"` |
| `toggles/` | Language switch (`language_switch.templ`) |

## Heuristic

Pure presentation → **`components`**. Section scaffolds → **`blocks/<group>/`**. Fetch/state → **`widgets/`**.
