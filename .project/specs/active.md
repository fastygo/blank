# Active implementation spec — Block 04

**Feature:** Registry boundary for layout artifacts  
**Architecture:** [next-shadcn-architecture.md](./next-shadcn-architecture.md)  
**Progress block:** Block 04 in [next-shadcn-refactor-progress.md](../next-shadcn-refactor-progress.md)

---

## Intent

Freeze where reusable layout-related UI artifacts live before sidebar/app shell variants are implemented. Separate runtime router wiring from shadcn-like copy-paste registry accumulation.

This slice is **documentation only** — no runtime UI changes, no sidebar implementation.

---

## Deliverables

| File | Updates |
|------|---------|
| [`next-shadcn-architecture.md`](./next-shadcn-architecture.md) | Registry boundary section, folder convention, route adapter policy |
| [`internal/ui/README.md`](../../internal/ui/README.md) | Where-does-this-go table, runtime vs registry |
| [`internal/ui/layout/README.md`](../../internal/ui/layout/README.md) | Permanent app chrome vs layout organisms |
| [`internal/ui/blocks/README.md`](../../internal/ui/blocks/README.md) | Domain folder convention, layout organisms |
| [`internal/ui/widgets/README.md`](../../internal/ui/widgets/README.md) | Widgets vs blocks for layout + behavior |
| [`blank-ui-structure.mdc`](../../.cursor/rules/blank-ui-structure.mdc) | Updated site package layout, add-page flow |

---

## Frozen decisions

- **`internal/ui/layout/`** — document/chrome infrastructure only; stays in app.
- **`internal/ui/blocks/<domain>/<organism>/`** — layout organisms (e.g. `dashboard/sidebar_app`, `docs/toc_shell`, `marketing/topnav_shell`).
- **No** `internal/ui/blocks/layout/`, `recipes`, `elements`, or `ui/`.
- **`views.*Shell`** — thin runtime adapters; reusable markup moves to registry in Block 05+.
- **`internal/site/router.go`** — runtime wiring only; not a registry.

---

## Acceptance criteria

- [x] Future layout work has an explicit registry home under `blocks/<domain>/<organism>/`.
- [x] Team can distinguish runtime router wiring from copy-paste UI artifacts.
- [x] `internal/ui/layout` vs `blocks` vs `widgets` vs `components` responsibilities documented.
- [x] `views.AppShell` policy documented as temporary thin adapter.
- [x] No sidebar UI implemented.
- [x] No forbidden folder names introduced.
- [x] No custom JS guidance beyond `@ui8kit/aria`/Sheet policy.

**Block 04 complete.** Validation: doc review for stale `feature.go` route-registration and forbidden folders.

---

## After Block 04

Proceed to Block 05: extract current topnav shell from `views` into a registry artifact; keep `views.AppShell` as thin adapter.
