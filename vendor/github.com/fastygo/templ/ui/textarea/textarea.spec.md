---
id: ui.textarea
layer: primitive
kind: input
package: github.com/fastygo/templ/ui/textarea
facade: github.com/fastygo/templ/ui
templ: Textarea
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
  Rows:
    role: visible-lines
    type: int
    default: 4
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
showcase:
  - id: variant.default
    props: { Name: message, Placeholder: "Your message" }
  - id: state.disabled
    props: { Disabled: true, Name: locked, Value: "Fixed text" }
  - id: rows.custom
    props: { Rows: 6, Name: bio, Placeholder: "Bio" }
semantics:
  root: textarea
  role: textbox
  behavior: interactive

---
## Summary

Textarea renders a multi-line native text field.
TextareaRows defaults Rows to four when zero.

## Use Cases

- Collect long feedback in a contact form
- Collect bio text with custom row count

## Semantics

- InputClasses recipe from utils styles the surface
- Value renders as textarea body content

## Example variant.default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Textarea(ui.TextareaProps{
		Name:        "message",
		Placeholder: "Your message",
	})
}
```

## Example state.disabled

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Textarea(ui.TextareaProps{
		Disabled: true,
		Name:     "locked",
		Value:    "Fixed text",
	})
}
```

## Example rows.custom

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Textarea(ui.TextareaProps{
		Rows:        6,
		Name:        "bio",
		Placeholder: "Bio",
	})
}
```
