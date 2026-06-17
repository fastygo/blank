---
id: ui.block
layer: primitive
kind: layout
package: github.com/fastygo/templ/ui/block
facade: github.com/fastygo/templ/ui
templ: Block
api:
  Tag:
    role: landmark-tag
    type: string
    enum: [div, main, header, section, footer, nav, article, aside, figure, search, hgroup]
    default: div
    allow-list-source: utils.tags.TagGroupLayout
  Class:
    role: style-extension
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
    accepts: any
showcase:
  - id: tag.main
    props: { Tag: main, Class: "min-h-screen" }
  - id: tag.section
    props: { Tag: section, Class: "py-8" }
  - id: tag.header
    props: { Tag: header, Class: "border-b border-border" }
  - id: tag.footer
    props: { Tag: footer, Class: "border-t border-border" }
  - id: tag.nav
    props: { Tag: nav, Class: "py-2" }
semantics:
  root: resolved from Tag
  role: none
  behavior: static
  rule: do-not-nest-block-in-block

---
## Summary

Block renders a top-level page landmark element.
Do not nest Block inside another Block.

## Use Cases

- Define main page region with Tag main
- Define page header and footer landmarks
- Define content sections with Tag section

## Semantics

- All values in api.Tag.enum are valid landmarks
- ResolveTag uses TagGroupLayout from utils.tags
- Invalid Tag falls back to div

## Example tag.main

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Block(ui.BlockProps{Tag: "main", Class: "min-h-screen"}) {
		Page content
	}
}
```

## Example tag.section

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Block(ui.BlockProps{Tag: "section", Class: "py-8"}) {
		Section content
	}
}
```

## Example tag.header

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Block(ui.BlockProps{Tag: "header", Class: "border-b border-border"}) {
		Site header
	}
}
```

## Example tag.footer

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Block(ui.BlockProps{Tag: "footer", Class: "border-t border-border"}) {
		Site footer
	}
}
```

## Example tag.nav

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Block(ui.BlockProps{Tag: "nav", Class: "py-2"}) {
		Primary navigation
	}
}
```
