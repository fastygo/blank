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
- Layouts should follow the shadcn pattern: **compound parts + small enum axes + presets/blocks**, not one giant layout engine.
- Sidebars are **regions** that may be desktop-static and mobile-sheet. They are not separate applications or branches.
- Custom JS is not allowed for covered interaction patterns. Use `@ui8kit/aria` through existing pattern hooks. Add custom JS only when a required W3C/APG pattern is not covered by `@ui8kit/aria`, and document that exception.

## Current Baseline To Re-check In Each Slice

- Routes, locale resolution, nav, layout data, and render handlers are currently mixed in `internal/site/feature.go`.
- Route layout wrapper currently exists as `views.SiteShell` in `internal/views/layout.templ`.
- Document/chrome shell currently exists as `layout.Shell` in `internal/ui/layout/shell.templ`.
- Mobile navigation currently uses `data-ui8kit` dialog/sheet hooks directly in the shell.
- `web/static/js/manifest.json` currently includes the `dialog` pattern.
- `github.com/fastygo/templ/components` already has `Sheet`, `SheetTrigger`, `SheetOverlay`, `SheetContent`, `SheetHeader`, `SheetTitle`, `SheetClose`.
- `@Templ/examples/ui/blocks/home` and `dashboard` already demonstrate the desired desktop aside + mobile sheet pairing.

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
- [ ] 01. Named route shells exist: `AppShell`, `MarketingShell`, `DocsShell`, with `SiteShell` as temporary alias.
- [ ] 02. `internal/site` is split into route manifest, render helper, nav, and layout data.
- [ ] 03. React onboarding docs explain route -> shell -> page and the add-page workflow.
- [ ] 04. `internal/ui/layout/layout.spec.md` defines compound layout parts and enum axes.
- [ ] 05. Mobile sheet/trigger parts use existing `templ/components` Sheet APIs and `@ui8kit/aria`.
- [ ] 06. `AsideRegion`, `MainInset`, and `ShellBand` exist as reusable layout parts.
- [ ] 07. `topnav` and `sidebar_app` presets are implemented.
- [ ] 08. Three sidebars wireframes are documented or showcased as preset examples, not core runtime names.
- [ ] 09. Sidebar app preset is wired into `AppShell` without breaking topnav pages.
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
   - layout part
   - preset
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

---

## Block 02 — Route Manifest And Thin Render Helper

Copy this block into Plan mode:

```text
Refactor @Blank/internal/site from handler-per-page boilerplate into a readable route manifest.

Goal:
Make adding a page feel like a Next route table:
- routes live in internal/site/routes.go
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
- PageSpec should include method, pattern, active path, title resolver, shell renderer, body renderer, and optional nav item resolver.
- Feature.Routes should loop over specs and register GET routes.
- Custom handlers remain possible for future POST/API/auth flows.

Plan should decide:
1. Exact PageSpec shape.
2. How to avoid import cycles with views and templ.Component.
3. How nav items are derived from route specs.
4. Whether siteNav remains as a helper or is generated from specs.
5. Which tests/docs must change.

Hard constraints:
- No routes.yaml or codegen yet.
- No behavior change for / and /sample.
- No custom JS.
```

Acceptance:

- New page workflow is: add fixture fields + locale JSON + view + one route spec.
- `/` and `/sample` still render with active nav and localized copy.

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

---

## Block 04 — Layout Spec For shadcn-like Parts

Copy this block into Plan mode:

