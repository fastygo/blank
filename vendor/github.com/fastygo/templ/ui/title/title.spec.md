---
id: ui.title
layer: primitive
kind: typography
package: github.com/fastygo/templ/ui/title
facade: github.com/fastygo/templ/ui
templ: Title
api:
  Order:
    role: heading-level
    type: int
    enum: [1, 2, 3, 4, 5, 6]
    allow-list-source: utils.helpers.TitleTag
    default: 2
  Class:
    role: style-extension
    type: string
  value:
    role: copy
    type: string
    position: argument
showcase:
  - id: order.h1
    props: { Order: 1 }
  - id: order.h2
    props: { Order: 2 }
  - id: order.h3
    props: { Order: 3, Class: "text-lg" }
semantics:
  root: h1 | h2 | h3 | h4 | h5 | h6
  role: heading
  behavior: static

---
## Summary

Title renders one heading from Order prop.
TitleTag maps Order to h1 through h6.

## Use Cases

- Render page title with Order 1
- Render section title with Order 2

## Semantics

- Order 0 and 2 default to h2
- Order 1 maps to h1, Order 3 to h3

## Example order.h1

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Title(ui.TitleProps{Order: 1}, "Page title")
}
```

## Example order.h2

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Title(ui.TitleProps{Order: 2}, "Section title")
}
```

## Example order.h3

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Title(ui.TitleProps{Order: 3, Class: "text-lg"}, "Subsection title")
}
```
