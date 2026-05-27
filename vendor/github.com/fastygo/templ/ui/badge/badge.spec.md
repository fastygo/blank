---
id: ui.badge
layer: primitive
kind: label
package: github.com/fastygo/templ/ui/badge
facade: github.com/fastygo/templ/ui
templ: Badge
api:
  Variant:
    role: appearance
    type: string
    enum: [default, secondary, destructive, outline]
    allow-list-source: ui.badge.BadgeVariants
    default: default
  Size:
    role: density
    type: string
    enum: [default, sm, lg]
    allow-list-source: ui.badge.BadgeVariants
    default: default
  Class:
    role: style-extension
    type: string
slots:
  default:
    required: true
    accepts: text
showcase:
  - id: variant.default
    props: { Variant: default, Size: default }
  - id: variant.secondary
    props: { Variant: secondary, Size: default }
  - id: variant.destructive
    props: { Variant: destructive, Size: default }
  - id: variant.outline
    props: { Variant: outline, Size: default }
  - id: size.sm
    props: { Variant: default, Size: sm }
  - id: size.lg
    props: { Variant: default, Size: lg }
semantics:
  root: div
  role: none
  behavior: static

---
## Summary

Badge shows a short status label.
Badge uses variant and size presets.

## Use Cases

- Show item status in a list
- Show a count next to a title
- Show a category tag on content

## Semantics

- Root element is div
- Badge has no interactive role
- Badge content comes from children

## Example variant.default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Badge(ui.BadgeProps{Variant: "default"}) {
		Default
	}
}
```

## Example variant.secondary

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Badge(ui.BadgeProps{Variant: "secondary"}) {
		Secondary
	}
}
```

## Example variant.destructive

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Badge(ui.BadgeProps{Variant: "destructive"}) {
		Destructive
	}
}
```

## Example variant.outline

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Badge(ui.BadgeProps{Variant: "outline"}) {
		Outline
	}
}
```

## Example size.sm

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Badge(ui.BadgeProps{Size: "sm"}) {
		Small
	}
}
```

## Example size.lg

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Badge(ui.BadgeProps{Size: "lg"}) {
		Large
	}
}
```
