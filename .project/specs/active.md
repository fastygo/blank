# Active implementation spec — root layout split

**Feature:** Split `Shell` into `RootLayout` + `TopnavLayout` + `DashboardLayout`; collapse `views.*From` to `views.<Page>(d, f)`.
**Prior spec:** [page-composes-layout.md](./page-composes-layout.md) (registry co-location slice).

---

## Delivered

- [x] [`internal/ui/layout/root_layout.templ`](../../internal/ui/layout/root_layout.templ) — document frame only (`html`, `head`, `body`).
- [x] [`internal/ui/layout/topnav_layout.templ`](../../internal/ui/layout/topnav_layout.templ) — header, main, footer, mobile sheet.
- [x] [`internal/ui/layout/dashboard_layout.templ`](../../internal/ui/layout/dashboard_layout.templ) — TopnavLayout + desktop aside.
- [x] [`internal/ui/layout/props.go`](../../internal/ui/layout/props.go) — `DocumentProps`, `TopnavLayoutProps`, `DashboardLayoutProps`.
- [x] [`internal/ui/layout/data.go`](../../internal/ui/layout/data.go) — `Document()`, `Topnav()`, `Dashboard(title)`.
- [x] [`internal/ui/layout/build*.go`](../../internal/ui/layout/) — split builders per layer.
- [x] Removed `shell.templ`, `sidebar_shell.templ`, `views/home.go`, `views/sample.go`.
- [x] [`internal/views/home.templ`](../../internal/views/home.templ), [`sample.templ`](../../internal/views/sample.templ) — `HomePage(d, f)` / `SamplePage(d, f)`.
- [x] [`internal/site/router.go`](../../internal/site/router.go) — `Body: views.HomePage` / `views.SamplePage`.
- [x] Tests: [`layout_render_test.go`](../../internal/ui/layout/layout_render_test.go), updated [`wireframe_render_test.go`](../../internal/views/wireframe_render_test.go).
- [x] Docs + rules aligned; [`.project/.chat/shell-layout.md`](../.chat/shell-layout.md) resolved.

---

## Verification

```bash
bun run verify
```

---

## Maintainer checklist (add a page)

1. Add route copy to [`internal/fixtures/fixtures.go`](../../internal/fixtures/fixtures.go) and **every** [`internal/fixtures/locale/*.json`](../../internal/fixtures/locale/) file.
2. Add `internal/views/<page>.templ` — `templ <Page>(d layout.Data, f fixtures.Locale)` composing `@layout.RootLayout` + `@layout.TopnavLayout` or `@layout.DashboardLayout`.
3. Add one [`PageSpec`](../../internal/site/router.go) with `Body: views.<Page>`.
4. Run **`bun run verify`**.

Layout request data: [`internal/ui/layout/data.go`](../../internal/ui/layout/data.go) (`layout.Data`, `Document()`, `Topnav()`, `Dashboard(title)`).
