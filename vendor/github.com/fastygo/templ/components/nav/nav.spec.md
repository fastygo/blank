---
id: components.nav
layer: composite
kind: navigation
package: github.com/fastygo/templ/components/nav
facade: github.com/fastygo/templ/components
templ: Nav
parts:
  - templ: Nav
    props: [Class, AriaLabel, DataUI8Kit, Attrs]
  - templ: NavList
    props: [Class, Orientation, Gap, Attrs]
  - templ: NavItem
    props: [Class, Attrs]
  - templ: NavLink
    props: [Href, Variant, Size, Class, Active, Disabled, AriaCurrent, AriaLabel, Attrs]
api:
  Orientation:
    role: direction
    type: string
    enum: [vertical, horizontal]
    allow-list-source: components.nav.NavListVariants
    default: vertical
  LinkVariant:
    role: appearance
    type: string
    enum: [default, active, ghost, muted, unstyled]
    allow-list-source: components.nav.NavLinkVariants.variant
    default: default
  Size:
    role: density
    type: string
    enum: [sm, default, lg]
    default: default
showcase:
  - id: layout.vertical
    props: { AriaLabel: "Primary navigation" }
  - id: state.active
    props: { Active: true }
semantics:
  root: nav
  list-root: ul
  item-root: li
  link-root: a | span
  behavior: static
---

## Summary

Nav composes labeled navigation lists with active and disabled links.

## Use Cases

- Render sidebars and top navigation
- Render current-page states without custom class helpers

## Semantics

- Nav root accepts an explicit AriaLabel
- Active links default aria-current to page
- Disabled or empty links render span with aria-disabled

## Example layout.vertical

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Nav(cmp.NavProps{AriaLabel: "Primary navigation"}) {
		@cmp.NavList(cmp.NavListProps{}) {
			@cmp.NavItem(cmp.NavItemProps{}) {
				@cmp.NavLink(cmp.NavLinkProps{Href: "/", Active: true}) { Home }
			}
		}
	}
}
```

## Example state.active

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.NavLink(cmp.NavLinkProps{Href: "/docs", Active: true}) { Docs }
}
```
