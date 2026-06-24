# `registry:widgets`

**UI + behavior** — loading state, user actions, API calls, timers, glue to app services.
Composes `components`, `blocks`, and `github.com/fastygo/templ/*` internally.

Widgets are registry artifacts with **orchestration**. They are not HTTP handlers and do not own routes.

## When to use widgets vs blocks

| Situation | Use |
|-----------|-----|
| Props-only layout or section wireframe | `blocks/<domain>/<organism>/` |
| Renders resolved props only | `components/` or `blocks/` |
| Fetches data, manages loading/error, coordinates side effects | `widgets/` |

Example future widgets: live notification shell, authenticated nav that calls an API, dashboard widget that polls metrics.

Layout **geometry** without fetch → **`blocks/`**. Layout **plus behavior** → **`widgets/`**.

## Rules

- Handlers / features pass **resolved** view models; widgets may call services but **must not** own routing.
- Prefer composing existing **blocks** before duplicating section or shell markup.
- Same styling contract as the rest of the registry: Tailwind + tokens, ui8px, no raw tags.
- Covered interaction: `@ui8kit/aria` — no custom JS for patterns already in the manifest.

## Extraction

When stable → **`github.com/fastygo/widgets/<name>`**. App keeps routes in `internal/site/router.go` and wires dependencies.

## Related

- Registry index: [`../README.md`](../README.md)
- Blocks (layout organisms): [`../blocks/README.md`](../blocks/README.md)
