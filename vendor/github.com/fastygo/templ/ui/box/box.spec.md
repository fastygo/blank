---
id: ui.box
layer: primitive
kind: layout
package: github.com/fastygo/templ/ui/box
facade: github.com/fastygo/templ/ui
templ: Box
api:
  Tag:
    role: layout-tag
    type: string
    enum: [div, section, article, aside, header, footer, main, nav, figure, search, hgroup]
    allow-list-source: utils.tags.TagGroupLayout
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
    props: { Tag: div }
  - id: tag.section
    props: { Tag: section, Class: "p-4" }
semantics:
  root: resolved from Tag
  role: none
  behavior: static
  rule: internal-layout-only-not-landmark

---
## Summary

Box wraps internal layout sections inside a Block.
Box is not for top-level page landmarks.

## Use Cases

- Pad content inside a card body
- Group inline elements with flex utilities on Class

## Semantics

- ResolveTag uses TagGroupLayout from utils
- Attrs spread directly on the root element
- Prefer Block for main, header, and footer regions

## Example tag.default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Box(ui.BoxProps{Class: "mx-auto max-w-3xl px-6"}) {
		Inner content
	}
}
```

## Example tag.section

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Box(ui.BoxProps{Tag: "section", Class: "p-4"}) {
		Section body
	}
}
```
