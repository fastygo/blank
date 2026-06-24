---
id: components.breadcrumb
layer: composite
kind: navigation
package: github.com/fastygo/templ/components/breadcrumb
facade: github.com/fastygo/templ/components
templ: Breadcrumb
api:
  Items:
    role: trail
    type: "[]BreadcrumbItem"
  Class:
    role: style-extension
    type: string
  AriaLabel:
    role: accessible-name
    type: string
  DataUI8Kit:
    role: behavior-hook
    type: string
  Attrs:
    role: html-attrs
    type: templ.Attributes
item-props:
  Label:
    role: crumb-label
    type: string
  Href:
    role: crumb-link
    type: string
  Current:
    role: current-page
    type: bool
  Disabled:
    role: disabled-crumb
    type: bool
showcase:
  - id: layout.standard
    props:
      Items:
        - { Label: Home, Href: / }
        - { Label: Docs, Href: /docs }
        - { Label: Button, Current: true }
  - id: state.disabled
    props:
      Items:
        - { Label: Home, Href: / }
        - { Label: Locked, Disabled: true }
semantics:
  root: nav
  list-root: ol
  role: navigation
  aria-label: from AriaLabel prop
  behavior: static
  current-attr: aria-current page

---
## Summary

Breadcrumb shows the current page trail.
Breadcrumb sets aria-current on the active item.

## Use Cases

- Show docs hierarchy above page title
- Mark the current page without a link

## Semantics

- Root element is nav
- AriaLabel and DataUI8Kit are explicit opt-in props
- Items render as ol li elements
- Href renders anchor when link is enabled

## Example layout.standard

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Breadcrumb(cmp.BreadcrumbProps{
		AriaLabel: "Breadcrumb",
		Items: []cmp.BreadcrumbItem{
			{Label: "Home", Href: "/"},
			{Label: "Docs", Href: "/docs"},
			{Label: "Button", Current: true},
		},
	})
}
```

## Example state.disabled

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Breadcrumb(cmp.BreadcrumbProps{
		AriaLabel: "Breadcrumb",
		Items: []cmp.BreadcrumbItem{
			{Label: "Home", Href: "/"},
			{Label: "Locked", Disabled: true},
		},
	})
}
```
