# Active implementation spec — Block 03

**Feature:** React onboarding docs  
**Architecture:** [next-shadcn-architecture.md](./next-shadcn-architecture.md)  
**Progress block:** Block 03 in [next-shadcn-refactor-progress.md](../next-shadcn-refactor-progress.md)

---

## Intent

Update onboarding docs so a React developer who knows Next App Router and shadcn/ui can add a basic page without reading framework internals.

This slice is **documentation only** — no runtime code changes.

---

## Deliverables

| File | Updates |
|------|---------|
| [`docs/for-react-devs.md`](../../docs/for-react-devs.md) | Request flow, file map, add-page cookbook, registry terms, honest dev loop |
| [`README.md`](../../README.md) | Quick start, scripts table, project layout, add-page checklist aligned with `router.go` |

---

## Acceptance criteria

- [x] Request flow documented: `GET /path` → `PageSpec` → route shell → page.
- [x] File map covers Next/shadcn equivalents (`layout.tsx`, `page.tsx`, `components/ui`, `vite.config`, locale JSON).
- [x] Add-page cookbook uses one `PageSpec` in `internal/site/router.go`.
- [x] Route shell vs document shell vs blocks/components/widgets explained.
- [x] Dev loop is honest: no HMR, no Go auto-restart, CSS watch is a second terminal.
- [x] No claim that `routes.yaml` or codegen is available.
- [x] Custom JS guidance limited to `@ui8kit/aria` pattern policy.

**Block 03 complete.** Validation: manual review of updated docs for stale `feature.go` route-registration references.

---

## After Block 03

Proceed to Block 04: freeze registry boundary for layout artifacts (`internal/ui/*` vs runtime adapters).
