# Retag — composition without raw HTML

**retag** = replace raw HTML tags in composition layers with kit atoms / registry components.

Normative.

**Allowlist registry (SoT):** [`fastygo.config.mjs`](../fastygo.config.mjs) → `retag.allow`  
Tooling reads it via `bun scripts/config-get.mjs retag.allowHiddenFiles`.

## Zones

| Zone | Paths | Raw HTML |
| --- | --- | --- |
| **U — atoms** | `internal/kit/ui/**` | Required. Do not retag / do not hand-edit codegen. |
| **L — document** | `internal/ui/layout/**` | Allowlist only: `html`, `head`, `body`, `meta`, `title`, `link`, `script`, `noscript`, `main`. |
| **C — composition** | `internal/ui/{blocks,components,widgets}/**`, `internal/views/**` | **0 raw**, except registry / marker below. |

## Composition allowlist

1. **Hidden inputs (registry)** — `<input type="hidden">` only in files listed under `retag.allow` in `fastygo.config.mjs`.  
   Comment `// retag:allow: hidden input — …` stays as local documentation; **it is not enough** without a config entry.
2. **Other escapes** — rare non-hidden raw tags need `// retag:allow: <reason>` on the previous line (prefer extending kit instead).

## Tag → atom map

| Raw | Prefer (`internal/kit/ui`) |
| --- | --- |
| `div` | `box`, `stack`, `group`, `grid`, `container` |
| `p` / `span` | `text` / `inline` |
| `h1`–`h6` | `title` (`As`) |
| landmarks | `block` (`Tag`) |
| `a` (prose) | `link` |
| `a` (CTA) | `button` + `Href` |
| controls / `img` | matching kit atoms |
| repeated clusters | `internal/ui/components` (L4) |

## Tooling

```bash
bun run retag                     # ratchet check
./scripts/retag-check.sh report
./scripts/retag-check.sh write-baseline
bun scripts/config-get.mjs retag.allowHiddenFiles
```

Config also documents CSS/templ paths (`css`, `templ`) for Bun scripts — Go server does not load this file.
