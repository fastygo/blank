---
id: ui.list
layer: primitive
kind: layout
package: github.com/fastygo/templ/ui/list
facade: github.com/fastygo/templ/ui
templ: List
parts:
  - templ: List
    props: [Class, Tag, Attrs]
  - templ: ListItem
    props: [Class, Tag, Value, Attrs]
    slot: item
api:
  Tag:
    role: list-tag
    type: string
    enum: [ul, ol, dl, menu]
    allow-list-source: utils.tags.TagGroupList
    default: ul
  Value:
    role: ordered-value
    type: int
    applies-to: ListItem
  ItemTag:
    role: item-tag
    type: string
    enum: [li, dt, dd]
    allow-list-source: utils.tags.TagGroupListItem
    default: li
showcase:
  - id: tag.ul
    props: { Tag: ul, Class: "gap-2" }
  - id: tag.ol
    props: { Tag: ol, Class: "list-decimal pl-4" }
semantics:
  root: ul | ol | dl | menu
  item-root: li | dt | dd
  behavior: static

---
## Summary

List renders ul, ol, dl, or menu containers.
ListItem renders one li row inside List.

## Use Cases

- Show bullet navigation links
- Show ordered steps in a wizard

## Semantics

- ResolveTag uses TagGroupList from utils
- ListItem Value sets li value attribute when positive

## Example tag.ul

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.List(ui.ListProps{Class: "gap-2"}) {
		@ui.ListItem(ui.ListItemProps{}) { First item }
		@ui.ListItem(ui.ListItemProps{}) { Second item }
	}
}
```

## Example tag.ol

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.List(ui.ListProps{Tag: "ol", Class: "list-decimal pl-4 gap-2"}) {
		@ui.ListItem(ui.ListItemProps{Value: 1}) { Step one }
		@ui.ListItem(ui.ListItemProps{Value: 2}) { Step two }
	}
}
```
