---
id: ui.form
layer: composite
kind: input
package: github.com/fastygo/templ/ui/form
facade: github.com/fastygo/templ/ui
parts:
  - templ: Form
    props: [ID, Class, Action, Method, Enctype, Autocomplete, Name, Target, NoValidate, Attrs]
    slot: root
  - templ: FormItem
    props: [Class]
    slot: item
  - templ: FormDescription
    props: [Class]
    slot: description
    value: string
  - templ: FormMessage
    props: [Class]
    slot: message
    value: string
api:
  Action:
    role: submit-url
    type: string
  Method:
    role: http-method
    type: string
    enum: [get, post]
  NoValidate:
    role: validation
    type: bool
showcase:
  - id: layout.field-with-label
    parts: [Form, FormItem, FormDescription]
    props: { Method: post }
  - id: layout.with-message
    parts: [Form, FormItem, FormMessage]
semantics:
  root: form
  role: form
  behavior: interactive
  message-role: alert

---
## Summary

Form wraps fields in a vertical grid container.
FormItem groups label, control, and helper text.

## Use Cases

- Submit user data with Method post
- Show field hint with FormDescription
- Show validation error with FormMessage

## Semantics

- Form root uses grid gap-4 layout
- FormMessage renders role alert on p element

## Example layout.field-with-label

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Form(ui.FormProps{Action: "/signup", Method: "post"}) {
		@ui.FormItem(ui.FormItemProps{}) {
			@ui.Label(ui.LabelProps{HTMLFor: "email"}) { Email }
			@ui.Input(ui.InputProps{ID: "email", Name: "email", Type: "email"})
			@ui.FormDescription(ui.FormDescriptionProps{}, "We never share your email.")
		}
	}
}
```

## Example layout.with-message

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Form(ui.FormProps{}) {
		@ui.FormItem(ui.FormItemProps{}) {
			@ui.Label(ui.LabelProps{HTMLFor: "name"}) { Name }
			@ui.Input(ui.InputProps{ID: "name", Name: "name", Required: true})
			@ui.FormMessage(ui.FormMessageProps{}, "Name is required.")
		}
	}
}
```
