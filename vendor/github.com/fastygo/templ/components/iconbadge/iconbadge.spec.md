---
id: components.iconbadge
layer: composite
kind: media
package: github.com/fastygo/templ/components/iconbadge
facade: github.com/fastygo/templ/components
templ: IconBadge
api:
  Variant:
    role: appearance
    type: string
    enum: [default, accent, secondary, outline, destructive, unstyled]
    allow-list-source: components.iconbadge.IconBadgeVariants
    default: default
  Size:
    role: density
    type: string
    enum: [xs, sm, default, lg]
    default: default
  Name:
    role: icon-name
    type: string
  BaseClass:
    role: icon-pack-base-class
    type: string
  Prefix:
    role: icon-name-prefix
    type: string
  Text:
    role: text-fallback
    type: string
showcase:
  - id: variant.default
    props: { Name: home }
  - id: icon.pack
    props: { BaseClass: icon, Prefix: "icon-", Name: home }
semantics:
  root: span
  behavior: static
---

## Summary

IconBadge renders a small icon surface for nav rows, cards, and status blocks.

## Use Cases

- Render navigation icons with a consistent surface
- Render workflow and feature card pictograms

## Semantics

- IconBadge uses ui.Icon when BaseClass or svg type is provided
- IconBadge falls back to Text or the first rune of Name

## Example variant.default

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.IconBadge(cmp.IconBadgeProps{Name: "home"})
}
```

## Example icon.pack

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.IconBadge(cmp.IconBadgeProps{BaseClass: "icon", Prefix: "icon-", Name: "home"})
}
```
