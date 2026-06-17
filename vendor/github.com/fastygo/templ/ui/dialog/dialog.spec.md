---
id: ui.dialog
layer: primitive
kind: interactive
package: github.com/fastygo/templ/ui/dialog
facade: github.com/fastygo/templ/ui
templ: Dialog
api:
  Variant:
    role: appearance
    type: string
    enum: [default, card, sheet, unstyled]
    allow-list-source: ui.dialog.DialogVariants
    default: default
  Size:
    role: density
    type: string
    enum: [sm, default, lg, xl, full]
    default: default
  Open:
    role: state
    type: bool
    default: false
  DataUI8Kit:
    role: behavior-hook
    type: string
    default: ""
showcase:
  - id: variant.default
    props: { Variant: default, Size: default, AriaLabel: "Dialog" }
  - id: variant.sheet
    props: { Variant: sheet, Size: full, AriaLabel: "Panel" }
semantics:
  root: dialog
  behavior: native-dialog
---

## Summary

Dialog renders a native dialog root with optional behavior hook.

## Use Cases

- Render modal content
- Render sheet-like panels with native dialog semantics

## Semantics

- DataUI8Kit is opt-in and only renders when provided
- Use AriaLabel or AriaLabelledBy for the accessible name

## Example variant.default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Dialog(ui.DialogProps{AriaLabel: "Dialog"}) {
		@ui.Text(ui.TextProps{}, "Dialog body.")
	}
}
```

## Example variant.sheet

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Dialog(ui.DialogProps{Variant: "sheet", Size: "full", AriaLabel: "Panel"}) {
		@ui.Text(ui.TextProps{}, "Panel body.")
	}
}
```
