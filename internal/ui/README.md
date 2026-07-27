# In-app UI registry (`internal/ui`)

**Reference layout** for staging reusable UI before promotion to shared modules.
Use this tree in Blank and product apps the same way — copy the structure, fill
packages incrementally, extract when stable.

## Positioning

| Layer | Module | Role |
|-------|--------|------|
| Atoms / helpers | `github.com/fastygo/blank/internal/kit/{ui,utils}` | Cn, CVA, tags, ARIA, Button, Stack, … |
| Kit composites (interim) | `github.com/fastygo/templ/components` | Sheet (until kit ships it) |
| App registry | **`internal/ui/*`** (this tree) | Chrome, blocks, widgets, app components |

**Not Go dependencies during staging:** `github.com/fastygo/blocks`, `github.com/fastygo/widgets`.
Develop here first; `require` shared modules only after extraction.

**Not part of this registry:** `internal/views/` (route pages), `internal/site/` (runtime route manifest).

## Tree

```
internal/ui/
  layout/       # registry:layout — shells, data.go, build.go, shell_head, header_trailing
  components/   # registry:components — icon, toggles, navigation, appsidebar, …
  blocks/       # registry:blocks — full scaffolds (staging → fastygo/blocks)
    marketing/  # hero/ (live); add organisms as needed
  widgets/      # registry:widgets — staging stub
  variants/     # registry:variants — staging stub
  utils/        # registry:utils — thin helpers on kit/utils
```

## Where does this UI go?

| You are building… | Put it in… |
|-------------------|------------|
| Document frame (`html`, `head`, `body`) | `layout/root_layout.templ` |
| Topnav chrome (header, footer, mobile sheet) | `layout/topnav_layout.templ` |
| Dashboard chrome (topnav + aside) | `layout/dashboard_layout.templ` |
| Small reusable control (icon, toggle) | `components/<area>/` |
| Aside / sidebar content | `components/appsidebar/` |
| Full scaffold with default copy | `blocks/<domain>/<organism>/` |
| Shell that fetches or orchestrates | `widgets/` |
| Route page | `internal/views/<page>.templ` |
| HTTP route + page choice | `internal/site/router.go` (`PageSpec`) |

## Composition rules

- **No raw HTML** in composition — use `internal/kit/ui` (+ Sheet from templ/components for now).
- Document shell only: `<!DOCTYPE>`, `<html>`, `<head>`, `<body>`, … in `layout/root_layout.templ` or `layout/shell_head.templ`.
- Policy: [`.project/retag.md`](../../.project/retag.md), Cursor `templ-retag.mdc` / `ui-registry.mdc`.
- Covered interaction: `@ui8kit/aria` + Sheet with `Behavior: "ui8kit"`.

## Related docs

- Architecture: [`.project/architecture.md`](../../.project/architecture.md)
- Kit: [`../kit/README.md`](../kit/README.md)
- App README: [`../../README.md`](../../README.md)
