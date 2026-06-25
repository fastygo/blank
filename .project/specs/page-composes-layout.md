# Page composes layout (shadcn parity)

**Status:** delivered (2026-06-25); extended by registry co-location slice — see [active.md](./active.md)
**Supersedes:** [archive/block-11-final.md](./archive/block-11-final.md)
**Driver:** [`.cursor/rules/blank-spec-driver.mdc`](../../.cursor/rules/blank-spec-driver.mdc), brainstorm [`brainstorm/layout-structure.md`](../brainstorm/layout-structure.md)

> **Superseded paths (registry co-location, 2026-06-25):** `views/models.go`, `views/layout_helpers.go`, `views/partials/`, `sample_stub.templ` — replaced by `layout/data.go`, `layout/build.go`, co-located `views/<page>.go`, and `views/sample.templ`.

---

## Intent

Collapse Blank's route-shell stack so that opening `internal/views/<route>.templ` shows the **entire layout composition tree** in a single file — the same cognitive experience as opening `app/dashboard/page.tsx` in a shadcn project. Eliminate empty wrapper layers that exist only to forward props.

Today, understanding `/sample` requires reading 5 files in chain:

```text
router.go  →  views.SidebarAppShell  →  blocks/sidebar_app  →  layout.Shell  →  components/navigation
```

After this slice, opening `views/sample.templ` directly shows:

```text
layout.SidebarShell + appsidebar.AppSidebar + content
```

No registry indirection, no `PageSpec.Layout` adapter chain.

---

## Affected surface

- `internal/site/router.go` — `PageSpec.Layout` removed; `Body` returns a fully composed document.
- `internal/site/render.go` — single `web.Render(ctx, w, page.Body(data, fix))` call; no `page.Layout(...)` wrapping.
- `internal/views/layout.templ` — **deleted** (the 4 route adapters `AppShell`, `SidebarAppShell`, `MarketingShell`, `DocsShell` were renames).
- `internal/views/home.templ`, `sample_stub.templ` — each page composes its layout shell directly (`@layout.Shell { ... }` or `@layout.SidebarShell(..., aside) { ... }`).
- `internal/views/models.go` — page-data structs gain a `Shell layout.ShellProps` field (and `Sidebar appsidebar.Props` where relevant).
- `internal/views/layout_helpers.go` — `shellProps()` becomes public `ShellPropsFor()`; add `SidebarPropsFor()` builder.
- `internal/ui/layout/sidebar_shell.templ` — **new** first-class sidebar shell; analogous to shadcn `SidebarProvider + SidebarInset`.
- `internal/ui/components/appsidebar/` — **new** local sidebar component; analogous to shadcn `components/app-sidebar.tsx`.
- `internal/ui/blocks/dashboard/sidebar_app/`, `internal/ui/blocks/dashboard/app_shell/` — **deleted** (absorbed into `layout/` and `components/appsidebar/`).
- Tests: existing render tests in deleted blocks are removed; `views/wireframe_render_test.go` exercises pages directly through the new composition.
- Rules / docs: `blank-ui-structure.mdc`, `internal/ui/{README.md, layout/README.md, blocks/README.md}`, `docs/for-react-devs.md`, `.project/specs/next-shadcn-architecture.md`.

---

## Design

### Two named shells in `internal/ui/layout/`

| Shell | Composition | shadcn analogue |
|---|---|---|
| `layout.Shell` (kept) | `html + head + body + Header + main{children} + Footer` | `SidebarProvider` without sidebar |
| `layout.SidebarShell` (new) | `Shell + (aside slot, main{children})` | `SidebarProvider + AppSidebar + SidebarInset` |

`SidebarShell` signature uses two parameters for explicitness:

```go
templ SidebarShell(shell ShellProps, sidebar templ.Component) { ... }
```

The `sidebar` slot is `nil`-safe so callers may omit it (rare; but matches Next "no sidebar variant").

### Local sidebar component

`internal/ui/components/appsidebar/`:

```go
type Props struct {
    Title     string
    AriaLabel string
    Items     []navigation.Item
    Active    string
    Class     string
}
```

`AppSidebar(props)` renders the `<aside>` content (title + vertical `navigation.Nav`). It is **the** copy-paste-friendly artifact: a project can swap or extend it without touching `layout.SidebarShell`.

### Page = composition

`internal/views/home.templ`:

```templ
templ HomePage(d HomePageData) {
    @layout.Shell(d.Shell) {
        @ui.Box(...) { ...content... }
    }
}
```

`internal/views/sample_stub.templ`:

