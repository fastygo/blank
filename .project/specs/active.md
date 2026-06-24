# Active implementation spec — Block 01

**Feature:** Named route shells for Next/shadcn onboarding  
**Architecture:** [next-shadcn-architecture.md](./next-shadcn-architecture.md)  
**Progress block:** Block 01 in [next-shadcn-refactor-progress.md](../next-shadcn-refactor-progress.md)

---

## Intent

Add recognizable route-level layout names so React developers see `AppShell` / `MarketingShell` / `DocsShell` instead of a single ambiguous `SiteShell`, without changing visual behavior or introducing sidebar work.

## Assumptions

- Current topnav behavior in `layout.Shell` is the correct default for `AppShell`.
- `MarketingShell` and `DocsShell` may delegate to the same chrome as `AppShell` in this slice if that minimizes diff.
- Handlers may keep calling `SiteShell` during migration; alias must remain valid.

## Affected surfaces

| File | Change |
|------|--------|
| [`internal/views/layout.templ`](../../internal/views/layout.templ) | Add `AppShell`, `MarketingShell`, `DocsShell`; refactor `SiteShell` to alias `AppShell` |
| [`internal/views/wireframe_render_test.go`](../../internal/views/wireframe_render_test.go) | Optionally add `TestAppShell_homeRenders` or rename existing test; keep coverage |
| [`docs/for-react-devs.md`](../../docs/for-react-devs.md) | Short section: document shell vs route shell; `SiteShell` → `AppShell` |
| [`README.md`](../../README.md) | Optional one-line pointer to architecture spec |

**Out of scope for Block 01:**

- `internal/site/feature.go` handler renames (optional; not required if alias works)
- Sidebar presets, layout parts, `routes.go`
- Fixture/locale changes
- Custom JS or manifest changes

## UI structure

```templ
templ AppShell(d LayoutData, body templ.Component) {
  @layout.Shell(/* same ShellProps as today */) { @body }
}

templ MarketingShell(d LayoutData, body templ.Component) {
  @AppShell(d, body)  // or minimal future divergence
}

templ DocsShell(d LayoutData, body templ.Component) {
  @AppShell(d, body)  // placeholder until docs routes
}

templ SiteShell(d LayoutData, body templ.Component) {
  @AppShell(d, body)
}
```

Extract shared `ShellProps` construction to a private helper in the same file if it avoids duplication (e.g. `appShellProps(d LayoutData) layout.ShellProps`).

## Accessibility

- No ARIA or manifest changes expected.
- Existing mobile sheet, theme toggle, and language switch markup unchanged.

## Validation plan

```bash
bun run templ
go test ./internal/views/...
```

If only docs + templ with no class changes: ui8px lint optional. Full `bun run verify` if unsure.

## Acceptance criteria

- [x] `views.AppShell`, `MarketingShell`, `DocsShell` exist in `layout.templ`
- [x] `views.SiteShell` delegates to `AppShell`
- [x] `/` and `/sample` render identically to pre-refactor (manual or existing render test)
- [x] Docs explain two layout layers and alias policy
- [x] No custom JS added

## Block 01 completion notes

- **Completed:** Named route shells in `layout.templ`; `TestAppShell_homeRenders` + `TestSiteShell_aliasRenders`; onboarding docs updated.
- **Validation:** `bun run templ`, `go test ./internal/views/...` — pass.

## After Block 01

Proceed to Block 02: split `internal/site/feature.go` into route manifest + render helper.
