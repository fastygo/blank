---
id: ui.select
layer: primitive
kind: input
package: github.com/fastygo/templ/ui/select
facade: github.com/fastygo/templ/ui
templ: Select
api:
  Variant:
    role: appearance
    type: string
    enum: [default, outline, ghost, unstyled]
    allow-list-source: utils.recipes.InputChrome
    default: default
  Size:
    role: density
    type: string
    enum: [default, sm, lg, xs]
    allow-list-source: utils.recipes.InputSize
    default: default
  Name:
    role: form-name
    type: string
  Value:
    role: selected-value
    type: string
  Options:
    role: choices
    type: "[]Option"
  Disabled:
    role: state
    type: bool
  Required:
    role: state
    type: bool
showcase:
  - id: variant.default
    props:
      Name: country
      Value: us
      Options:
        - { Value: us, Label: United States }
        - { Value: ca, Label: Canada }
  - id: state.disabled
    props:
      Disabled: true
      Name: locked
      Options:
        - { Value: a, Label: Option A }
semantics:
  root: select
  role: combobox
  behavior: interactive

---
## Summary

Select renders a native dropdown from Options slice.
Select marks selected option when Value matches Option Value.

## Use Cases

- Pick one country in an address form
- Pick one category in a filter form

## Semantics

- InputClasses recipe from utils styles the surface
- Options render as option elements inside select

## Example variant.default

```templ
import "github.com/fastygo/templ/ui"
import "github.com/fastygo/templ/ui/select"

templ Example() {
	@ui.Select(ui.SelectProps{
		Name:  "country",
		Value: "us",
		Options: []selectfield.Option{
			{Value: "us", Label: "United States"},
			{Value: "ca", Label: "Canada"},
		},
	})
}
```

## Example state.disabled

```templ
import "github.com/fastygo/templ/ui"
import "github.com/fastygo/templ/ui/select"

templ Example() {
	@ui.Select(ui.SelectProps{
		Disabled: true,
		Name:     "locked",
		Options: []selectfield.Option{
			{Value: "a", Label: "Option A"},
		},
	})
}
```
