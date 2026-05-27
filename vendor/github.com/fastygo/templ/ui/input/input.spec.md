---
id: ui.input
layer: primitive
kind: input
package: github.com/fastygo/templ/ui/input
facade: github.com/fastygo/templ/ui
templ: Input
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
  Type:
    role: input-type
    type: string
    enum: [text, email, password, number, search, tel, url]
    default: text
  Class:
    role: style-extension
    type: string
  Name:
    role: form-name
    type: string
  Value:
    role: value
    type: string
  Placeholder:
    role: placeholder
    type: string
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
  - id: variant.default
    props: { Variant: default, Type: text, Name: email, Placeholder: "name@example.com" }
  - id: type.email
    props: { Type: email, Name: email, ID: demo-email }
  - id: size.sm
    props: { Size: sm, Name: q, Placeholder: "Search" }
  - id: state.disabled
    props: { Disabled: true, Name: locked, Value: "Read only" }
  - id: state.required
    props: { Required: true, Name: name, Placeholder: "Required field" }
semantics:
  root: input
  role: textbox
  behavior: interactive

---
## Summary

Input renders a single-line native text field.
Input uses InputClasses recipe from utils.

## Use Cases

- Collect email in a login form
- Collect search query in a toolbar field

## Semantics

- DefaultInputType returns text when Type is empty
- ControlAttrs merges aria-label without duplicate id
- Disabled sets native disabled attribute

## Example variant.default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{
		Name:        "email",
		Placeholder: "name@example.com",
	})
}
```

## Example type.email

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{
		ID:   "demo-email",
		Type: "email",
		Name: "email",
	})
}
```

## Example size.sm

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{
		Size:        "sm",
		Name:        "q",
		Placeholder: "Search",
	})
}
```

## Example state.disabled

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{
		Disabled: true,
		Name:     "locked",
		Value:    "Read only",
	})
}
```

## Example state.required

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{
		Required:    true,
		Name:        "name",
		Placeholder: "Required field",
	})
}
```
