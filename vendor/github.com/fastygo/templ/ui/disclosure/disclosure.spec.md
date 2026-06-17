---
id: ui.disclosure
layer: primitive
kind: interactive
package: github.com/fastygo/templ/ui/disclosure
facade: github.com/fastygo/templ/ui
templ: Disclosure
parts:
  - templ: Disclosure
    props: [Variant, Size, Class, Open, Attrs]
  - templ: Summary
    props: [Variant, Size, Class, Attrs]
    slot: trigger
api:
  Variant:
    role: appearance
    type: string
    enum: [default, ghost, card, unstyled]
    allow-list-source: ui.disclosure.DisclosureVariants
    default: default
  Size:
    role: density
    type: string
    enum: [sm, default, lg]
    default: default
  Open:
    role: state
    type: bool
    default: false
showcase:
  - id: variant.default
    props: { Variant: default }
  - id: state.open
    props: { Open: true }
semantics:
  root: details
  trigger-root: summary
  behavior: native-disclosure
---

## Summary

Disclosure renders native details and summary elements.

## Use Cases

- Render FAQ rows
- Render compact settings and optional form sections

## Semantics

- Disclosure uses native details open state
- Summary must be the visible trigger

## Example variant.default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Disclosure(ui.DisclosureProps{}) {
		@ui.Summary(ui.SummaryProps{}) { More options }
		@ui.Text(ui.TextProps{}, "Hidden content.")
	}
}
```

## Example state.open

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Disclosure(ui.DisclosureProps{Open: true}) {
		@ui.Summary(ui.SummaryProps{}) { Details }
		@ui.Text(ui.TextProps{}, "Visible content.")
	}
}
```
