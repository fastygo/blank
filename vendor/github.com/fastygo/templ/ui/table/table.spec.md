---
id: ui.table
layer: composite
kind: data
package: github.com/fastygo/templ/ui/table
facade: github.com/fastygo/templ/ui
parts:
  - templ: Table
    props: [Class, Attrs]
  - templ: TableCaption
    props: [Class]
    slot: caption
  - templ: TableHead
    props: [Class]
    slot: head
  - templ: TableBody
    props: [Class]
    slot: body
  - templ: TableFoot
    props: [Class]
    slot: foot
  - templ: TableRow
    props: [Class]
    slot: row
  - templ: TableHeadCell
    props: [Class, Scope, Abbr]
    slot: head-cell
  - templ: TableCell
    props: [Class, ColSpan, RowSpan, Headers]
    slot: cell
api:
  Scope:
    role: header-scope
    type: string
    enum: [col, row, colgroup, rowgroup]
    applies-to: TableHeadCell
  ColSpan:
    role: column-span
    type: int
    applies-to: TableCell
showcase:
  - id: layout.basic
    parts: [Table, TableHead, TableBody, TableRow, TableHeadCell, TableCell]
    props: { Scope: col }
  - id: cell.colspan
    parts: [Table, TableBody, TableRow, TableCell]
semantics:
  root: table
  role: table
  behavior: static

---
## Summary

Table composes caption, sections, rows, and cells.
TableHeadCell supports scope on th elements.

## Use Cases

- Show user list with column headers
- Span one cell across two columns with ColSpan

## Semantics

- Table root is native table element
- TableHeadCell sets scope when Scope prop is valid

## Example layout.basic

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Table(ui.TableProps{Class: "w-full text-sm"}) {
		@ui.TableHead(ui.TableSectionProps{}) {
			@ui.TableRow(ui.TableRowProps{}) {
				@ui.TableHeadCell(ui.TableCellProps{Scope: "col"}) { Name }
				@ui.TableHeadCell(ui.TableCellProps{Scope: "col"}) { Role }
			}
		}
		@ui.TableBody(ui.TableSectionProps{}) {
			@ui.TableRow(ui.TableRowProps{}) {
				@ui.TableCell(ui.TableCellProps{}) { Alex }
				@ui.TableCell(ui.TableCellProps{}) { Admin }
			}
		}
	}
}
```

## Example cell.colspan

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Table(ui.TableProps{}) {
		@ui.TableBody(ui.TableSectionProps{}) {
			@ui.TableRow(ui.TableRowProps{}) {
				@ui.TableCell(ui.TableCellProps{ColSpan: 2}) { Spans two columns }
			}
		}
	}
}
```
