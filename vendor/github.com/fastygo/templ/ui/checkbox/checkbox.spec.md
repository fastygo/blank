---
id: ui.checkbox
layer: primitive
kind: input
package: github.com/fastygo/templ/ui/checkbox
facade: github.com/fastygo/templ/ui
templ: Checkbox
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
  Value:
    role: submit-value
    type: string
  Checked:
    role: state
    type: bool
  Disabled:
    role: state
    type: bool
  Required:
    role: state
    type: bool
  AriaLabel:
    role: accessible-name
    type: string
showcase:
  - id: state.unchecked
    props: { Name: terms, ID: terms }
  - id: state.checked
    props: { Name: terms, ID: terms, Checked: true }
  - id: state.disabled
    props: { Name: locked, Disabled: true, Checked: true }
semantics:
  root: input
  input-type: checkbox
  role: checkbox
  behavior: interactive

---
## Summary

Checkbox renders a native boolean form control.
Checkbox uses ControlClasses recipe from utils.

## Use Cases

- Accept terms in a signup form
- Toggle one optional setting in a list

## Semantics

- type attribute is always checkbox
- ControlAttrs merges aria-label without duplicate id

## Example state.unchecked

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Checkbox(ui.CheckboxProps{Name: "terms", ID: "terms"})
}
```

## Example state.checked

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Checkbox(ui.CheckboxProps{Name: "terms", ID: "terms", Checked: true})
}
```

## Example state.disabled

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Checkbox(ui.CheckboxProps{Name: "locked", Disabled: true, Checked: true})
}
```
