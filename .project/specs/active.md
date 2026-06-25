# Active implementation spec — registry co-location

**Feature:** Consolidate layout glue into `internal/ui/layout/`, co-locate page builders with each route, add first marketing block.
**Spec:** [page-composes-layout.md](./page-composes-layout.md) (prior slice) + registry co-location plan (2026-06-25).

---

## Delivered

- [x] [`internal/ui/layout/data.go`](../../internal/ui/layout/data.go) — `layout.Data`, `ShellProps()`, `SidebarProps()`.
- [x] [`internal/ui/layout/build.go`](../../internal/ui/layout/build.go) — `BuildData(BuildParams)`; asset + locale switch assembly.
- [x] [`internal/ui/layout/shell_head.templ`](../../internal/ui/layout/shell_head.templ), [`header_trailing.templ`](../../internal/ui/layout/header_trailing.templ) — moved from `views/partials/`.
- [x] [`internal/site/layout_data.go`](../../internal/site/layout_data.go) — thin wrapper calling `layout.BuildData`.
- [x] [`internal/site/router.go`](../../internal/site/router.go) — `PageRenderer func(layout.Data, fixtures.Locale)`; `Body: views.HomePageFrom` / `views.SamplePageFrom`.
- [x] [`internal/views/home.go`](../../internal/views/home.go) + [`home.templ`](../../internal/views/home.templ) — co-located builder + page; uses `@hero.Hero`.
- [x] [`internal/views/sample.go`](../../internal/views/sample.go) + [`sample.templ`](../../internal/views/sample.templ) — renamed from `sample_stub.templ`.
- [x] [`internal/ui/blocks/marketing/hero/`](../../internal/ui/blocks/marketing/hero/) — first copy-paste block extracted from home hero markup.
- [x] Removed `views/models.go`, `views/layout_helpers.go`, `views/partials/`, empty `blocks/dashboard/`, `blocks/docs/`, `blocks/marketing/doc.go`.
- [x] Tests: [`wireframe_render_test.go`](../../internal/views/wireframe_render_test.go), [`hero_render_test.go`](../../internal/ui/blocks/marketing/hero/hero_render_test.go).
- [x] Docs + rules aligned.

---

## Verification

```bash
bun run verify
```

Result: **pass** — templ, ui8px, validate:aria, `go test ./...` green.

---

## Maintainer checklist (add a page)

1. Add route copy to [`internal/fixtures/fixtures.go`](../../internal/fixtures/fixtures.go) and **every** [`internal/fixtures/locale/*.json`](../../internal/fixtures/locale/) file.
2. Add `internal/views/<page>.go` — `func <Page>From(d layout.Data, f fixtures.Locale) templ.Component`.
3. Add `internal/views/<page>.templ` — compose `@layout.Shell(d.ShellProps()) { ... }` or `@layout.SidebarShell(d.ShellProps(), appsidebar.AppSidebar(d.SidebarProps(title))) { ... }`.
4. Add one [`PageSpec`](../../internal/site/router.go) with `Body: views.<Page>From`.
5. Run **`bun run verify`**.

There is **no** central `models.go`, **no** `layout_helpers.go`, and **no** `PageSpec.Layout`.

Layout request data: [`internal/ui/layout/data.go`](../../internal/ui/layout/data.go) (`layout.Data`, `ShellProps()`, `SidebarProps()`).

Blocks: full scaffolds under [`internal/ui/blocks/<domain>/<organism>/`](../../internal/ui/blocks/) — e.g. [`marketing/hero`](../../internal/ui/blocks/marketing/hero/).
