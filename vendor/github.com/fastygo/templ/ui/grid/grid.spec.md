---
id: ui.grid
layer: primitive
kind: layout
package: github.com/fastygo/templ/ui/grid
facade: github.com/fastygo/templ/ui
templ: Grid
parts:
  - templ: Grid
    props: [Class]
  - templ: GridCol
    props: [Class]
    slot: column
api:
  Class:
    role: style-extension
    type: string
showcase:
  - id: layout.two-col
    props: { Class: "grid-cols-2 gap-4" }
  - id: layout.three-col
    props: { Class: "grid-cols-3 gap-4" }
semantics:
  root: div
  role: none
  behavior: static
  layout: grid
---

## Summary

Grid lays out children in a CSS grid container.
GridCol wraps one column cell inside Grid.

## Use Cases

- Show two KPI cards in one row
- Split form label and control in two columns

## Semantics

- Grid root is div with grid class
- GridCol root is div without default grid class

## Example layout.two-col

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Grid(ui.GridProps{Class: "grid-cols-2 gap-4"}) {
		@ui.GridCol(ui.GridColProps{}) {
			@ui.Text(ui.TextProps{}, "Column A")
		}
		@ui.GridCol(ui.GridColProps{}) {
			@ui.Text(ui.TextProps{}, "Column B")
		}
	}
}
```

## Example layout.three-col

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Grid(ui.GridProps{Class: "grid-cols-3 gap-4"}) {
		@ui.GridCol(ui.GridColProps{}) { A }
		@ui.GridCol(ui.GridColProps{}) { B }
		@ui.GridCol(ui.GridColProps{}) { C }
	}
}
```
