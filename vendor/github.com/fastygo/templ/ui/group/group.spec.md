---
id: ui.group
layer: primitive
kind: layout
package: github.com/fastygo/templ/ui/group
facade: github.com/fastygo/templ/ui
templ: Group
api:
  Tag:
    role: group-tag
    type: string
    enum: [div, section, article, aside, header, footer, main, nav, figure, fieldset]
    allow-list-source: utils.tags.TagGroupGroup
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
    props: { Tag: div, Class: "gap-2" }
  - id: tag.fieldset
    props: { Tag: fieldset, Class: "gap-2" }
semantics:
  root: resolved from Tag
  role: none
  behavior: static
  layout: flex-row

---
## Summary

Group arranges children in a horizontal flex row.
Group supports fieldset for related form controls.

## Use Cases

- Place buttons side by side with gap utilities
- Group radio options inside a fieldset

## Semantics

- Base classes include flex min-w-0
- ResolveTag uses TagGroupGroup from utils

## Example tag.default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "gap-2"}) {
		@ui.Button(ui.ButtonProps{Variant: "default"}) { Save }
		@ui.Button(ui.ButtonProps{Variant: "outline"}) { Cancel }
	}
}
```

## Example tag.fieldset

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Tag: "fieldset", Class: "gap-4"}) {
		@ui.Radio(ui.RadioProps{Name: "plan", Value: "free", ID: "plan-free"})
		@ui.Radio(ui.RadioProps{Name: "plan", Value: "pro", ID: "plan-pro"})
	}
}
```
