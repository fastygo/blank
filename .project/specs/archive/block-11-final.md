# Active implementation spec — Block 11 (final)

**Feature:** Final onboarding verification  
**Architecture:** [next-shadcn-architecture.md](./next-shadcn-architecture.md)  
**Progress block:** Block 11 in [next-shadcn-refactor-progress.md](../next-shadcn-refactor-progress.md)

---

## Intent

Validate that a React/shadcn developer can understand and extend Blank quickly. No new architecture, routing model, custom JS, or UI behavior in this slice.

---

## Delivered

- [x] [`README.md`](../../README.md) — topnav home + sidebar sample routes; normal feature work checklist.
- [x] [`docs/for-react-devs.md`](../../docs/for-react-devs.md) — `/about` cookbook with exact file touchpoints; sidebar/mobile discoverability in `internal/ui`.
- [x] [`internal/ui/blocks/dashboard/sidebar_app/README.md`](../../internal/ui/blocks/dashboard/sidebar_app/README.md) — `/sample` uses `views.SidebarAppShell`; other routes choose `PageSpec.Layout` explicitly.
- [x] Removed `views.SiteShell` compatibility alias from [`internal/views/layout.templ`](../../internal/views/layout.templ); removed `TestSiteShell_aliasRenders`.

**Block 11 complete.** Validation: `bun run templ`, `bun run lint:ui8px`, `bun run validate:aria`, `go test ./...`, `bun run verify`.

---

## Refactor complete (Blocks 00–11)

Next/shadcn refactor is complete. Use the checklist below for normal feature work.

---

## Maintainer checklist (add a page)

1. Add route copy to [`internal/fixtures/fixtures.go`](../../internal/fixtures/fixtures.go) and **every** [`internal/fixtures/locale/*.json`](../../internal/fixtures/locale/) file.
2. Add page data/template under [`internal/views/`](../../internal/views/) (`models.go` + `<page>.templ`).
3. Add one [`PageSpec`](../../internal/site/router.go) in `internal/site/router.go`.
4. Choose layout in `PageSpec.Layout`:
   - `views.AppShell` — topnav app zone
   - `views.SidebarAppShell` — desktop aside + mobile sheet
   - `views.MarketingShell` — public/landing (currently delegates to `AppShell`)
   - `views.DocsShell` — docs zone (currently delegates to `AppShell`)
5. Reuse registry artifacts under [`internal/ui/`](../../internal/ui/) for chrome and sections — **not** `fastygo.config.mjs`.
6. Run **`bun run verify`** before landing the change.

Sidebar and mobile sheet UI: [`internal/ui/blocks/dashboard/sidebar_app/`](../../internal/ui/blocks/dashboard/sidebar_app/) + [`internal/ui/components/navigation/`](../../internal/ui/components/navigation/). `@ui8kit/aria` is already wired for covered patterns.
