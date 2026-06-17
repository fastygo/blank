---
id: ui.link
layer: primitive
kind: navigation
package: github.com/fastygo/templ/ui/link
facade: github.com/fastygo/templ/ui
templ: Link
api:
  Variant:
    role: appearance
    type: string
    enum: [default, muted, nav, block, external, unstyled]
    allow-list-source: ui.link.LinkVariants
    default: default
  Size:
    role: density
    type: string
    enum: [sm, default, lg]
    allow-list-source: ui.link.LinkVariants
    default: default
  Href:
    role: destination
    type: string
  External:
    role: security
    type: bool
    default: false
  AriaCurrent:
    role: current-state
    type: string
showcase:
  - id: variant.default
    props: { Href: "/docs" }
  - id: variant.external
    props: { Href: "https://example.com", External: true }
semantics:
  root: a
  behavior: interactive
---

## Summary

Link renders a semantic anchor for navigation and prose.
Use Button for actions.

## Use Cases

- Render navigation links
- Render prose links and external resources

## Semantics

- External links default target blank and rel noopener noreferrer
- Disabled links set aria-disabled and tabindex -1

## Example variant.default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Link(ui.LinkProps{Href: "/docs"}) { Docs }
}
```

## Example variant.external

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Link(ui.LinkProps{Href: "https://example.com", External: true}) { External }
}
```
