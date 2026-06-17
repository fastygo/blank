---
id: ui.stack
layer: primitive
kind: layout
package: github.com/fastygo/templ/ui/stack
facade: github.com/fastygo/templ/ui
templ: Stack
api:
  Tag:
    role: stack-tag
    type: string
    enum: [div, nav, section, article, aside, header, footer, main, figure, search, hgroup, ul, ol]
    allow-list-source: utils.tags.TagGroupStack
    default: div
  Class:
    role: style-extension
    type: string
  Attrs:
    role: html-attrs
    type: templ.Attributes
slots:
  default:
    required: true
    accepts: any
showcase:
  - id: tag.default
    props: { Tag: div, Class: "gap-4" }
  - id: tag.ul
    props: { Tag: ul, Class: "gap-2" }
semantics:
  root: resolved from Tag
  role: none
  behavior: static
  layout: flex-col

---
## Summary

Stack arranges children in a vertical flex column.
Stack is the default spacing primitive for sections.

## Use Cases

- Space form fields vertically with gap utilities
- Stack headings and body copy in a section

## Semantics

- Base classes include flex flex-col items-start
- ResolveTag uses TagGroupStack from utils

## Example tag.default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-4"}) {
		@ui.Text(ui.TextProps{}, "First line")
		@ui.Text(ui.TextProps{}, "Second line")
	}
}
```

## Example tag.ul

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Tag: "ul", Class: "gap-2"}) {
		@ui.Text(ui.TextProps{Tag: "span"}, "Item one")
	}
}
```
