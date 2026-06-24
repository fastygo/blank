# Blank Next/shadcn Refactor Progress

This file is a copy-paste progress driver for iterative Plan mode chats.
Copy one block at a time into a new chat or current chat in Plan mode.
Ask the agent to produce a bounded implementation plan from that block, then switch to Agent mode for the implementation slice.

## Fixed Decisions

- Blank should feel like **Next App Router + shadcn/ui** for React developers.
- The core mental model is **route -> route layout -> page content**.
- `fastygo.config.mjs` stays primarily a **Vite-like tooling config**: server env, templ, CSS, JS bundles, validation. It is not the route registry.
- `go.mod` should stay thin. Do not add `github.com/fastygo/blocks` or `github.com/fastygo/widgets` during Blank showcase/starter work.
- Runtime routes should first become a readable Go route manifest. `routes.yaml`/codegen is a later DX layer, not the first refactor.
- Layout work should follow the shadcn registry pattern: **copy-pasteable UI artifacts in `internal/ui/*` plus thin runtime wiring**, not one giant layout engine.
- Do not confuse **runtime app routing** with **registry accumulation**. Blank is a runtime app, but reusable UI mass must still land under `internal/ui/{components,blocks,widgets,layout,variants,utils}`.
- `internal/ui/layout/` is app-owned chrome/frame that stays in the app. Reusable layout organisms (sidebars, dashboard shells, docs toc shells, etc.) should be considered `blocks` or `widgets`, not hidden branches inside one `views/layout.templ`.
- Sidebars are reusable UI organisms that may include desktop-static regions and mobile-sheet behavior. They are not separate applications or branches.
- Custom JS is not allowed for covered interaction patterns. Use `@ui8kit/aria` through existing pattern hooks. Add custom JS only when a required W3C/APG pattern is not covered by `@ui8kit/aria`, and document that exception.

## Current Baseline To Re-check In Each Slice

- Runtime route specs live in `internal/site/router.go`; render, nav, and layout data helpers are split across `render.go`, `nav.go`, and `layout_data.go`. `feature.go` keeps feature wiring only.
- Route layout wrapper currently exists as `views.SiteShell` / `views.AppShell` in `internal/views/layout.templ`; this is a temporary runtime adapter, not the long-term registry home for every layout variant.
- Document/chrome shell currently exists as `layout.Shell` in `internal/ui/layout/shell.templ`.
- Mobile navigation currently uses `data-ui8kit` dialog/sheet hooks directly in the shell.
- `web/static/js/manifest.json` currently includes the `dialog` pattern.
- `github.com/fastygo/templ/components` already has `Sheet`, `SheetTrigger`, `SheetOverlay`, `SheetContent`, `SheetHeader`, `SheetTitle`, `SheetClose`.
- `@Templ/examples/ui/blocks/home` and `dashboard` already demonstrate the desired desktop aside + mobile sheet pairing.
- `internal/ui/README.md` and FastyGoUI registry policy define the target tree: `layout`, `components`, `blocks`, `widgets`, `variants`, `utils`. Do not introduce parallel `elements`, `ui`, `recipes`, or runtime-only layout catalogs.

## Validation Defaults

For implementation slices, prefer the narrowest useful check first:

- `.templ` or class string changes: `bun run lint:ui8px`
- ARIA hooks or JS manifest changes: `bun run build:js` then `bun run validate:aria`
- Route/view/model changes: `bun run templ` then `go test ./...`
- Broad layout refactor: `bun run verify`

If a command cannot be run, report why and what remains unverified.

---

## Progress Checklist

- [x] 00. Architecture spec and glossary are frozen. → [`.project/specs/next-shadcn-architecture.md`](specs/next-shadcn-architecture.md)
- [x] 01. Named route shells exist: `AppShell`, `MarketingShell`, `DocsShell`, with `SiteShell` as temporary alias.
- [x] 02. `internal/site` is split into runtime router manifest, render helper, nav, and layout data.
- [x] 03. React onboarding docs explain route -> shell -> page and the add-page workflow.
- [x] 04. Registry boundary is frozen: atoms/molecules/organisms/templates map to `Templ` and `internal/ui/*`.
- [x] 05. Current topnav shell is extracted from `views` into the chosen `internal/ui` registry location, with `views.AppShell` as thin adapter.
- [ ] 06. Mobile sheet/nav reusable UI is factored through existing `templ/components` Sheet APIs and `@ui8kit/aria`.
- [ ] 07. Sidebar/app layout organism is implemented under `internal/ui/blocks` or `internal/ui/widgets` as decided by Block 04.
- [ ] 08. Three sidebar wireframes are documented or showcased as registry artifacts, not runtime engine variants.
- [ ] 09. Runtime router chooses layout organisms explicitly without hiding layout choice in a global switch.
- [ ] 10. Final docs and tests are aligned for React/shadcn onboarding.

