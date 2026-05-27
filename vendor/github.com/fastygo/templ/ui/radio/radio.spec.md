---
id: ui.radio
layer: primitive
kind: input
package: github.com/fastygo/templ/ui/radio
facade: github.com/fastygo/templ/ui
templ: Radio
api:
  Variant:
    role: appearance
    type: string
    enum: [default, outline, ghost, unstyled]
    allow-list-source: utils.recipes.ControlChrome
    default: default
  Size:
    role: density
    type: string
    enum: [default, sm, lg, xs]
    allow-list-source: utils.recipes.ControlSize
    default: default
  Name:
    role: group-name
    type: string
  Value:
    role: option-value
    type: string
  Checked:
    role: state
    type: bool
  Disabled:
    role: state
    type: bool
  AriaLabel:
    role: accessible-name
    type: string
showcase:
  - id: layout.group
    props: { Name: plan }
  - id: state.disabled
    props: { Name: plan, Value: pro, Disabled: true }
semantics:
  root: input
  input-type: radio
  role: radio
  behavior: interactive
  rule: same-name-groups-options

---
## Summary

Radio renders one option in a named radio group.
All radios in a group share the same Name value.

## Use Cases

- Pick one plan in a pricing form
- Pick one payment method in checkout

## Semantics

- type attribute is always radio
- Checked marks the selected option in the group

## Example layout.group

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2"}) {
		@ui.Radio(ui.RadioProps{Name: "plan", Value: "free", ID: "plan-free", Checked: true})
		@ui.Radio(ui.RadioProps{Name: "plan", Value: "pro", ID: "plan-pro"})
	}
}
```

## Example state.disabled

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Radio(ui.RadioProps{Name: "plan", Value: "pro", ID: "plan-pro", Disabled: true})
}
```