```templ
templ SamplePage(d SamplePageData) {
    @layout.SidebarShell(d.Shell, appsidebar.AppSidebar(d.Sidebar)) {
        @ui.Box(...) { ...content... }
    }
}
```

One file, one tree. This is the shadcn ideal in templ.

### Router stays thin

```go
type PageSpec struct {
    Method  string
    Pattern string
    Active  string
    Title   TitleResolver
    Body    PageRenderer
    Nav     NavResolver
}

type PageRenderer func(views.LayoutData, fixtures.Locale) templ.Component
```

`Body` returns a **full HTML document** (because the page composes the shell). `handlePage` becomes one `web.Render` call.

---

## UI and component choices

- Primitives from `github.com/fastygo/templ/ui` only — no raw HTML except in the existing document-shell exceptions.
- `navigation.Nav` / `navigation.MobileSheet` / `MobileSheetTrigger` unchanged.
- `appsidebar.AppSidebar` composes `ui.Box` (`Tag: "aside"`), `ui.Stack`, `ui.Text`, and `navigation.Nav`.
- `layout.SidebarShell` composes `layout.Shell` (its document chrome is unchanged), wrapping `flex md:flex-row` body + sidebar slot + main column.

---

## Accessibility and interaction

- Sidebar `<aside>` keeps `aria-label` from `Props.AriaLabel` (falls back to shell's `MainNavigation`).
- Active item in vertical nav keeps `aria-current="page"` (already in `navigation.Nav`).
- Mobile sheet behavior is unchanged — same `data-ui8kit` hooks served by `layout.Shell`. No manifest changes required.

---

## Fixtures and locale

- No new locale keys. `appsidebar.Props` is built from existing `views.LayoutData` (nav items, active route, navigation aria labels) via a new `views.SidebarPropsFor(d, title)` helper.
- Per-route sidebar title falls back to brand when omitted.

---

## Validation

```bash
bun run templ
go build ./...
go test ./internal/...
bun run lint:ui8px
bun run validate:aria
bun run verify
```

Render-test invariants to keep green:

- `TestAppShell_homeRenders` (renamed to follow new path: `TestHomePage_renders`)
  - doctype, `<title>Home · FastyGo</title>`, no `<aside>`, sheet markup, theme toggle, footer, hero copy, language switch.
- `TestSidebarAppShell_sampleRenders` → `TestSamplePage_renders`
  - doctype, `<title>Sample · FastyGo</title>`, `<aside>`, `aria-label="Main navigation"`, `aria-current="page"`, sheet markup, page body.

---

## Acceptance criteria

1. Opening `internal/views/sample_stub.templ` shows the full layout tree (`SidebarShell + AppSidebar + content`) — no jump through `views/layout.templ` or `blocks/`.
2. `internal/site/router.go` has no `Layout` field on `PageSpec`. Adding a new page is: one fixture key, one `*PageData`, one `*Page` templ, one `PageSpec` entry.
3. `internal/ui/blocks/dashboard/sidebar_app/` and `internal/ui/blocks/dashboard/app_shell/` are removed; `internal/ui/blocks/dashboard/` retains `doc.go` as a domain stub for future blocks.
4. `bun run verify` passes.

---

## Non-goals

- Renaming `layout.Shell` to `layout.AppShell` (cosmetic; not worth the import churn — README documents the parity mapping).
- Introducing a layout engine, `APP_LAYOUT` config, or showcase enum.
- Splitting `site` / `app` / `docs` features.
- Adding shadcn-style icon-collapsing sidebar with cookie state.
- Touching `internal/devoverlay/`.

---

## Trade-offs

- Each page repeats one line `@layout.Shell(d.Shell) { ... }` (or `SidebarShell`). This is the visible-composition tax — and the whole point.
- `*PageData` gains a `Shell` field. Pages now know their chrome — explicit contract.
- `PageSpec.Layout` declarative perk goes away — but it was misleading (layout was not visible from the page anyway).

---

## Rollout

Single PR (no intermediate state where router and views disagree). Steps in order:

1. Add `internal/ui/layout/sidebar_shell.templ` + helpers.
2. Add `internal/ui/components/appsidebar/`.
3. Update `views/models.go`, `layout_helpers.go`, `home.templ`, `sample_stub.templ`.
4. Drop `views/layout.templ`; update `site/router.go` + `render.go`.
5. Delete `internal/ui/blocks/dashboard/{sidebar_app,app_shell}/`.
6. Update render tests in `views/`.
7. Update rules + READMEs + docs.
8. `bun run templ` + `bun run verify`.
