---
id: utils
layer: helper
kind: module
package: github.com/fastygo/templ/utils
facade: github.com/fastygo/templ/utils
go-package: uiutils
exports:
  class:
    - name: Cn
      role: class-merge
      signature: Cn(classes ...string) string
    - name: Compose
      role: variant-compose
      signature: Compose(v Variants, selection map[string]string, extra ...string) string
    - name: Variants
      role: cva-recipe
      type: struct
  attrs:
    - name: MergeAttrs
      role: attr-merge
    - name: DOMAttrs
      role: dom-id-role-tabindex
    - name: ControlAttrs
      role: form-control-attrs
    - name: SwitchAttrs
      role: switch-aria
  aria:
    - name: AriaExpanded
    - name: AriaControls
    - name: AriaCurrent
    - name: AriaLive
    - name: AriaModal
    - name: AriaLabel
    - name: AriaPressed
    - name: AriaHasPopup
  recipes:
    - name: InputChrome
      used-by: [ui.input, ui.textarea, ui.select]
      keys: [variant]
    - name: InputSize
      used-by: [ui.input, ui.textarea, ui.select]
      keys: [size]
    - name: ControlChrome
      used-by: [ui.checkbox, ui.radio, ui.switch]
      keys: [variant]
    - name: ControlSize
      used-by: [ui.checkbox, ui.radio, ui.switch]
      keys: [size]
    - name: CardVariants
      used-by: [components.card]
      keys: [variant]
    - name: AlertVariants
      used-by: [components.alert]
      keys: [variant]
  helpers:
    - name: InputClasses
      role: input-surface-classes
    - name: ControlClasses
      role: toggle-surface-classes
    - name: CardClasses
      role: card-surface-classes
    - name: AlertClasses
      role: alert-surface-classes
    - name: AlertAttrs
      role: alert-static-aria
    - name: BreadcrumbRootAttrs
      role: breadcrumb-nav-explicit-attrs
      signature: BreadcrumbRootAttrs(ariaLabel, dataUI8Kit string, attrs templ.Attributes) templ.Attributes
    - name: BreadcrumbItemAttrs
      role: breadcrumb-item-aria
    - name: TitleTag
      role: heading-level-resolver
    - name: DefaultInputType
      role: input-type-default
    - name: DefaultButtonType
      role: button-type-default
    - name: TextareaRows
      role: textarea-rows-default
companion: utils/tags.spec.md
examples:
  - id: compose.variant
  - id: merge.attrs
semantics:
  copy-once: true
  import-path: github.com/fastygo/templ/utils
---

## Summary

Utils holds shared Go helpers for registry bricks.
Utils provides class recipes, attrs merge, and ARIA helpers.

## Use Cases

- Copy utils once into a consumer app module
- Compose variant classes in button and input bricks
- Merge templ attributes without duplicate keys

## Semantics

- Package name in consumer code is uiutils
- Variants struct follows cva-style Base plus ByKey maps
- Recipes are vars; bricks call Compose or helper wrappers

## allow-list-source registry

Brick `api.*.allow-list-source` values resolve here:

| Source id | Code location |
|-----------|---------------|
| `utils.tags.TagGroupLayout` | `utils/tags.go` — Block, Box |
| `utils.tags.TagGroupStack` | `utils/tags.go` — Stack |
| `utils.tags.TagGroupGroup` | `utils/tags.go` — Group |
| `utils.tags.TagGroupContainer` | `utils/tags.go` — Container |
| `utils.tags.TagGroupList` | `utils/tags.go` — List |
| `utils.tags.TagGroupText` | `utils/tags.go` — Text |
| `utils.recipes.InputChrome` | `utils/utils.go` — InputChrome.ByKey.variant |
| `utils.recipes.ControlChrome` | `utils/utils.go` — ControlChrome.ByKey.variant |
| `utils.recipes.CardVariants` | `utils/utils.go` — CardVariants.ByKey.variant |
| `utils.recipes.AlertVariants` | `utils/utils.go` — AlertVariants.ByKey.variant |
| `utils.helpers.TitleTag` | `utils/utils.go` — TitleTag function |
| `ui.button.ButtonVariants` | `ui/button/button.templ` — brick-local Variants var |
| `ui.badge.BadgeVariants` | `ui/badge/badge.templ` — brick-local Variants var |

## Example compose.variant

```go
import uiutils "github.com/fastygo/templ/utils"

classes := uiutils.Compose(uiutils.InputChrome, map[string]string{
	"variant": "outline",
}, "mt-2")
```

## Example merge.attrs

```go
attrs := uiutils.MergeAttrs(
	uiutils.AriaLabel("Search"),
	templ.Attributes{"name": "q"},
)
```
