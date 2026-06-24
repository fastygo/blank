---
id: components.sheet
layer: composite
kind: overlay
package: github.com/fastygo/templ/components/sheet
facade: github.com/fastygo/templ/components
templ: Sheet
parts:
  - templ: Sheet
    props: [ID, Variant, Side, Size, Class, Open, AriaLabel, AriaLabelledBy, AriaDescribedBy, Behavior, Attrs]
  - templ: SheetTrigger
    props: [ID, For, Class, Variant, Size, Open, Behavior, AriaLabel, Attrs]
  - templ: SheetOverlay
    props: [For, Class, Open, Behavior, Attrs]
  - templ: SheetContent
    props: [ID, Class, Attrs]
  - templ: SheetHeader
    props: [Class, Attrs]
  - templ: SheetTitle
    props: [ID, Class, Attrs]
    value: string
  - templ: SheetDescription
    props: [ID, Class, Attrs]
    value: string
  - templ: SheetClose
    props: [For, Class, Variant, Size, Behavior, AriaLabel, Attrs]
api:
  Variant:
    role: appearance
    type: string
    enum: [default, card, ghost, unstyled]
    allow-list-source: components.sheet.SheetVariants
    default: default
  Side:
    role: placement
    type: string
    enum: [left, right, top, bottom]
    default: right
  Size:
    role: density
    type: string
    enum: [sm, default, md, lg, xl, full]
    default: default
  Behavior:
    role: behavior-hook
    type: string
    enum: ["", ui8kit]
    default: ""
showcase:
  - id: side.left
    props: { ID: demo-sheet, Side: left, Size: sm, AriaLabel: "Navigation" }
  - id: behavior.ui8kit
    props: { ID: demo-sheet, Behavior: ui8kit }
semantics:
  root: div
  role: dialog
  behavior: optional
---

## Summary

Sheet composes a side panel with trigger, overlay, content, and close parts.
Behavior hooks are opt-in through Behavior.

## Use Cases

- Render mobile navigation panels
- Render settings and inspector side panels

## Semantics

- Sheet root uses role dialog and aria-modal true
- Trigger wires aria-haspopup, aria-controls, and aria-expanded
- Behavior ui8kit adds data-ui8kit dialog hooks

## Example side.left

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.SheetTrigger(cmp.SheetTriggerProps{For: "demo-sheet", AriaLabel: "Open navigation"}) { Menu }
	@cmp.Sheet(cmp.SheetProps{ID: "demo-sheet", Side: "left", AriaLabel: "Navigation"}) {
		@cmp.SheetContent(cmp.SheetContentProps{}) {
			Navigation
		}
	}
}
```

## Example behavior.ui8kit

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.SheetTrigger(cmp.SheetTriggerProps{For: "demo-sheet", Behavior: "ui8kit"}) { Open }
	@cmp.Sheet(cmp.SheetProps{ID: "demo-sheet", Behavior: "ui8kit", AriaLabel: "Panel"}) {
		@cmp.SheetOverlay(cmp.SheetOverlayProps{For: "demo-sheet", Behavior: "ui8kit"})
		@cmp.SheetContent(cmp.SheetContentProps{}) { Panel }
	}
}
```