---

## Block 00 — Architecture Spec And Naming

Copy this block into Plan mode:

```text
We are refactoring @Blank for React/shadcn developers.

Goal:
Create a small architecture/spec slice that freezes the naming and responsibilities before implementation.

Context:
- Read @Blank/.project/shadcn-like.md
- Read @Blank/.project/sidebar-like.md
- Read @Blank/.project/vite-like.md
- Read @Blank/.project/next-shadcn-refactor-progress.md
- Re-check current files:
  - @Blank/internal/site/feature.go
  - @Blank/internal/views/layout.templ
  - @Blank/internal/ui/layout/shell.templ
  - @Blank/internal/ui/layout/README.md
  - @Blank/docs/for-react-devs.md

Required plan:
1. Define the final glossary:
   - document shell
   - route shell / route layout
   - page
   - chrome
   - registry artifact
   - layout organism
   - runtime adapter
   - block/showcase
2. Decide exact public names:
   - views.AppShell
   - views.MarketingShell
   - views.DocsShell
   - views.SiteShell as temporary compatibility alias
3. Decide what stays out of scope:
   - routes.yaml/codegen
   - custom JS
   - icon-collapse sidebar state
   - full layout engine
4. Propose the first implementation slice with files and validation.

Hard constraints:
- No custom JS for covered patterns. Use @ui8kit/aria.
- Do not add Go dependencies unless the plan proves they are required.
- Keep the plan small enough for one implementation pass.
```

Acceptance:

- A short durable spec exists or docs are updated with the glossary.
- The next implementation slice is unambiguous.

**Done:** See [`.project/specs/next-shadcn-architecture.md`](specs/next-shadcn-architecture.md) and Block 01 ready spec [`.project/specs/active.md`](specs/active.md).

---

## Block 01 — Named Route Shells

Copy this block into Plan mode:

```text
Implement the first code slice for Next-like named route layouts in @Blank.

Goal:
Make route-level layout selection readable to React developers by adding named shell wrappers:
- views.AppShell
- views.MarketingShell
- views.DocsShell
- views.SiteShell as compatibility alias to AppShell

Context:
- Read @Blank/internal/views/layout.templ
- Read @Blank/internal/views/models.go
- Read @Blank/internal/ui/layout/shell.templ
- Read @Blank/internal/site/feature.go
- Read @Blank/internal/views/wireframe_render_test.go
- Read @Blank/docs/for-react-devs.md

Design intent:
- internal/ui/layout.Shell remains the document/chrome shell.
- internal/views/*Shell are route layouts, analogous to Next route group layout.tsx files.
- AppShell initially preserves current topnav behavior.
- MarketingShell and DocsShell may initially delegate to AppShell or use minimal prop adjustments if that is the smallest safe slice.

Plan should include:
1. Exact changes to templ functions.
2. Whether render tests need rename/addition.
3. How docs should explain SiteShell vs AppShell.
4. Validation commands.

Hard constraints:
- No visual redesign in this slice.
- No sidebar implementation in this slice.
- No custom JS.
```

Acceptance:

- Current pages still render.
- New code can call `views.AppShell(data, body)`.
- Existing `views.SiteShell(data, body)` still works during migration.

**Done:** `internal/views/layout.templ`, render tests, [`docs/for-react-devs.md`](../docs/for-react-devs.md).

---

## Block 02 — Runtime Router Manifest And Thin Render Helper

Copy this block into Plan mode:

