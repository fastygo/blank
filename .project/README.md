# Blank project docs

- [Architecture](architecture.md)
- [Retag](retag.md) — composition without raw HTML
- ADRs: [001](adr/001-ssr-htmx.md) · [002](adr/002-fixture-first.md) · [004](adr/004-templ-composition-retag.md)
- UI registry: [`internal/ui/README.md`](../internal/ui/README.md) · kit: [`internal/kit/README.md`](../internal/kit/README.md)
- Tooling SoT: [`fastygo.config.mjs`](../fastygo.config.mjs) (`registry`, `retag`, `css`)
- Cursor: `.cursor/rules/ui-registry.mdc` · `templ-retag.mdc`

## Local commands

| Script | Purpose |
| --- | --- |
| `bun run dev` / `bun run g` | CSS + JS + overlay build, then `scripts/dev.sh` (`go run` on `:3000`) |
| `bun run retag` | Zoned raw-HTML check (allowlist from `fastygo.config.mjs`) |
| `bun run test` | Retag ratchet + `go test ./...` |
| `bun run build:css` | Tailwind via paths in `fastygo.config.mjs` |
| `bun run config` | Dump tooling config JSON |
| `bun run verify` | templ + CSS + JS + retag + tests |

## Stack

- Framework Feature modules + site router
- templ registry: `internal/ui/{layout,components,blocks,widgets}` + thin `views`
- Atoms in `internal/kit/ui` (codegen), helpers in `internal/kit/utils`
- Locale UI copy: `internal/fixtures/locale` (en + ru)
- Design tokens: `web/static/css/tokens.css` + `theme-blank.css`
- Sheet composite still from `github.com/fastygo/templ/components` until kit ships Sheet
