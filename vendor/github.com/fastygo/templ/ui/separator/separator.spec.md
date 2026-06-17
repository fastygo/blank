---
id: ui.separator
layer: primitive
kind: layout
package: github.com/fastygo/templ/ui/separator
facade: github.com/fastygo/templ/ui
templ: Separator
api:
  Variant:
    role: appearance
    type: string
    enum: [default, muted, strong, spaced, unstyled]
    allow-list-source: ui.separator.SeparatorVariants
    default: default
  Orientation:
    role: direction
    type: string
    enum: [horizontal, vertical]
    default: horizontal
  Decorative:
    role: accessibility
    type: bool
    default: false
showcase:
  - id: variant.default
    props: { Variant: default }
  - id: orientation.vertical
    props: { Orientation: vertical, Decorative: true }
semantics:
  root: hr
  behavior: static
---

## Summary

Separator renders a semantic or decorative hr.

## Use Cases

- Separate form sections
- Separate toolbar or navigation groups

## Semantics

- Decorative separators set role presentation and aria-hidden true

## Example variant.default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Separator(ui.SeparatorProps{})
}
```

## Example orientation.vertical

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Separator(ui.SeparatorProps{Orientation: "vertical", Decorative: true})
}
```