```text
Refactor @Blank/internal/site from handler-per-page boilerplate into a readable runtime router manifest.

Goal:
Make adding a page feel like a Next route table:
- runtime route specs live in internal/site/router.go (or routes.go if the implementation chooses that name deliberately)
- page render boilerplate lives in internal/site/render.go
- nav construction lives in internal/site/nav.go
- layout data construction lives in internal/site/layout_data.go
- Feature wiring stays in internal/site/feature.go

Context:
- Read @Blank/internal/site/feature.go
- Read @Blank/internal/views/models.go
- Read @Blank/internal/views/home.templ
- Read @Blank/internal/views/sample_stub.templ
- Read @Blank/internal/fixtures/fixtures.go
- Read @Blank/internal/fixtures/locale/en.json
- Read @Blank/internal/fixtures/locale/ru.json
- Read @Blank/vendor/github.com/fastygo/framework/pkg/app/feature.go

Desired model:
- Introduce a small PageSpec or equivalent.
- PageSpec should include method, pattern, active path, title resolver, explicit layout/render adapter, body renderer, and optional nav item resolver.
- Do not hide layout choice in `views/layout.templ`. The router should make runtime layout selection visible: route -> layout artifact/adapter -> page body.
- Feature.Routes should loop over specs and register GET routes.
- Custom handlers remain possible for future POST/API/auth flows.

Plan should decide:
1. Exact PageSpec shape.
2. How to avoid import cycles with views and templ.Component.
3. Whether the layout field points to `views.AppShell` temporarily or a typed adapter prepared for future `internal/ui/blocks` / `widgets` layout organisms.
4. How nav items are derived from route specs.
5. Whether siteNav remains as a helper or is generated from specs.
6. Which tests/docs must change.

Hard constraints:
- No routes.yaml or codegen yet.
- No behavior change for / and /sample.
- No custom JS.
```

Acceptance:

- New page workflow is: add fixture fields + locale JSON + view + one route spec.
- Runtime layout choice is visible in the route spec.
- `/` and `/sample` still render with active nav and localized copy.

**Done:** `internal/site/{feature.go,router.go,render.go,nav.go,layout_data.go}`.

---

## Block 03 — React Onboarding Docs

Copy this block into Plan mode:

```text
Update @Blank docs for React/shadcn onboarding after named shells and route manifest exist.

Goal:
Create a short, accurate onboarding path for a React developer who knows Next App Router and shadcn/ui.

Context:
- Read @Blank/README.md
- Read @Blank/docs/for-react-devs.md
- Read @Blank/.project/next-shadcn-refactor-progress.md
- Re-check the current implementation after previous slices.

Docs must explain:
1. Request flow: GET /about -> site route spec -> route shell -> page component.
2. File map:
   - app/layout.tsx -> internal/views/*Shell + internal/ui/layout.Shell
   - app/**/page.tsx -> internal/views/*.templ
   - components/ui -> github.com/fastygo/templ/ui and templ/components
   - app config/vite config -> fastygo.config.mjs
   - messages/*.json -> internal/fixtures/locale/*.json
3. Add-page cookbook.
4. Difference between route shell, chrome shell, blocks, components, widgets.
5. Honest dev loop: what watches automatically and what requires restart.

Hard constraints:
- Do not claim HMR or Go auto-restart if it is not implemented.
- Do not describe routes.yaml/codegen as available.
- No custom JS guidance except @ui8kit/aria pattern policy.
```

Acceptance:

- A React developer can add a basic page without reading framework internals.

**Done:** [`docs/for-react-devs.md`](../docs/for-react-devs.md), [`README.md`](../README.md).

---

## Block 04 — Registry Boundary For Layout Artifacts

Copy this block into Plan mode:

