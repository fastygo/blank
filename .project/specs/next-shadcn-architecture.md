# Blank Next/shadcn Architecture

Durable architecture spec for the Blank refactor aimed at React developers who know **Next App Router** and **shadcn/ui**. This document freezes vocabulary and responsibilities before runtime changes.

**Status:** Block 03 complete — React onboarding docs updated.  
**Next slice:** Block 04 — registry boundary for layout artifacts (see [active.md](./active.md)).

---

## Intent

Blank is a **server-rendered BFF frontend** (Go + templ + Tailwind). React developers should onboard using the same mental model as Next:

```text
route -> route layout -> page content
```

They should **not** need to read framework internals to add a page, choose a layout, or understand where chrome vs page markup lives.

---

## Glossary

| Term | Blank location | Role |
|------|----------------|------|
| **Document shell** | `internal/ui/layout.Shell` | Full HTML document: `DOCTYPE`, `html`, `head`, `body`, asset hooks, top-level chrome composition, `@ui8kit/aria` markup hooks. Analogous to root `app/layout.tsx` document frame only — not route-group choice. |
| **Route shell / route layout** | `views.AppShell`, `views.MarketingShell`, `views.DocsShell` | Temporary runtime adapters that wrap `{children}` page body. Analogous to `app/(app)/layout.tsx`, `(marketing)/layout.tsx`, `(docs)/layout.tsx`; later blocks move reusable UI mass into `internal/ui/*`. |
| **Page** | `internal/views/*.templ` (e.g. `HomePage`, `SamplePage`) | Route content only. Receives resolved props/strings. No fixture loading inside `.templ`. No header/footer/sidebar chrome. Analogous to `page.tsx`. |
| **Chrome** | `internal/ui/layout/*`, `internal/ui/components/*` | App-owned shell/frame UI: document shell, header, footer, nav, mobile sheet hosts, theme/language controls. `internal/ui/layout` stays in the app. |
| **Registry artifact** | `internal/ui/{components,blocks,widgets,variants,utils}` | Copy-pasteable or reusable UI mass accumulated in the app registry before extraction. This is the shadcn-like layer. |
| **Layout organism** | Usually `internal/ui/blocks/*` or `internal/ui/widgets/*` | Reusable composed UI such as sidebar app shell, dashboard shell, docs toc shell. It may compose `templ/ui`, `templ/components`, app components, and widgets. |
| **Block / showcase** | `internal/ui/blocks/*`, `@Templ/examples/ui/blocks/*` | Copy-paste wireframe scaffolds (shadcn Blocks). **Not** the runtime router itself. Runtime routes may choose adapters that render these artifacts. |

### Two “layout” layers (common confusion)

| Layer | Name in docs | File today | Next analogy |
|-------|--------------|------------|--------------|
| 1 | Document shell | `internal/ui/layout/shell.templ` → `layout.Shell` | Root document + providers frame |
| 2 | Route shell adapter | `internal/views/layout.templ` → `views.SiteShell` (→ `AppShell`) | Temporary route group adapter |

**Rule:** When onboarding says “layout”, specify **document shell** vs **route shell**.

---

## Next / shadcn mapping

| Next App Router | shadcn | Blank (target) |
|-----------------|--------|----------------|
| `next.config.ts` / `vite.config.ts` | — | [`fastygo.config.mjs`](../../fastygo.config.mjs) — **tooling only** |
| Root document | — | `internal/ui/layout.Shell` |
| `app/(app)/layout.tsx` | `SidebarProvider` + inset | `views.AppShell` |
| `app/(marketing)/layout.tsx` | minimal header | `views.MarketingShell` |
| `app/(docs)/layout.tsx` | toc + content | `views.DocsShell` |
| `app/**/page.tsx` | block content | `internal/views/<page>.templ` |
| `components/ui/*` | primitives | `github.com/fastygo/templ/ui` + `templ/components` |
| App components | local UI | `internal/ui/components/*` |
| shadcn blocks | copy-paste | `internal/ui/blocks/*` + Templ examples |
| Route table | file-based routes | `internal/site/router.go` (target runtime manifest; `routes.go` acceptable if named deliberately; today `feature.go`) |
| `middleware.ts` | — | `internal/serverapp/app.go` (locales, security) |
| `messages/*.json` | — | `internal/fixtures/locale/*.json` |

---

## Request flow

```mermaid
flowchart LR
  subgraph request [HTTP request]
    R["GET /sample"]
  end

  subgraph site [internal/site]
    RF["router.go / feature.go"]
    LD["layout_data.go"]
  end

  subgraph views [internal/views]
    RS["AppShell / MarketingShell / DocsShell"]
    PG["SamplePage"]
  end

  subgraph chrome [internal/ui]
    DS["layout.Shell"]
    LP["registry artifacts"]
  end

  R --> RF
  RF --> LD
  RF --> RS
  RS --> DS
  DS --> LP
  RS --> PG
```

**Onboarding rule:** React devs touch **`internal/views/*.templ`** and **`fixtures/locale/*.json`** daily; **`internal/site/`** stays thin runtime routing; reusable UI mass lives in **`internal/ui/*`**.

---

## Public naming (frozen)

