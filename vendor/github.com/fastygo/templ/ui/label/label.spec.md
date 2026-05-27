---
id: ui.label
layer: primitive
kind: input
package: github.com/fastygo/templ/ui/label
facade: github.com/fastygo/templ/ui
templ: Label
api:
  HTMLFor:
    role: control-id
    type: string
  Class:
    role: style-extension
    type: string
  Attrs:
    role: html-attrs
    type: templ.Attributes
slots:
  default:
    required: true
    accepts: text
showcase:
  - id: layout.with-control
    props: { HTMLFor: demo-email }
semantics:
  root: label
  role: none
  behavior: static

---
## Summary

Label links visible text to one form control.
Label sets for attribute when HTMLFor is set.

## Use Cases

- Name an email input field
- Name a checkbox in a form item

## Semantics

- Root element is label
- for attribute references control id from HTMLFor

## Example layout.with-control

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Label(ui.LabelProps{HTMLFor: "demo-email"}) {
		Email
	}
}
```
