# Active implementation spec — page composes layout

**Feature:** Collapse route-shell stack so `views/<route>.templ` shows the full layout tree in one file.
**Spec:** [page-composes-layout.md](./page-composes-layout.md)
**Supersedes:** [archive/block-11-final.md](./archive/block-11-final.md)

---

## Delivered

- [x] [`internal/ui/layout/sidebar_shell.templ`](../../internal/ui/layout/sidebar_shell.templ) — `layout.SidebarShell(shell, sidebar)`; promoted from the deleted `sidebar_app` block.
- [x] [`internal/ui/components/appsidebar/`](../../internal/ui/components/appsidebar/) — `appsidebar.AppSidebar(props)` local aside; analogue of shadcn `components/app-sidebar.tsx`, with its own [`README.md`](../../internal/ui/components/appsidebar/README.md).
- [x] [`internal/views/home.templ`](../../internal/views/home.templ) — composes `@layout.Shell(d.Shell) { ... }` directly; layout visible at the top of the file.
- [x] [`internal/views/sample_stub.templ`](../../internal/views/sample_stub.templ) — composes `@layout.SidebarShell(d.Shell, appsidebar.AppSidebar(d.Sidebar)) { ... }` directly.
- [x] [`internal/views/models.go`](../../internal/views/models.go) — `HomePageData` / `SamplePageData` carry `Shell layout.ShellProps` (and `Sidebar appsidebar.Props` for the sidebar page).
- [x] [`internal/views/layout_helpers.go`](../../internal/views/layout_helpers.go) — public `ShellPropsFor(d)` and `SidebarPropsFor(d, title)` builders.
- [x] `internal/views/layout.templ` **deleted** — route-adapter functions (`AppShell`, `SidebarAppShell`, `MarketingShell`, `DocsShell`) gone.
- [x] [`internal/site/router.go`](../../internal/site/router.go) — `PageSpec.Layout` removed; `Body PageRenderer` now returns a fully composed document.
- [x] [`internal/site/render.go`](../../internal/site/render.go) — one `web.Render(ctx, w, page.Body(data, fix))` call, no wrapping adapter.
- [x] `internal/ui/blocks/dashboard/{sidebar_app,app_shell}/` **deleted** — `blocks/dashboard/doc.go` kept as a domain stub for future full scaffolds.
- [x] [`internal/views/wireframe_render_test.go`](../../internal/views/wireframe_render_test.go) — `TestHomePage_composesTopnavShell` and `TestSamplePage_composesSidebarShell` exercise pages directly.
- [x] Rules + READMEs aligned: [`.cursor/rules/blank-ui-structure.mdc`](../../.cursor/rules/blank-ui-structure.mdc), [`internal/ui/README.md`](../../internal/ui/README.md), [`internal/ui/layout/README.md`](../../internal/ui/layout/README.md), [`internal/ui/blocks/README.md`](../../internal/ui/blocks/README.md), [`README.md`](../../README.md).
- [x] Docs + architecture spec aligned: [`docs/for-react-devs.md`](../../docs/for-react-devs.md), [`.project/specs/next-shadcn-architecture.md`](./next-shadcn-architecture.md). Historical brainstorms ([`shadcn-like.md`](../shadcn-like.md), [`sidebar-like.md`](../sidebar-like.md), [`next-shadcn-refactor-progress.md`](../next-shadcn-refactor-progress.md)) annotated as superseded.

---

## Verification (2026-06-25)

```bash
go mod vendor          # required once after fresh go/bun install (vendor/ was out of sync with go.mod)
bun run verify
```

Result: **pass** — `templ generate` (updates=0), `lint:ui8px` (0 violations), `validate:aria` (0 violations), `go test ./...` green including `TestHomePage_composesTopnavShell` and `TestSamplePage_composesSidebarShell`.

---

## Maintainer checklist (add a page, post-refactor)

1. Add route copy to [`internal/fixtures/fixtures.go`](../../internal/fixtures/fixtures.go) and **every** [`internal/fixtures/locale/*.json`](../../internal/fixtures/locale/) file.
2. Add `*PageData` to [`internal/views/models.go`](../../internal/views/models.go) — include `Shell layout.ShellProps` (and `Sidebar appsidebar.Props` if the page uses `SidebarShell`).
3. Add `internal/views/<page>.templ` — the template **composes** `@layout.Shell(d.Shell) { ... }` or `@layout.SidebarShell(d.Shell, appsidebar.AppSidebar(d.Sidebar)) { ... }`.
4. Add one [`PageSpec`](../../internal/site/router.go) in `internal/site/router.go` with `Body` returning the page directly.
5. Run **`bun run verify`** before landing the change.

There is **no** `Layout` field on `PageSpec` and **no** `views.*Shell` adapter — pages compose their own shell.

Shells live in [`internal/ui/layout/`](../../internal/ui/layout/):

- `layout.Shell` — topnav document chrome
- `layout.SidebarShell` — topnav + desktop aside slot

Local sidebar UI lives in [`internal/ui/components/appsidebar/`](../../internal/ui/components/appsidebar/).