| Symbol | Status | Meaning |
|--------|--------|---------|
| `views.AppShell` | **Temporary adapter** | App-zone route layout adapter (topnav today; delegates to `internal/ui` artifact later). |
| `views.MarketingShell` | **Temporary adapter** | Public/landing layout adapter. |
| `views.DocsShell` | **Temporary adapter** | Docs layout adapter. |
| `views.SiteShell` | **Temporary alias** | Delegates to `AppShell` until call sites migrate. |
| `layout.Shell` | **Keep** | Document/chrome shell — do not rename to avoid colliding with route shells. |

---

## Layer responsibilities

### `internal/site/`

- HTTP route registration
- Locale resolution per request
- Building `views.LayoutData` from fixtures
- Choosing route layout adapter / UI artifact + page component
- **Must not** contain page markup

### `internal/views/`

- Route shells (`*Shell`)
- Page templates (`*Page`)
- View models (`models.go`)
- **Must not** load fixtures inside `.templ`

### `internal/ui/*`

- `layout/`: app-owned document/chrome frame that stays in the app
- `components/`: small props-only app UI
- `blocks/`: reusable section/layout organisms with default demo data
- `widgets/`: reusable UI with behavior/data orchestration
- `variants/` and `utils/`: named class maps and thin helpers
- Preserves `data-ui8kit-*` / `@ui8kit/aria` contracts where behavior is needed

### `internal/fixtures/`

- Embedded locale JSON + typed `Locale`
- User-visible copy including ARIA labels for chrome

### `fastygo.config.mjs`

- Dev/build tooling: server env, templ, CSS, JS bundles, ui8px validation
- **Not** the route registry or primary layout switch for production routes

---

## Interaction policy (no custom JS)

- Covered patterns use **`@ui8kit/aria`** via committed `web/static/js/ui8kit.js` and [`manifest.json`](../../web/static/js/manifest.json).
- Mobile navigation/sheets use existing **dialog/sheet** hooks (`data-ui8kit`, `data-ui8kit-dialog-*`) or `templ/components` Sheet with `Behavior: "ui8kit"`.
- **`theme.js`** is allowed for theme toggle (existing app behavior).
- Do **not** add custom client state for sidebar collapse, cookie persistence, or keyboard shortcuts until a spec explicitly requires it and `@ui8kit/aria` does not cover the pattern.
- New manifest patterns require **`bun run build:js`** and **`bun run validate:aria`**.

---

## Non-goals (this refactor track)

| Item | Reason |
|------|--------|
| `routes.yaml` + codegen | Later DX layer; Go route manifest first |
| Custom JS for sidebar/sheet | Use `@ui8kit/aria` + templ Sheet |
| Icon-collapse sidebar + cookie state | shadcn parity deferred; wireframe phase |
| Full layout engine / JSON layout config | Use registry artifacts + explicit runtime route wiring |
| Split `site` / `app` / `docs` features | Only when multiple layout groups need different nav |
| `github.com/fastygo/blocks` / `widgets` deps | Staging stays in `internal/ui/*` |
| Go auto-restart in dev | Document honest restart-after-Go-change workflow |

---

## Sidebar direction (summary)

Sidebars are **registry artifacts / organisms**, not separate repos or branches:

- **Desktop:** static `aside` in layout grid
- **Mobile:** same nav content in `MobileSheet` + `SheetTrigger`
- **Geometry:** represented by the specific block/widget composition, not a global runtime engine
- **Three wireframe images** are **showcase artifacts** (`sidebars_full`, `sidebars_main`, `sidebars_header`), not core runtime names

Reference implementations: `@Templ/examples/ui/blocks/home/page.templ`, `dashboard/page.templ`.

---

## One-line onboarding (target README phrase)

> **Runtime routes live in `internal/site/router.go`. Pages live in `internal/views/*.templ`. Route layout adapters are `AppShell` / `MarketingShell` / `DocsShell` for now. Reusable UI artifacts live in `internal/ui/*`. Copy lives in `fixtures/locale/*.json`.**

Runtime route specs live in `internal/site/router.go`; `feature.go` registers handlers from the manifest.

---

## Refactor sequence (progress driver)

See [next-shadcn-refactor-progress.md](../next-shadcn-refactor-progress.md) for copy-paste blocks.

| Block | Deliverable |
|-------|-------------|
| 00 | This spec |
| 01 | Named route shells + `SiteShell` alias |
| 02 | `router.go` + render helper |
| 03 | React onboarding docs |
| 04–09 | Registry boundary, UI artifacts, sidebar organism, runtime wiring |
| 10–11 | Tooling honesty + final verification |

---

## Block 00 completion notes

- **Completed:** Architecture glossary, naming, non-goals, and request-flow documented.
- **Baseline unchanged:** No runtime code modified in Block 00.
- **Next:** Implement Block 01 per [active.md](./active.md).

## Block 02 completion notes

- **Completed:** `PageSpec` route manifest in `internal/site/router.go`; `handlePage` in `render.go`; nav derived from specs in `nav.go`; layout data in `layout_data.go`; `feature.go` wiring only.
- **Layout adapter:** Current routes visibly use `views.AppShell` in route specs.
- **Next:** Block 03 — update React onboarding docs to reference the route manifest workflow.

## Block 03 completion notes

- **Completed:** [`docs/for-react-devs.md`](../../docs/for-react-devs.md) and [`README.md`](../../README.md) explain request flow, file map, add-page cookbook, registry terms, and honest dev loop.
- **No runtime changes:** Documentation-only slice.
- **Next:** Block 04 — freeze registry boundary for layout artifacts.