```text
Design the @Blank/internal/ui/layout component spec for shadcn-like layout parts.

Goal:
Create or update internal/ui/layout/layout.spec.md as the source of truth for layout parts and enum axes.

Context:
- Read @Blank/.project/sidebar-like.md
- Read @Blank/internal/ui/layout/shell.templ
- Read @Blank/internal/ui/layout/props.go
- Read @Blank/internal/ui/layout/helpers.go
- Read @Blank/internal/ui/layout/README.md
- Read @Templ/components/sheet/sheet.spec.md if available in workspace.
- Read @Templ/.cursor/rules/templ-component-spec.mdc if available.

Required model:
- Do not make one giant Variant.
- Define compound parts:
  - Shell
  - ShellBand
  - MainInset
  - AsideRegion
  - MobileSheet
  - SheetTrigger
  - NavList or SidebarNav
  - optional PageHeader/AppFooter
- Define enum axes:
  - AsideSide: left | right
  - AsideScope: viewport | content_row | main_column
  - AsideDesktopMode: static | sticky | hidden
  - AsideMobileMode: sheet | hidden | inline
  - Collapsible: none | offcanvas | icon
  - BandScope: shell_full | main_column | content_row
  - TriggerPlacement: header_start | header_end | main_start | main_end

First implementation should support only:
- Collapsible: none | offcanvas
- AsideMobileMode: sheet | hidden

Hard constraints:
- This slice may be spec/docs only if that is safer.
- Do not implement icon-collapse yet.
- No custom JS.
- Any mobile sheet behavior must use @ui8kit/aria via existing Sheet/dialog hooks.
```

Acceptance:

- Future layout code has a bounded API contract.
- The three sidebar images are described as showcase/preset examples, not as the whole architecture.

---

## Block 05 — MobileSheet And SheetTrigger Parts

Copy this block into Plan mode:

```text
Extract reusable MobileSheet and SheetTrigger layout parts in @Blank using existing templ/components Sheet APIs.

Goal:
Stop hand-rolling mobile sheet markup inside layout.Shell. Create reusable layout parts that can be used by topnav, left sidebar, right sidebar, docs toc, etc.

Context:
- Read @Blank/internal/ui/layout/shell.templ
- Read @Blank/internal/ui/layout/header.templ
- Read @Blank/internal/ui/layout/helpers.go
- Read @Blank/internal/ui/layout/props.go
- Read @Blank/web/static/js/manifest.json
- Read @Blank/scripts/ui8kit-entry.mjs
- Read @Templ/components/sheet/sheet.templ if available.
- Read @Templ/examples/ui/blocks/home/page.templ if available.
- Read @Templ/examples/ui/blocks/dashboard/page.templ if available.

Desired result:
- Introduce MobileSheetProps and SheetTrigger props/helpers inside internal/ui/layout.
- Use import shape `import cmp "github.com/fastygo/templ/components"` for Sheet parts.
- Preserve current IDs or migrate carefully with tests:
  - ui8kit-mobile-sheet-panel
  - ui8kit-mobile-sheet-trigger
  - ui8kit-mobile-sheet-title
- Preserve @ui8kit/aria dialog/sheet behavior and manifest expectations.

Plan should include:
1. Exact part names and files.
2. How current topnav mobile menu maps to the new parts.
3. ARIA labels and fixture copy requirements.
4. Render test updates.
5. Validation commands.

Hard constraints:
- No custom JS.
- Do not introduce a new client state system.
- Do not change visible layout beyond unavoidable markup normalization.
```

Acceptance:

- Existing topnav mobile menu still works.
- The same MobileSheet part can later host sidebar nav.
- `bun run validate:aria` remains green.

---

## Block 06 — AsideRegion, MainInset, ShellBand

Copy this block into Plan mode:

```text
Implement the core shadcn-like layout geometry parts in @Blank.

Goal:
Add reusable layout parts:
- ShellBand
- MainInset
- AsideRegion

Context:
- Read @Blank/internal/ui/layout/layout.spec.md
- Read @Blank/internal/ui/layout/shell.templ
- Read @Blank/internal/ui/layout/props.go
- Read @Blank/internal/ui/layout/helpers.go
- Read @Blank/internal/ui/layout/nav.templ
- Read @Blank/internal/ui/layout/header.templ
- Read @Blank/.project/sidebar-like.md

Design:
- AsideRegion is a desktop region that can be left or right.
- AsideRegion can also declare a mobile sheet, but mobile rendering should use MobileSheet from the previous slice.
- MainInset is the main column/content wrapper, like shadcn SidebarInset.
- ShellBand hosts header/footer/subheader in shell_full or main_column scope.

Plan should decide:
1. Minimal props for each part.
2. Which class computations belong in helpers.go.
3. How to keep current topnav shell working.
4. How to write or adjust render tests without overfitting class strings.

Hard constraints:
- No icon-collapse behavior yet.
- No custom JS.
- Use templ/ui and templ/components. No raw content/layout tags except documented shell exceptions.
```

