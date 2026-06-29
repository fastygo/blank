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

**Not part of this registry:** `internal/views/` (route pages + co-located builders), `internal/site/` (runtime route manifest).

## Tree

```
internal/ui/
  layout/       # registry:layout — shells, data.go, build.go, shell_head, header_trailing
  components/   # registry:components — icon, toggles, navigation, appsidebar, …
  blocks/       # registry:blocks — full scaffolds (staging → fastygo/blocks)
    marketing/  # hero/ (live); add organisms as needed
  widgets/      # registry:widgets — staging stub (README + doc.go)
  variants/     # registry:variants — staging stub (README + doc.go)
  utils/        # registry:utils — thin helpers on templ/utils
```

| Label | Path | Role |
|--------|------|------|
| `registry:layout` | `layout/` | Document + chrome layers composed directly by pages (`RootLayout`, `TopnavLayout`, `DashboardLayout`); stays in app |
| `registry:blocks` | `blocks/<domain>/` | **Full scaffolds** with in-package default copy — not 4-line adapter wrappers |
| `registry:components` | `components/` | Props in, markup out; no HTTP, domain, or fetch |
| `registry:widgets` | `widgets/` | Fetch, state, orchestration; composes blocks/components |
| `registry:variants` | `variants/` | Named, ui8px-safe utility presets |
| `registry:utils` | `utils/` | App-specific helpers; generic logic stays in **templ/utils** |

There is **no** `internal/ui/elements`, **no** `internal/ui/ui/`, and **no** `internal/ui/blocks/layout/` (conflicts with `layout/` app chrome).

## Where does this UI go?

| You are building… | Put it in… |
|-------------------|------------|
| Document frame (`html`, `head`, `body`) | `layout/root_layout.templ` |
| Topnav chrome (header, footer, mobile sheet) | `layout/topnav_layout.templ` |
| Dashboard chrome (topnav + aside) | `layout/dashboard_layout.templ` |
| Small reusable control (icon, toggle) | `components/<area>/` |
| Aside / sidebar content consumed by `DashboardLayout` | `components/appsidebar/` (or fork it) |
| Full scaffold with default copy (dashboard, hero, docs toc) | `blocks/<domain>/<organism>/` |
| Shell that fetches or orchestrates API/state | `widgets/` |
| Route page (composes its own layout layers) | `internal/views/<page>.templ` |
| HTTP route + page choice | `internal/site/router.go` (`PageSpec`) |

Routes register `views.<Page>` from `internal/site/router.go`; each page composes
`@layout.RootLayout(d.Document()) { @layout.TopnavLayout(d.Topnav()) { ... } }` or
`@layout.RootLayout(d.Document()) { @layout.DashboardLayout(d.Dashboard(title)) { ... } }`.

## `components` vs `widgets` vs `blocks`

- **`blocks`** — section-level or layout-level markup + default demo data; portable wireframe (shadcn block copy-paste).
- **`components`** — small reusable pieces; props only.
- **`widgets`** — same UI boundaries plus **behavior** (API, loading, subscriptions).

If it only renders props → **`components`** or **`blocks`**. If it **fetches** or coordinates side effects → **`widgets`**.

## Runtime vs registry (do not collapse)

```text
internal/site/router.go  →  internal/views/<Page>  →  @layout.RootLayout  →  @layout.TopnavLayout|DashboardLayout  →  @ui.* atoms
```

- **`internal/site/`** — runtime wiring only; no reusable markup, no layout selection.
- **`internal/views/<page>.templ`** — composes its own layout shell. The composition tree is visible in one file (shadcn parity).
- **`internal/ui/layout/`** — named document shells used by pages directly.
- **`internal/ui/components/`** — small reusable parts and aside content (e.g. `appsidebar`).
- **`internal/ui/blocks/*`** — shadcn-style full scaffolds; copy-paste or `require` after extraction.

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
- Document shell only: `<!DOCTYPE>`, `<html>`, `<head>`, `<body>`, … in `layout/root_layout.templ` or `layout/shell_head.templ`.
- Tailwind utilities must pass **`Blank/.ui8px/policy`** (`bun run lint:ui8px`).
- Covered interaction: `@ui8kit/aria` + `templ/components` Sheet with `Behavior: "ui8kit"` — no custom JS for covered patterns.

## Related docs

- [`layout/README.md`](layout/README.md)
- [`components/README.md`](components/README.md)
- [`blocks/README.md`](blocks/README.md)
- [`widgets/README.md`](widgets/README.md)
- [`variants/README.md`](variants/README.md)
- [`utils/README.md`](utils/README.md)
- Architecture: [`docs/architecture.md`](../../docs/architecture.md)
- App rules: `@Blank/.cursor/rules/blank-ui-structure.mdc`, `blank-atomic-ui8px.mdc`

## Source policy

Registry folder rules match **FastyGoUI** design-system policy (`registry:*` labels).
When FastyGoUI is not in the workspace, **this README + subtree READMEs** are the on-disk index.
