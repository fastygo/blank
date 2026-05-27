---
id: ui.button
layer: primitive
kind: action
package: github.com/fastygo/templ/ui/button
facade: github.com/fastygo/templ/ui
templ: Button
api:
  Variant:
    role: appearance
    type: string
    enum: [default, secondary, destructive, outline, ghost, link, unstyled]
    allow-list-source: ui.button.ButtonVariants
    default: default
  Size:
    role: density
    type: string
    enum: [default, sm, lg, icon]
    allow-list-source: ui.button.ButtonVariants
    default: default
  Class:
    role: style-extension
    type: string
  Type:
    role: control-type
    type: string
    enum: [button, submit, reset]
    default: button
  Form:
    role: association
    type: string
  Disabled:
    role: state
    type: bool
    default: false
  Href:
    role: navigation
    type: string
  AriaLabel:
    role: accessible-name
    type: string
  ID:
    role: identity
    type: string
  Attrs:
    role: html-attrs
    type: templ.Attributes
slots:
  default:
    required: true
    accepts: text
showcase:
  - id: variant.default
    props: { Variant: default, Size: default }
  - id: variant.secondary
    props: { Variant: secondary, Size: default }
  - id: variant.destructive
    props: { Variant: destructive, Size: default }
  - id: variant.outline
    props: { Variant: outline, Size: sm }
  - id: variant.ghost
    props: { Variant: ghost, Size: sm }
  - id: variant.link
    props: { Variant: link, Href: "/docs" }
  - id: variant.unstyled
    props: { Variant: unstyled, Class: "underline" }
  - id: size.sm
    props: { Variant: default, Size: sm }
  - id: size.lg
    props: { Variant: default, Size: lg }
  - id: size.icon
    props: { Variant: default, Size: icon, AriaLabel: "Settings" }
  - id: state.disabled
    props: { Variant: default, Disabled: true }
  - id: state.disabled-link
    props: { Variant: link, Href: "/locked", Disabled: true }
  - id: state.submit
    props: { Variant: default, Type: submit }
  - id: state.reset
    props: { Variant: outline, Type: reset }
semantics:
  root: button
  root-when-href: a
  role: button
  role-when-href: link
  behavior: interactive

---
## Summary

Button triggers an action on click.
Button navigates when Href is set.

## Use Cases

- Submit a form action
- Open a page from a primary action
- Show a destructive confirm action
- Show a low-emphasis ghost action

## Semantics

- Root element is button when Href is empty
- Root element is a when Href is set
- Disabled sets native disabled on button
- Disabled link sets aria-disabled true

## Example variant.default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "default"}) {
		Primary
	}
}
```

## Example variant.secondary

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "secondary"}) {
		Secondary
	}
}
```

## Example variant.destructive

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "destructive"}) {
		Delete
	}
}
```

## Example variant.outline

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "outline", Size: "sm"}) {
		Outline
	}
}
```

## Example variant.ghost

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "ghost", Size: "sm"}) {
		Ghost
	}
}
```

## Example variant.link

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "link", Href: "/docs"}) {
		Learn more
	}
}
```

## Example variant.unstyled

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "unstyled", Class: "underline"}) {
		Custom
	}
}
```

## Example size.sm

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Size: "sm"}) {
		Small
	}
}
```

## Example size.lg

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Size: "lg"}) {
		Large
	}
}
```

## Example size.icon

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Size: "icon", AriaLabel: "Settings"}) {
		+
	}
}
```

## Example state.disabled

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Disabled: true}) {
		Disabled
	}
}
```

## Example state.submit

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Type: "submit"}) {
		Submit
	}
}
```

## Example state.disabled-link

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "link", Href: "/locked", Disabled: true}) {
		Disabled link
	}
}
```

## Example state.reset

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "outline", Type: "reset"}) {
		Reset
	}
}
```
