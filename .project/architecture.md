# Architecture

```text
cmd/server                 composition root (DI, health, config)
internal/serverapp         Framework wiring (locales, security, site, overlay)
internal/site              HTTP routes (PageSpec → views)
internal/fixtures          embedded UI locale copy (en + ru)
internal/kit/ui/*          UI atoms — codegen (stand-in for fastygo/templ/ui)
internal/kit/utils         Compose / CVA / tag / ARIA helpers
internal/ui/layout         document shell + app chrome (app-owned)
internal/ui/components     registry composites (CLI-ready later)
internal/ui/blocks/*       page sections → future github.com/fastygo/blocks
internal/ui/widgets        interactive shells → future github.com/fastygo/widgets
internal/ui/variants       optional utility maps
internal/ui/utils          thin helpers over kit/utils
internal/views             thin page composition only
internal/devoverlay        optional loopback status overlay
web/static/css             tokens.css (contract) + theme-blank.css (starter pack)
```

## Boundaries

- Framework owns HTTP perimeter, middleware, health, cookie sessions.
- **Kit vs registry (FastyGoUI-aligned):** atoms in `internal/kit/**`; app
  registry in `internal/ui/{layout,components,blocks,widgets,...}` — see
  [internal/ui/README.md](../internal/ui/README.md).
- **Templ zones (retag):** raw HTML only in `internal/kit/ui/**` and the
  document allowlist in `internal/ui/layout/**`. Composition uses atoms —
  see [retag.md](retag.md) and [ADR 004](adr/004-templ-composition-retag.md).
- UI chrome copy lives in `internal/fixtures/locale/` (i18n-ready).
- Design tokens: semantic bridge in `tokens.css`; starter theme pack in
  `theme-blank.css`. ui8px is deferred until promoting into FastyGoUI.
- Closed composition escapes: `<input type="hidden">` only in files listed in
  `fastygo.config.mjs` → `retag.allow`; other rares via `// retag:allow:`.
  CTA navigation uses `button.Button` with `Href`.
- Tooling paths (CSS, retag) live in `fastygo.config.mjs`; Go runtime does not
  import it — Bun scripts and `scripts/retag-check.sh` do.
