# In-app UI registry (`internal/ui`)

**Reference layout** for staging reusable UI before promotion to shared modules.
Use this tree in **Blank**, **BuildY**, and product apps the same way — copy the
structure, fill packages incrementally, extract when stable.

## Positioning

| Layer | Module | Role |
|-------|--------|------|
| Atoms / helpers | `github.com/fastygo/templ/utils` | Cn, CVA, tags, ARIA |
| Molecules | `github.com/fastygo/templ/ui` | Button, Stack, Form, Table, … |
| Kit composites | `github.com/fastygo/templ/components` | Card, Alert, Breadcrumb, … |
| App registry | **`internal/ui/*`** (this tree) | Shell, blocks, widgets, app components |

**Not Go dependencies during staging:** `github.com/fastygo/blocks`, `github.com/fastygo/widgets`.
Develop here first; `require` shared modules only after extraction.

## Tree

```
internal/ui/
  layout/       # registry:layout — shell, header, footer, mobile sheet (stays in app)
  components/   # registry:components — dumb reusable UI (icon, toggles, …)
  blocks/       # registry:blocks — section organisms (staging → fastygo/blocks)
    dashboard/
    marketing/
    docs/
  widgets/      # registry:widgets — UI + behavior (staging → fastygo/widgets)
  variants/     # registry:variants — optional wireframe utility maps
  utils/        # registry:utils — thin helpers on templ/utils
```

| Label | Path | Role |
|--------|------|------|
| `registry:layout` | `layout/` | App chrome and route frame; **not** extracted to blocks. |
| `registry:blocks` | `blocks/<group>/` | Section scaffolds + in-package default copy. |
| `registry:components` | `components/` | Props in, markup out; no HTTP, domain, or fetch. |
| `registry:widgets` | `widgets/` | Fetch, state, orchestration; composes blocks/components. |
| `registry:variants` | `variants/` | Named, ui8px-safe utility presets. |
| `registry:utils` | `utils/` | App-specific helpers; generic logic stays in **templ/utils**. |

There is **no** `internal/ui/elements` and **no** `internal/ui/ui/` layer.

## `components` vs `widgets` vs `blocks`

- **`blocks`** — section-level markup + default demo data; portable wireframe.
- **`components`** — small reusable pieces; props only.
- **`widgets`** — same UI boundaries plus **behavior** (API, loading, subscriptions).

If it only renders props → **`components`**. If it **fetches** or coordinates side effects → **`widgets`**.

## Data (three layers)

1. **Block/component defaults** — English wireframe copy in-package (`defaults.go` / `placeholders.go`).
2. **`internal/fixtures`** — app structs and embedded locale JSON.
3. **Views** — `internal/views/` compose layout + **resolved** strings; no fixture loads inside `.templ`.

## Promotion & extraction

| Stable & generic | Destination |
|------------------|-------------|
| Primitive / small composite | `github.com/fastygo/templ` (`ui/` or `components/`) |
| Section block | `github.com/fastygo/blocks/<group>` |
| Interactive / API shell | `github.com/fastygo/widgets` |
| App shell | **keep** `internal/ui/layout` |

**Freeze before extract:** stable props, 2+ uses or explicit registry intent, default data in block package, dependencies only on **templ** (+ framework for widgets if needed).

## Composition rules

- **No raw HTML** layout/content tags — use `templ/ui` (+ `templ/components`).
- Document shell only: `<!DOCTYPE>`, `<html>`, `<head>`, `<body>`, … in `layout/shell.templ` or `views/partials/shell_head.templ`.
- Tailwind utilities must pass **`Blank/.ui8px/policy`** (`bun run lint:ui8px`).

## Related docs

- [`layout/README.md`](layout/README.md)
- [`components/README.md`](components/README.md)
- [`blocks/README.md`](blocks/README.md)
- [`widgets/README.md`](widgets/README.md)
- [`variants/README.md`](variants/README.md)
- [`utils/README.md`](utils/README.md)
- App rules: `@Blank/.cursor/rules/blank-ui-structure.mdc`, `blank-atomic-ui8px.mdc`

## Source policy

Registry folder rules match **FastyGoUI** design-system policy (`registry:*` labels).
When FastyGoUI is not in the workspace, **this README + subtree READMEs** are the on-disk index.
