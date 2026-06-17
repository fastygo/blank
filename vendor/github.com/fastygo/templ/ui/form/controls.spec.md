---
id: ui.form.controls
layer: primitive
kind: form
package: github.com/fastygo/templ/ui/form
facade: github.com/fastygo/templ/ui
templ: Fieldset
parts:
  - templ: Fieldset
    props: [Variant, Size, Class, Name, Form, Disabled, Attrs]
  - templ: Legend
    props: [Variant, Size, Class, Attrs]
  - templ: DataList
    props: [ID, Class, Attrs]
  - templ: DataOption
    props: [Value, Label, Attrs]
  - templ: Output
    props: [ID, Class, Name, For, Value, Attrs]
  - templ: Meter
    props: [Variant, Size, ID, Class, Value, Min, Max, Low, High, Optimum, Attrs]
  - templ: Progress
    props: [Variant, Size, ID, Class, Value, Max, Attrs]
api:
  FieldsetVariant:
    role: appearance
    type: string
    enum: [default, card, inline, unstyled]
    allow-list-source: ui.form.FieldsetVariants.variant
    default: default
  LegendVariant:
    role: appearance
    type: string
    enum: [default, muted, unstyled]
    allow-list-source: ui.form.LegendVariants.variant
    default: default
  MeterVariant:
    role: appearance
    type: string
    enum: [default, muted, unstyled]
    allow-list-source: ui.form.MeterVariants.variant
    default: default
  ProgressVariant:
    role: appearance
    type: string
    enum: [default, muted, unstyled]
    allow-list-source: ui.form.ProgressVariants.variant
    default: default
showcase:
  - id: group.fieldset
    props: { Variant: card }
  - id: status.progress
    props: { Value: "40", Max: "100" }
semantics:
  root: fieldset | legend | datalist | option | output | meter | progress
  behavior: native-form
---

## Summary

Form controls covers fieldset, legend, datalist, output, meter, and progress.

## Use Cases

- Group related form controls
- Render autocomplete values, calculated output, quotas, and progress

## Semantics

- Fieldset and Legend preserve native grouping semantics
- Meter and Progress use native status elements

## Example group.fieldset

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Fieldset(ui.FieldsetProps{Variant: "card"}) {
		@ui.Legend(ui.LegendProps{}) { Account }
		@ui.Input(ui.InputProps{Name: "email"})
	}
}
```

## Example status.progress

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Progress(ui.ProgressProps{Value: "40", Max: "100"})
}
```
