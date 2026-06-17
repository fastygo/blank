---
id: ui.linebreak
layer: primitive
kind: typography
package: github.com/fastygo/templ/ui/linebreak
facade: github.com/fastygo/templ/ui
templ: Break
api:
  Type:
    role: break-type
    type: string
    enum: [br, wbr, word]
    default: br
  Class:
    role: style-extension
    type: string
showcase:
  - id: type.br
    props: { Type: br }
  - id: type.wbr
    props: { Type: wbr }
semantics:
  root: br | wbr
  behavior: static
---

## Summary

Break renders br or wbr for controlled inline wrapping.

## Use Cases

- Render intentional line breaks in addresses
- Add optional break points in long tokens

## Semantics

- Type br renders a line break
- Type wbr or word renders an optional word break

## Example type.br

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Break(ui.BreakProps{})
}
```

## Example type.wbr

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Break(ui.BreakProps{Type: "wbr"})
}
```
