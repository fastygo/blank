---
id: ui.switch
layer: primitive
kind: input
package: github.com/fastygo/templ/ui/switch
facade: github.com/fastygo/templ/ui
templ: Switch
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
    role: form-name
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
  - id: state.off
    props: { Name: notifications, ID: notifications, AriaLabel: "Notifications" }
  - id: state.on
    props: { Name: notifications, ID: notifications, Checked: true, AriaLabel: "Notifications" }
  - id: state.disabled
    props: { Name: locked, Disabled: true, Checked: true, AriaLabel: "Locked setting" }
semantics:
  root: input
  input-type: checkbox
  role: switch
  aria-checked: dynamic
  behavior: interactive

---
## Summary

Switch renders a toggle with switch ARIA semantics.
SwitchAttrs sets role switch and aria-checked.

## Use Cases

- Toggle notifications on or off
- Toggle feature flags in settings forms

## Semantics

- Native input type is checkbox
- SwitchAttrs overrides role to switch
- aria-checked mirrors Checked prop at render time

## Example state.off

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Switch(ui.SwitchProps{
		Name:      "notifications",
		ID:        "notifications",
		AriaLabel: "Notifications",
	})
}
```

## Example state.on

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Switch(ui.SwitchProps{
		Name:      "notifications",
		ID:        "notifications",
		Checked:   true,
		AriaLabel: "Notifications",
	})
}
```

## Example state.disabled

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Switch(ui.SwitchProps{
		Name:      "locked",
		Disabled:  true,
		Checked:   true,
		AriaLabel: "Locked setting",
	})
}
```