```text
Redesign the @Blank layout plan around the internal/ui registry model.

Goal:
Freeze where reusable layout-related UI artifacts live before adding sidebar/app shell variants.

Context:
- Read @Blank/.project/specs/next-shadcn-architecture.md
- Read @Blank/.project/sidebar-like.md
- Read @Blank/internal/ui/README.md
- Read @Blank/.cursor/rules/blank-ui-structure.mdc
- Read @FastyGoUI/.cursor/rules/fastygo-ui-design-system-registry.mdc
- Read @Blank/internal/ui/layout/shell.templ
- Read @Templ/components/sheet/sheet.spec.md if available in workspace.
- Read @Templ/.cursor/rules/templ-component-spec.mdc if available.

Required model:
- `github.com/fastygo/templ/ui` = atoms / primitive tags.
- `github.com/fastygo/templ/components` = molecules / neutral composites.
- `internal/ui/components/*` = small props-only app components.
- `internal/ui/blocks/*` = reusable section/layout organisms with in-package defaults, future `fastygo/blocks` candidates.
- `internal/ui/widgets/*` = organisms with behavior/data orchestration, future `fastygo/widgets` candidates.
- `internal/ui/layout/*` = app-owned document/chrome frame that stays in the app, not a dumping ground for every copy-paste layout variant.
- `internal/views/*` = pages and thin runtime adapters only; not a registry.
- `internal/site/router.go` (or chosen name) = runtime wiring: route -> resolved data -> selected UI artifact/layout adapter -> page.

Plan should decide:
1. Exact folder convention for layout organisms:
   - examples: `internal/ui/blocks/layout/topnav`, `internal/ui/blocks/sidebar/app`, `internal/ui/blocks/docs/toc_shell`, or another policy-consistent shape.
2. What remains in `internal/ui/layout` permanently.
3. What belongs in `blocks` vs `widgets` vs `components`.
4. How to describe this in `.project/specs/next-shadcn-architecture.md`.
5. Whether `views.AppShell` stays as a thin adapter or should later disappear behind route specs.

Hard constraints:
- This slice may be spec/docs only if that is safer.
- Do not introduce `internal/ui/recipes`, `internal/ui/elements`, or `internal/ui/ui`.
- Do not implement sidebar UI yet.
- No custom JS.
```

Acceptance:

- Future layout work has an explicit registry home.
- The team can distinguish runtime router wiring from shadcn-like copy-paste UI artifacts.

**Done:** [`.project/specs/next-shadcn-architecture.md`](specs/next-shadcn-architecture.md) (Registry boundary section), [`internal/ui/README.md`](../internal/ui/README.md), subtree READMEs, [`blank-ui-structure.mdc`](../.cursor/rules/blank-ui-structure.mdc).

**Frozen folder policy:** `blocks/<domain>/<organism>/` (e.g. `dashboard/sidebar_app`, `docs/toc_shell`) — **not** `blocks/layout/`.

---

## Block 05 — Extract Current Topnav Shell Into Registry Artifact

Copy this block into Plan mode:

```text
Move the current topnav app shell composition out of the central views facade and into the chosen internal/ui registry location.

Goal:
Make the current app shell a reusable UI artifact, while keeping `views.AppShell` as a thin runtime adapter for compatibility.

Context:
- Read @Blank/.project/specs/next-shadcn-architecture.md
- Read @Blank/internal/views/layout.templ
- Read @Blank/internal/views/models.go
- Read @Blank/internal/ui/layout/shell.templ
- Read @Blank/internal/ui/README.md
- Read the Block 04 registry boundary decision/spec.

Desired result:
- Current `AppShell` behavior lives in a named UI artifact under `internal/ui/*`, using the folder convention decided in Block 04.
- `views.AppShell` delegates to that artifact.
- `MarketingShell`, `DocsShell`, and `SiteShell` remain thin aliases/adapters until they diverge.
- No new layout variants are introduced in this slice.

Plan should include:
1. Exact folder and package names.
2. How to avoid import cycles between `views`, `internal/ui/*`, and view models.
3. Whether the UI artifact receives `layout.ShellProps`, a small app-shell props struct, or the existing `views.LayoutData`.
4. Render test updates.
5. Docs/progress updates.
6. Validation commands.

Hard constraints:
- No custom JS.
- Do not introduce sidebar, docs toc, or marketing divergence yet.
- Do not change visible layout beyond unavoidable package movement.
```

Acceptance:

- `views.AppShell(data, body)` still works.
- Current topnav shell is discoverable as an `internal/ui` registry artifact.
- Render tests still cover current shell behavior.

**Done:** [`internal/ui/blocks/dashboard/app_shell/`](../internal/ui/blocks/dashboard/app_shell/) (`appshell.AppShell`), [`internal/views/layout.templ`](../internal/views/layout.templ) delegates from `views.AppShell`.

---

## Block 06 — Reusable Mobile Sheet And Nav UI

Copy this block into Plan mode:

```text
Extract reusable mobile sheet and navigation UI for future sidebar/docs blocks.

Goal:
Stop treating mobile navigation as private shell markup. Create reusable UI under the correct registry directory so topnav, sidebar, docs toc, and other blocks can share it.

Context:
- Read @Blank/internal/ui/layout/shell.templ
- Read @Blank/internal/ui/layout/header.templ
- Read @Blank/internal/ui/layout/props.go
- Read @Blank/internal/ui/layout/helpers.go
- Read @Blank/internal/ui/layout/nav.templ
- Read @Blank/web/static/js/manifest.json
- Read @Blank/scripts/ui8kit-entry.mjs
- Read @Templ/components/sheet/sheet.templ if available.
- Read @Templ/examples/ui/blocks/home/page.templ if available.
- Read @Templ/examples/ui/blocks/dashboard/page.templ if available.
- Read the Block 04 registry boundary decision/spec.

Desired result:
- Use existing `templ/components` Sheet APIs with `Behavior: "ui8kit"` where possible.
- Preserve or deliberately migrate current IDs with tests:
  - ui8kit-mobile-sheet-panel
  - ui8kit-mobile-sheet-trigger
  - ui8kit-mobile-sheet-title
- Reusable nav/sheet UI lands under `internal/ui/components`, `blocks`, or `widgets` according to Block 04.
- `internal/ui/layout.Shell` consumes the reusable UI but remains the app-owned document/chrome frame.

Plan should include:
1. Exact component/block/widget boundaries.
2. How current topnav mobile menu maps to the reusable UI.
3. ARIA labels and fixture copy requirements.
4. Render test and aria validation plan.

Hard constraints:
- No custom JS.
- Do not introduce a new client state system.
- Do not change visible layout beyond unavoidable markup normalization.
```

Acceptance:

- Existing topnav mobile menu still works.
- The same reusable UI can later host sidebar or docs navigation.
- `bun run validate:aria` remains green.

---

## Block 07 — Sidebar App Organism

Copy this block into Plan mode:

```text
Implement the first sidebar app organism in the internal/ui registry.

Goal:
Create a reusable sidebar/app shell organism, not a runtime-only preset hidden inside `views/layout.templ`.

Context:
- Read @Blank/.project/sidebar-like.md
- Read @Blank/internal/ui/README.md
- Read @Blank/internal/ui/layout/shell.templ
- Read @Blank/internal/ui/layout/header.templ
- Read reusable mobile sheet/nav UI from Block 06.

Design:
- The sidebar app organism should live under the `internal/ui` location frozen in Block 04: `internal/ui/blocks/dashboard/sidebar_app` (or similar under `blocks/<domain>/<organism>/`).
- It composes existing atoms/molecules (`templ/ui`, `templ/components`) and app components/widgets.
- It may reuse `internal/ui/layout.Shell` or consume shell/chrome props from the adapter.
- It supports desktop sidebar + mobile sheet using `@ui8kit/aria`.

Plan should decide:
1. Exact package/folder name.
2. Props shape and default/demo data policy.
3. What belongs inside the block vs extracted component/widget.
4. How to keep current topnav shell working.
5. How to write focused render tests without overfitting class strings.

Hard constraints:
- No icon-collapse behavior yet.
- No custom JS.
- Use templ/ui and templ/components. No raw content/layout tags except documented shell exceptions.
```

Acceptance:

- Sidebar app organism exists as a reusable `internal/ui` artifact.
- It is not selected globally and does not replace current routes until Block 09.

---

## Block 08 — Sidebar Wireframe Showcase Artifacts

Copy this block into Plan mode:

```text
Document or implement showcase artifacts for the three sidebar wireframes in @Blank.

Goal:
Represent the three provided sidebar geometry examples as registry/showcase artifacts, not runtime engine variants.

Context:
- Read @Blank/.project/sidebar-like.md
- Read @Blank/internal/ui/layout/README.md
- Read @Blank/internal/ui/README.md
- Read the Block 04 registry boundary decision/spec.
- Re-check current sidebar app organism from Block 07.

Wireframes:
1. sidebars_full:
   - left aside scope: viewport
   - right aside scope: content_row
   - header scope: main_column
2. sidebars_main:
   - header scope: shell_full
   - footer scope: shell_full
   - left/right asides scope: content_row
3. sidebars_header:
   - shell header full width
   - page/header band below it
   - left/right asides in content row

Plan should decide:
1. Whether this slice is docs-only or includes renderable examples.
2. Where showcase examples should live.
3. How to avoid turning these examples into runtime engine variants.
4. Validation approach.

Hard constraints:
- These are examples/showcase artifacts, not mandatory runtime names.
- No custom JS.
- No brand polish; wireframe structure only.
```

Acceptance:

- The team can point to each image and explain it using regions, scopes, and registry artifacts.
- Future variants are built by composing/forking registry artifacts, not adding ad hoc shell forks.

---

## Block 09 — Runtime Wiring To UI Artifacts

Copy this block into Plan mode:

```text
Wire runtime routes to selected UI artifacts while preserving easy onboarding.

Goal:
Make route specs explicitly choose the app/marketing/docs layout artifact or adapter. Do not hide layout choice in one global switch.

Context:
- Read @Blank/internal/views/layout.templ
- Read @Blank/internal/views/models.go
- Read @Blank/internal/site/router.go or routes.go
- Read @Blank/internal/ui/layout/shell.templ
- Read the topnav artifact from Block 05.
- Read the sidebar app organism from Block 07.
- Read @Blank/internal/views/wireframe_render_test.go
- Read @Blank/docs/for-react-devs.md

Plan should decide:
1. Whether current pages remain topnav or one route demonstrates sidebar_app.
2. Whether `PageSpec.Layout` points to a view adapter, block renderer, or typed route layout function.
3. How a React developer will choose route layout in one route spec line.
4. Required fixture/locale copy for sidebar accessibility labels if sidebar is wired.
5. Tests and validation.

Hard constraints:
- No global hidden switch that makes routes hard to reason about.
- No custom JS.
- Do not break mobile sheet accessibility.
```

Acceptance:

- A route spec clearly shows which layout/UI artifact wraps the page.
- Sidebar_app behavior is discoverable from `internal/ui` registry files and route wiring.

---

## Block 10 — Tooling Config Cleanup

Copy this block into Plan mode:

```text
Review @Blank/fastygo.config.mjs and scripts after the routing/layout refactor.

Goal:
Make tooling config honest and useful for React/Vite developers without turning it into runtime routing/layout config.

Context:
- Read @Blank/fastygo.config.mjs
- Read @Blank/scripts/load-config.d.ts
- Read @Blank/scripts/dev.mjs
- Read @Blank/scripts/start.mjs
- Read @Blank/scripts/build.mjs
- Read @Blank/scripts/watch-css.mjs
- Read @Blank/docs/for-react-devs.md
- Read @Blank/README.md

Plan should decide:
1. Whether config typing needs app/tooling additions.
2. Whether docs overclaim watch/HMR behavior.
3. Whether dev script should be improved in a separate slice.
4. What must not go into config:
   - route table
   - runtime layout artifact selection
   - full layout engine

Hard constraints:
- Do not add runtime config if it creates a second source of truth.
- No custom JS for UI behavior.
```

Acceptance:

- `fastygo.config.mjs` is clearly documented as tooling config.
- React developers know what to edit for routes/layouts vs dev/build settings.

---

## Block 11 — Final Onboarding Verification

Copy this block into Plan mode:

```text
Perform final onboarding verification for the @Blank Next/shadcn refactor.

Goal:
Validate that a React/shadcn developer can understand and extend Blank quickly.

Context:
- Read @Blank/README.md
- Read @Blank/docs/for-react-devs.md
- Read @Blank/.project/next-shadcn-refactor-progress.md
- Read current:
  - internal/site/*
  - internal/views/*
  - internal/ui/*
  - internal/fixtures/*

Verification scenario:
1. Add a hypothetical /about page mentally.
2. Identify exact files to touch.
3. Confirm route shell choice is obvious.
4. Confirm copy/i18n workflow is obvious.
5. Confirm sidebar/mobile behavior is discoverable in `internal/ui` registry artifacts and does not require custom JS.
6. Confirm validation commands are documented.

Plan should produce:
1. Any final docs/test fixes needed.
2. Any naming cleanup.
3. Any old compatibility alias that can be removed or should remain.
4. A final checklist for maintainers.

Hard constraints:
- Do not introduce new architecture in this final slice.
- No custom JS.
```

Acceptance:

- Docs, code names, and tests tell the same story.
- The refactor is ready for normal feature work.

