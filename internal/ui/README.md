# In-app UI registry (`internal/ui`)

**Reference layout** for staging reusable UI before promotion to shared modules.
Use this tree in **Blank**, **BuildY**, and product apps the same way — copy the
structure, fill packages incrementally, extract when stable.

## Positioning

| Layer | Module | Role |
|-------|--------|------|
| Atoms / helpers | `github.com/fastygo/templ/utils` | Cn, CVA, tags, ARIA |
| Molecules | `github.com/fastygo/templ/ui` | Button, Stack, Form, Table, … |
| Kit composites | `github.com/fastygo/templ/components` | Card, Alert, Breadcrumb, Sheet, … |
| App registry | **`internal/ui/*`** (this tree) | Chrome, blocks, widgets, app components |

**Not Go dependencies during staging:** `github.com/fastygo/blocks`, `github.com/fastygo/widgets`.
Develop here first; `require` shared modules only after extraction.

**Not part of this registry:** `internal/views/` (pages + thin route adapters), `internal/site/` (runtime route manifest).

## Tree

```
internal/ui/
  layout/       # registry:layout — document/chrome shell (stays in app)
  components/   # registry:components — icon, toggles, …
  blocks/       # registry:blocks — section/layout organisms (staging → fastygo/blocks)
    dashboard/  # e.g. app_shell (topnav today), sidebar_app
    marketing/  # e.g. topnav_shell, landing_shell
    docs/       # e.g. toc_shell
  widgets/      # registry:widgets — UI + behavior (staging → fastygo/widgets)
  variants/     # registry:variants — optional wireframe utility maps
  utils/        # registry:utils — thin helpers on templ/utils
```

| Label | Path | Role |
|--------|------|------|
| `registry:layout` | `layout/` | App document/chrome frame; **not** extracted; **not** every layout variant |
| `registry:blocks` | `blocks/<domain>/` | Section/layout organisms + in-package default copy |
| `registry:components` | `components/` | Props in, markup out; no HTTP, domain, or fetch |
| `registry:widgets` | `widgets/` | Fetch, state, orchestration; composes blocks/components |
| `registry:variants` | `variants/` | Named, ui8px-safe utility presets |
| `registry:utils` | `utils/` | App-specific helpers; generic logic stays in **templ/utils** |

There is **no** `internal/ui/elements`, **no** `internal/ui/ui/`, and **no** `internal/ui/blocks/layout/` (conflicts with `layout/` app chrome).

## Where does this UI go?

| You are building… | Put it in… |
|-------------------|------------|
| Document shell, header, footer, mobile sheet host | `layout/` |
| Small reusable control (icon, toggle) | `components/` |
| Sidebar app shell, dashboard grid, docs toc layout (props-only wireframe) | `blocks/<domain>/<organism>/` |
| Shell that fetches or orchestrates API/state | `widgets/` |
| Route-specific page content | `internal/views/*.templ` |
| HTTP route + layout choice | `internal/site/router.go` (`PageSpec`) |
| Thin Next-like route wrapper | `internal/views/*Shell` (adapter only) |

Layout organisms use **domain folders** first: `dashboard/sidebar_app`, `docs/toc_shell`, `marketing/topnav_shell` — not a generic `blocks/layout/` tree.

## `components` vs `widgets` vs `blocks`

- **`blocks`** — section-level or layout-level markup + default demo data; portable wireframe (shadcn block copy-paste).
- **`components`** — small reusable pieces; props only.
- **`widgets`** — same UI boundaries plus **behavior** (API, loading, subscriptions).

If it only renders props → **`components`** or **`blocks`**. If it **fetches** or coordinates side effects → **`widgets`**.

## Runtime vs registry (do not collapse)

```text
internal/site/router.go  →  views.AppShell (adapter)  →  internal/ui/blocks/... + layout.Shell  →  views.Page
```

- **`internal/site/`** — runtime wiring only; no reusable markup.
- **`internal/views/*Shell`** — temporary route adapters; should delegate to registry artifacts as they appear.
- **`internal/ui/*`** — shadcn-like accumulation layer; copy-paste or `require` after extraction.

## Data (three layers)

1. **Block/component defaults** — English wireframe copy in-package (`defaults.go` / `placeholders.go`).
2. **`internal/fixtures`** — app structs and embedded locale JSON.
3. **Views** — `internal/views/` compose layout + **resolved** strings; no fixture loads inside `.templ`.

## Promotion & extraction

| Stable & generic | Destination |
|------------------|-------------|
| Primitive / small composite | `github.com/fastygo/templ` (`ui/` or `components/`) |
| Section / layout block | `github.com/fastygo/blocks/<group>` |
| Interactive / API shell | `github.com/fastygo/widgets` |
| App document/chrome shell | **keep** `internal/ui/layout` |

**Freeze before extract:** stable props, 2+ uses or explicit registry intent, default data in block package, dependencies only on **templ** (+ framework for widgets if needed).

## Composition rules

- **No raw HTML** layout/content tags — use `templ/ui` (+ `templ/components`).
- Document shell only: `<!DOCTYPE>`, `<html>`, `<head>`, `<body>`, … in `layout/shell.templ` or `views/partials/shell_head.templ`.
- Tailwind utilities must pass **`Blank/.ui8px/policy`** (`bun run lint:ui8px`).
- Covered interaction: `@ui8kit/aria` + `templ/components` Sheet with `Behavior: "ui8kit"` — no custom JS for covered patterns.

## Related docs

- [`layout/README.md`](layout/README.md)
- [`components/README.md`](components/README.md)
- [`blocks/README.md`](blocks/README.md)
- [`widgets/README.md`](widgets/README.md)
- [`variants/README.md`](variants/README.md)
- [`utils/README.md`](utils/README.md)
- Architecture: [`.project/specs/next-shadcn-architecture.md`](../../.project/specs/next-shadcn-architecture.md)
- App rules: `@Blank/.cursor/rules/blank-ui-structure.mdc`, `blank-atomic-ui8px.mdc`

## Source policy

Registry folder rules match **FastyGoUI** design-system policy (`registry:*` labels).
When FastyGoUI is not in the workspace, **this README + subtree READMEs** are the on-disk index.
