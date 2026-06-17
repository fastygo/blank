---
id: ui.icon
layer: primitive
kind: media
package: github.com/fastygo/templ/ui/icon
facade: github.com/fastygo/templ/ui
templ: Icon
api:
  Type:
    role: render-mode
    type: string
    enum: [class, svg, text]
    allow-list-source: ui.icon.IconVariants
    default: class
  Size:
    role: density
    type: string
    enum: [xs, sm, default, lg, xl]
    allow-list-source: ui.icon.IconVariants
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
  Href:
    role: svg-use-reference
    type: string
  Title:
    role: accessible-name
    type: string
  Decorative:
    role: accessibility
    type: bool
    default: false
showcase:
  - id: type.class
    props: { Type: class, BaseClass: icon, Prefix: "icon-", Name: home, Size: sm }
  - id: type.svg
    props: { Type: svg, Href: "#check", Title: "Done" }
semantics:
  root: span | svg
  behavior: static
---

## Summary

Icon renders a class-based span, svg use reference, or text fallback.
Decorative icons are hidden from assistive technology by default.

## Use Cases

- Render navigation and toolbar icons
- Render inline SVG sprite references

## Semantics

- Empty Title and AriaLabel make the icon decorative
- Accessible icons receive role img

## Example type.class

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Icon(ui.IconProps{BaseClass: "icon", Prefix: "icon-", Name: "home", Size: "sm"})
}
```

## Example type.svg

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Icon(ui.IconProps{Type: "svg", Href: "#check", Title: "Done"})
}
```
