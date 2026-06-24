---
id: components.alert
layer: composite
kind: feedback
package: github.com/fastygo/templ/components/alert
facade: github.com/fastygo/templ/components
templ: Alert
api:
  Variant:
    role: appearance
    type: string
    enum: [default, destructive, success, warning]
    allow-list-source: utils.recipes.AlertVariants
    default: default
  Class:
    role: style-extension
    type: string
  Role:
    role: live-region-role
    type: string
    default: status
  AriaLive:
    role: live-region-politeness
    type: string
    enum: [off, polite, assertive]
    default: polite
  Attrs:
    role: html-attrs
    type: templ.Attributes
slots:
  default:
    required: true
    accepts: text
showcase:
  - id: variant.default
    props: { Variant: default }
  - id: variant.destructive
    props: { Variant: destructive }
  - id: variant.success
    props: { Variant: success }
  - id: variant.warning
    props: { Variant: warning }
semantics:
  root: section
  role: status
  aria-live: polite
  behavior: static
---

## Summary

Alert shows a short status message block.
Alert uses role status and aria-live polite.

## Use Cases

- Show form validation summary
- Show destructive action warning text
- Show success confirmation message

## Semantics

- Root element is section
- Role defaults to status
- AriaLive defaults to polite
- Children carry the visible message body

## Example variant.default

```templ
import "github.com/fastygo/templ/ui"
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Alert(cmp.AlertProps{Variant: "default"}) {
		@ui.Text(ui.TextProps{Class: "text-sm"}, "Static alert with role status.")
	}
}
```

## Example variant.destructive

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Alert(cmp.AlertProps{Variant: "destructive"}) {
		Something went wrong.
	}
}
```

## Example variant.success

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Alert(cmp.AlertProps{Variant: "success"}) {
		Changes saved.
	}
}
```

## Example variant.warning

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Alert(cmp.AlertProps{Variant: "warning"}) {
		Review settings before you continue.
	}
}
```