Acceptance:

- Existing topnav can be expressed through the new parts or can coexist while parts are introduced.
- Parts are generic enough for left/right/sidebar/docs layouts.

---

## Block 07 — Presets: topnav And sidebar_app

Copy this block into Plan mode:

```text
Add first production layout presets to @Blank.

Goal:
Implement preset factories for:
- topnav: current Blank behavior
- sidebar_app: desktop left sidebar + mobile sheet + main header trigger

Context:
- Read @Blank/internal/views/layout.templ
- Read @Blank/internal/views/models.go
- Read @Blank/internal/ui/layout/props.go
- Read @Blank/internal/ui/layout/shell.templ
- Read @Blank/internal/ui/layout/layout.spec.md
- Read @Blank/internal/site/routes.go if it exists
- Read @Blank/.project/sidebar-like.md

Design:
- Presets are Go/templ factory functions, not a global layout engine.
- `AppShell` chooses one preset.
- `MarketingShell` remains no-sidebar/topnav/minimal.
- `DocsShell` can remain a placeholder until docs routes exist.
- Power users can bypass presets later and assemble ShellProps manually.

Plan should decide:
1. Where preset factories live.
2. Whether preset is selected by a typed constant, config value, or direct wrapper function.
3. How `LayoutData.NavItems` maps into sidebar content and mobile sheet.
4. Render test coverage for topnav and sidebar_app.

Hard constraints:
- Do not wire `fastygo.config.mjs` as a runtime layout switch unless the plan includes a clean runtime bridge.
- No custom JS.
- Preserve current pages.
```

Acceptance:

- `topnav` still renders current shell behavior.
- `sidebar_app` can render a left desktop sidebar and mobile sheet from the same nav data.

---

## Block 08 — Sidebar Showcase Presets For Three Wireframes

Copy this block into Plan mode:

```text
Document or implement showcase presets for the three sidebar wireframes in @Blank.

Goal:
Represent the three provided sidebar geometry examples as showcase/preset examples, not as the core layout model.

Context:
- Read @Blank/.project/sidebar-like.md
- Read @Blank/internal/ui/layout/layout.spec.md
- Read @Blank/internal/ui/layout/README.md
- Re-check current layout parts and presets.

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
3. How to avoid turning these examples into a layout engine.
4. Validation approach.

Hard constraints:
- These are examples/presets, not mandatory runtime names.
- No custom JS.
- No brand polish; wireframe structure only.
```

Acceptance:

- The team can point to each image and explain it using parts and scopes.
- Future variants are built by composing parts, not adding ad hoc shell forks.

---

## Block 09 — AppShell Sidebar Wiring

Copy this block into Plan mode:

```text
Wire sidebar_app into AppShell in @Blank while preserving easy onboarding.

Goal:
Make AppShell capable of using sidebar_app as the main application layout.

Context:
- Read @Blank/internal/views/layout.templ
- Read @Blank/internal/views/models.go
- Read @Blank/internal/site/routes.go
- Read @Blank/internal/ui/layout/presets.go if it exists
- Read @Blank/internal/ui/layout/shell.templ
- Read @Blank/internal/views/wireframe_render_test.go
- Read @Blank/docs/for-react-devs.md

Plan should decide:
1. Whether AppShell defaults to topnav or sidebar_app at this milestone.
2. Whether MarketingShell is used for home and AppShell for sample/dashboard, or both current pages stay under AppShell.
3. How a React developer will choose route layout in PageSpec.
4. Required fixture/locale copy for sidebar accessibility labels.
5. Tests and validation.

Hard constraints:
- No global hidden switch that makes routes hard to reason about.
- No custom JS.
- Do not break mobile sheet accessibility.
```

Acceptance:

- A route spec clearly shows `Shell: views.AppShell` or `Shell: views.MarketingShell`.
- Sidebar_app behavior is discoverable without reading internal shell code.

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
   - route shell selection
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
  - internal/ui/layout/*
  - internal/fixtures/*

Verification scenario:
1. Add a hypothetical /about page mentally.
2. Identify exact files to touch.
3. Confirm route shell choice is obvious.
4. Confirm copy/i18n workflow is obvious.
5. Confirm sidebar/mobile behavior does not require custom JS.
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

