# ADR 004 — Templ composition retag

## Decision

Composition layers (`internal/ui/blocks`, `internal/ui/components`, `internal/ui/widgets`,
`internal/views`) must not own raw HTML. Raw tags live in `internal/kit/ui` atoms and in
the document allowlist inside `internal/ui/layout`. Debt is paid via **retag** and
enforced by `scripts/retag-check.sh` (ratchet baseline).

## Why

Architecture already required atom-based composition. Scattered raw tags
duplicate layout/typography, drift from design tokens, and block consistent
ARIA. A named, measurable pass (retag) beats an open-ended “refactor markup”.

`button.Button` now supports `Href` (ui8kit-codegen), so CTA anchors use the
Button atom — no raw `<a>` + `ButtonClasses` exception.

Folder layout matches FastyGoUI registry (`internal/ui/{layout,components,blocks,widgets}`)
so Blank stays the universal reference consumer for a future multi-app kit.

## Consequences

- Normative docs: [retag.md](../retag.md); Cursor rule `.cursor/rules/templ-retag.mdc`.
- CI/local: `retag-check` fails when violation counts exceed `.project/retag-baseline.txt`.
- Allowed escapes: document tags in layout; hidden inputs only via
  `fastygo.config.mjs` → `retag.allow`; other rares via `// retag:allow:`.
- Retag tasks change registry/views only — never codegen under `internal/kit/ui`.
- ui8px deferred until promoting components into FastyGoUI.
