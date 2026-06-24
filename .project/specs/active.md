# Active implementation spec — Block 05

**Feature:** Extract topnav shell into registry artifact  
**Architecture:** [next-shadcn-architecture.md](./next-shadcn-architecture.md)  
**Progress block:** Block 05 in [next-shadcn-refactor-progress.md](../next-shadcn-refactor-progress.md)

---

## Intent

Move the current topnav app shell composition from `views/layout.templ` into a reusable `internal/ui` registry artifact while keeping `views.AppShell(data, body)` as a thin runtime adapter.

No new sidebar, docs toc, marketing divergence, or visual changes.

---

## Deliverables

| Path | Role |
|------|------|
| [`internal/ui/blocks/dashboard/app_shell/`](../../internal/ui/blocks/dashboard/app_shell/) | `appshell.AppShell` — props-only block wrapping `layout.Shell` |
| [`internal/views/layout.templ`](../../internal/views/layout.templ) | `views.AppShell` maps `LayoutData` → `appshell.Props` and delegates |

---

## Acceptance criteria

- [x] `views.AppShell(data, body)` still works.
- [x] `MarketingShell`, `DocsShell`, `SiteShell` still delegate to `AppShell`.
- [x] Current topnav shell discoverable at `internal/ui/blocks/dashboard/app_shell`.
- [x] Artifact uses `layout.ShellProps`, not `views.LayoutData` (no import cycle).
- [x] `TestAppShell_homeRenders`, `TestSiteShell_aliasRenders`, and artifact render test pass.
- [x] No custom JS; no new layout variants.

**Block 05 complete.** Validation: `bun run templ`, `bun run lint:ui8px`, `go test ./...`.

---

## Block 06 — Reusable mobile sheet/nav UI

**Goal:** Extract mobile navigation sheet and nav rendering into reusable props-only components for topnav, sidebar, and docs blocks.

**Delivered:**

- [x] `internal/ui/components/navigation/` — `Nav`, `MobileSheet`, `MobileSheetTrigger` via `templ/components` Sheet (`Behavior: "ui8kit"`).
- [x] `layout.Shell` / `layout.Header` consume navigation components; `layout.NavItem` aliases `navigation.Item`.
- [x] Stable ui8kit IDs preserved (`ui8kit-mobile-sheet-panel`, `-trigger`, `-title`).
- [x] Focused render tests in `navigation_render_test.go`; shell/view tests still pass.
- [x] No manifest or JS bundle changes (`dialog` pattern sufficient).

**Block 06 complete.** Validation: `bun run templ`, `bun run lint:ui8px`, `bun run validate:aria`, `go test ./...`.

---

## Block 07 — Sidebar app layout organism

**Goal:** Add a reusable sidebar/app shell block under the registry without changing current route layout selection.

**Delivered:**

- [x] `internal/ui/blocks/dashboard/sidebar_app/` — `sidebarapp.SidebarApp` with `Props{Shell, Sidebar}`.
- [x] Desktop aside (`ui.Box` tag `aside`) + vertical `navigation.Nav`; mobile sheet from `layout.Shell`.
- [x] `DefaultProps()` for tests/showcase; no imports from `views` or `site`.
- [x] Focused render tests; `app_shell` and route adapters unchanged.

**Block 07 complete.** Validation: `bun run templ`, `bun run lint:ui8px`, `bun run validate:aria`, `go test ./...`.

---

## After Block 07

Proceed to Block 08: sidebar wireframe showcase artifacts (`sidebars_full`, `sidebars_main`, `sidebars_header`).
