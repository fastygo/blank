---
id: ui.container
layer: primitive
kind: layout
package: github.com/fastygo/templ/ui/container
facade: github.com/fastygo/templ/ui
templ: Container
api:
  Tag:
    role: shell-tag
    type: string
    enum: [div, main, section]
    allow-list-source: utils.tags.TagGroupContainer
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
    props: { Tag: div, Class: "mx-auto max-w-5xl px-6" }
  - id: tag.main
    props: { Tag: main, Class: "mx-auto max-w-5xl px-6" }
semantics:
  root: resolved from Tag
  role: none
  behavior: static

---
## Summary

Container wraps page content in a width shell.
Container supports div, main, or section root tags.

## Use Cases

- Center content with max width utilities on Class
- Wrap main element for app shell layouts

## Semantics

- ResolveTag uses TagGroupContainer from utils
- Invalid Tag falls back to div

## Example tag.default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Container(ui.ContainerProps{Class: "mx-auto max-w-5xl px-6"}) {
		Page width content
	}
}
```

## Example tag.main

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Container(ui.ContainerProps{Tag: "main", Class: "mx-auto max-w-5xl px-6"}) {
		Main shell content
	}
}
```
