# `sidebar_app` — dashboard sidebar shell block

**Registry artifact** (`sidebarapp.SidebarApp`): document chrome from `layout.Shell` plus a desktop sidebar and main content column. Mobile navigation uses the existing `components/navigation` sheet from the shell header.

This block is **not** selected by current routes. `views.AppShell` still delegates to `app_shell` until Block 09 wires explicit layout choice in the router.

## Block contract (Block 07)

```text
sidebarapp.SidebarApp(props, body)
  -> layout.Shell(props.Shell)
       -> navigation.MobileSheet + layout.Header (from shell)
       -> main slot:
            -> aside (desktop, md+) + vertical navigation.Nav
            -> content column -> body slot
```

| File | Role |
|------|------|
| `props.go` | `Props{Shell, Sidebar}` |
| `helpers.go` | Sidebar fallbacks, aside classes |
| `defaults.go` | `DefaultProps()` for tests/showcase |
| `sidebar_app.templ` | `SidebarApp` markup |

## Region vocabulary

Use these scopes when describing layout geometry. They are **documentation axes**, not runtime enum values.

| Scope | Meaning |
|-------|---------|
| `shell_full` | Spans the full document width (above or below the main grid) |
| `main_column` | Inside the primary content column (sidebar inset / main area) |
| `content_row` | Horizontal row that holds left aside + page + right aside |
| `viewport` | Fixed to the viewport edge (full-height sidebar column) |

| Region | Typical content |
|--------|-----------------|
| Shell header | Brand, theme toggle, mobile menu trigger |
| Shell footer | App footer copy |
| Aside left / right | Vertical nav, toc, filters |
| Main header band | Page title, breadcrumbs, actions |
| Main body | Route page content |

## Showcase wireframes (Block 08)

Three sidebar geometry examples are **showcase IDs** — names for explaining wireframes, not mandatory runtime presets. Future variants should be **forked or composed** as registry blocks under `blocks/<domain>/<organism>/`, not added as a global layout engine or `APP_LAYOUT` enum.

### `sidebars_full`

| Region | Scope |
|--------|-------|
| Left aside | `viewport` |
| Right aside | `content_row` |
| Header band | `main_column` |

Left nav is viewport-height; right panel sits in the content row beside page markup; page header lives in the main column above the row.

### `sidebars_main`

| Region | Scope |
|--------|-------|
| Header | `shell_full` |
| Footer | `shell_full` |
| Left aside | `content_row` |
| Right aside | `content_row` |

Full-width shell header and footer; both sidebars flank the page inside the content row.

### `sidebars_header`

| Region | Scope |
|--------|-------|
| Shell header | `shell_full` |
| Page/header band | Below shell header (main column entry) |
| Left aside | `content_row` |
| Right aside | `content_row` |

App chrome header spans full width; a secondary page header band sits below it; sidebars live in the content row under that band.

## Relation to current `sidebar_app`

The implemented block is closest to **`layout.sidebar_app`** / a single left aside in the main slot (Block 07). The three showcase IDs above describe **additional geometries** to fork from this artifact or from shared parts (`navigation`, future `AsideRegion` / `ShellBand` composites).

Do **not** wire showcase IDs into `internal/site/router.go` until Block 09 chooses layouts explicitly per route.

## How to add a new geometry

1. Copy or compose from `sidebar_app` (or extract shared parts into `components/` when reused twice).
2. Document scope mapping in the organism README (like this file).
3. Add render tests for landmarks and ARIA contracts, not utility class strings.
4. Register the block in route `PageSpec.Layout` only when the product needs that shell (Block 09+).

## Related

- [`.project/sidebar-like.md`](../../../../.project/sidebar-like.md) — shadcn-style compound layout model
- [`../../README.md`](../../README.md) — blocks registry index
- [`../../../layout/README.md`](../../../layout/README.md) — document/chrome infrastructure
