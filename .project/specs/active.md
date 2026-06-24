# Active implementation spec — Block 02

**Feature:** Runtime router manifest and thin render helper  
**Architecture:** [next-shadcn-architecture.md](./next-shadcn-architecture.md)  
**Progress block:** Block 02 in [next-shadcn-refactor-progress.md](../next-shadcn-refactor-progress.md)

---

## Intent

Refactor `internal/site/feature.go` from handler-per-page boilerplate into an explicit runtime route manifest.

This slice is about **runtime wiring only**:

```text
route -> resolved locale/layout data -> layout adapter -> page body
```

It does **not** move layout UI into `internal/ui` yet. `views.AppShell`, `MarketingShell`, and `DocsShell` are temporary route layout adapters until later blocks move reusable UI mass into the `internal/ui` registry.

---

## Current baseline

- `internal/site/feature.go` currently mixes:
  - route registration;
  - fixture locale resolution;
  - nav construction;
  - language switch construction;
  - layout data construction;
  - page body creation;
  - `web.Render`.
- Current routes:
  - `GET /{$}` -> `views.SiteShell(..., views.HomePage(...))`
  - `GET /sample` -> `views.SiteShell(..., views.SamplePage(...))`
- `views.SiteShell` is a compatibility alias to `views.AppShell`.

---

## Target files

| File | Role |
|------|------|
| [`internal/site/feature.go`](../../internal/site/feature.go) | Keep `Feature`, constructor, `ID`, `NavItems`, and `Routes` wiring only |
| `internal/site/router.go` | Runtime route manifest and `PageSpec` definitions |
| `internal/site/render.go` | `handlePage` / render helper |
| `internal/site/nav.go` | Navigation helpers derived from specs or existing helper |
| `internal/site/layout_data.go` | `layoutData`, `navigationProps`, `languageSwitch`, asset constants if useful |

The implementation may choose `routes.go` instead of `router.go` only if it keeps the runtime route-manifest intent clear.

---

## Proposed types

Use small typed function aliases to keep route specs readable and future-compatible with `internal/ui` layout artifacts:

```go
type LayoutRenderer func(views.LayoutData, templ.Component) templ.Component
type TitleResolver func(fixtures.Locale) string
type BodyRenderer func(fixtures.Locale) templ.Component
type NavResolver func(fixtures.Locale) (layout.NavItem, bool)

type PageSpec struct {
    Method  string
    Pattern string
    Active  string
    Layout  LayoutRenderer
    Title   TitleResolver
    Body    BodyRenderer
    Nav     NavResolver
}
```

Notes:

- `Layout` points to `views.AppShell` / `MarketingShell` / `DocsShell` for now.
- This is intentionally a **runtime adapter field**, not the final home of reusable layout UI.
- Later blocks may point the adapter to an `internal/ui/blocks` or `widgets` artifact without changing the route mental model.

---

## Initial route specs

Expected shape:

```go
var pages = []PageSpec{
    {
        Method:  "GET",
        Pattern: "/{$}",
        Active:  "/",
        Layout:  views.AppShell,
        Title:   func(f fixtures.Locale) string { return f.Home.Title },
        Body: func(f fixtures.Locale) templ.Component {
            return views.HomePage(views.HomeData{
                Welcome:      f.Home.Welcome,
                WelcomeBrand: f.Home.WelcomeBrand,
                Description:  f.Home.Description,
            })
        },
        Nav: func(f fixtures.Locale) (layout.NavItem, bool) {
            return layout.NavItem{Label: f.Home.NavLabel, Path: "/", Icon: "home"}, true
        },
    },
    {
        Method:  "GET",
        Pattern: "/sample",
        Active:  "/sample",
        Layout:  views.AppShell,
        Title:   func(f fixtures.Locale) string { return f.Sample.Title },
        Body: func(f fixtures.Locale) templ.Component {
            return views.SamplePage(views.SampleData{
                Title:       f.Sample.Title,
                Description: f.Sample.Description,
                Body:        f.Sample.Body,
            })
        },
        Nav: func(f fixtures.Locale) (layout.NavItem, bool) {
            return layout.NavItem{Label: f.Sample.NavLabel, Path: "/sample", Icon: "box"}, true
        },
    },
}
```

The exact names may vary, but the route spec must visibly show layout selection.

---

## Render flow

`Feature.Routes` should become a loop:

```go
func (f *Feature) Routes(mux *http.ServeMux) {
    for _, page := range pages {
        page := page
        mux.HandleFunc(page.Method+" "+page.Pattern, f.handlePage(page))
    }
}
```

`handlePage` should:

1. resolve `ctx`;
2. load fixture locale;
3. build `LayoutData`;
4. call `web.Render(ctx, w, page.Layout(data, page.Body(fix)))`.

---

## Navigation

Prefer deriving nav from `PageSpec.Nav` so path/label/icon are not duplicated.

If this creates too much complexity, keep `siteNav(fix)` as a helper for this slice, but document it as temporary and avoid duplicating route paths in more than one new place.

---

## Out of scope

- No `routes.yaml` or codegen.
- No `internal/ui` registry movement.
- No sidebar/topnav organism extraction.
- No fixture or locale shape changes.
- No custom JS or manifest changes.
- No behavior change for `/` and `/sample`.

---

## Accessibility

- Existing layout and mobile sheet markup remains unchanged.
- No new ARIA patterns expected.
- `bun run validate:aria` is not required unless implementation unexpectedly changes hooks/manifest.

---

## Validation plan

```bash
bun run templ
go test ./internal/site/... ./internal/views/...
```

If package-level tests are absent or the split touches imports broadly, run:

```bash
go test ./...
```

---

## Acceptance criteria

- [x] `internal/site/feature.go` no longer contains per-page render bodies.
- [x] Runtime route specs live in `internal/site/router.go` or clearly named equivalent.
- [x] Route specs show the layout adapter choice (`views.AppShell` for current routes).
- [x] New page workflow is: fixture fields + locale JSON + view + one route spec.
- [x] `/` and `/sample` still render with active nav and localized copy.
- [x] No custom JS added.

**Block 02 complete.** Validation: `bun run templ`, `go test ./...`.

---

## After Block 02

Proceed to Block 03: update React onboarding docs using the actual runtime route manifest.
