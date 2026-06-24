---
id: components.card
layer: composite
kind: surface
package: github.com/fastygo/templ/components/card
facade: github.com/fastygo/templ/components
parts:
  - templ: Card
    props: [Class, Variant, Tag, Attrs]
  - templ: CardHeader
    props: [Class, Attrs]
    slot: header
  - templ: CardTitle
    props: [Class, Order, Attrs]
    slot: title
    value: string
  - templ: CardDescription
    props: [Class, Attrs]
    slot: description
    value: string
  - templ: CardContent
    props: [Class, Attrs]
    slot: content
  - templ: CardFooter
    props: [Class, Attrs]
    slot: footer
api:
  Class:
    role: style-extension
    type: string
  Variant:
    role: appearance
    type: string
    enum: [default, raised, kpi, muted, ghost, compact, flat, accent]
    allow-list-source: utils.recipes.CardVariants
    default: default
  Tag:
    role: root-tag
    type: string
    enum: [div, section, article, aside, header, footer, main, nav, figure, search, hgroup]
    allow-list-source: utils.tags.TagGroupLayout
    default: div
  Attrs:
    role: html-attrs
    type: templ.Attributes
  Order:
    role: heading-level
    type: int
    applies-to: CardTitle
showcase:
  - id: variant.default
    props: { Variant: default }
  - id: variant.raised
    props: { Variant: raised }
  - id: variant.kpi
    props: { Variant: kpi }
  - id: variant.muted
    props: { Variant: muted }
  - id: variant.ghost
    props: { Variant: ghost }
  - id: variant.compact
    props: { Variant: compact }
  - id: variant.flat
    props: { Variant: flat }
  - id: variant.accent
    props: { Variant: accent }
  - id: layout.header-only
    parts: [Card, CardHeader, CardTitle]
  - id: title.order-h3
    props: { Variant: default }
    parts: [Card, CardHeader, CardTitle]
semantics:
  root: div | section | article | aside | header | footer | main | nav | figure | search | hgroup
  role: none
  behavior: static
  title-root: h1 | h2 | h3
  description-root: p

---
## Summary

Card groups related content in one bordered surface.
Card uses header, content, and footer slots.

## Use Cases

- Show a form inside a panel
- Show KPI metrics in a dashboard block
- Show article preview with title and actions

## Semantics

- Card root element is resolved from Tag
- CardTitle maps Order to h1, h2, or h3
- CardDescription renders a p element
- Card parts compose in document order

## Example variant.default

```templ
import "github.com/fastygo/templ/ui"
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Card(cmp.CardProps{Variant: "default"}) {
		@cmp.CardHeader(cmp.CardHeaderProps{}) {
			@cmp.CardTitle(cmp.CardTitleProps{Order: 2}, "Card title")
			@cmp.CardDescription(cmp.CardDescriptionProps{}, "Short supporting text.")
		}
		@cmp.CardContent(cmp.CardContentProps{}) {
			@ui.Text(ui.TextProps{}, "Main card body.")
		}
		@cmp.CardFooter(cmp.CardFooterProps{}) {
			@ui.Button(ui.ButtonProps{Variant: "outline", Size: "sm"}) {
				Action
			}
		}
	}
}
```

## Example variant.raised

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Card(cmp.CardProps{Variant: "raised"}) {
		@cmp.CardContent(cmp.CardContentProps{}) {
			Raised surface
		}
	}
}
```

## Example variant.kpi

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Card(cmp.CardProps{Variant: "kpi"}) {
		@cmp.CardHeader(cmp.CardHeaderProps{}) {
			@cmp.CardTitle(cmp.CardTitleProps{Order: 2}, "Revenue")
		}
		@cmp.CardContent(cmp.CardContentProps{}) {
			$12,450
		}
	}
}
```

## Example variant.muted

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Card(cmp.CardProps{Variant: "muted"}) {
		@cmp.CardContent(cmp.CardContentProps{}) {
			Muted surface
		}
	}
}
```

## Example variant.ghost

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Card(cmp.CardProps{Variant: "ghost"}) {
		@cmp.CardContent(cmp.CardContentProps{}) {
			Ghost dashed border
		}
	}
}
```

## Example variant.compact

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Card(cmp.CardProps{Variant: "compact"}) {
		@cmp.CardContent(cmp.CardContentProps{}) {
			Compact padding
		}
	}
}
```

## Example variant.flat

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Card(cmp.CardProps{Variant: "flat"}) {
		@cmp.CardContent(cmp.CardContentProps{}) {
			Flat no shadow
		}
	}
}
```

## Example variant.accent

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Card(cmp.CardProps{Variant: "accent"}) {
		@cmp.CardContent(cmp.CardContentProps{}) {
			Accent highlight
		}
	}
}
```

## Example layout.header-only

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Card(cmp.CardProps{}) {
		@cmp.CardHeader(cmp.CardHeaderProps{}) {
			@cmp.CardTitle(cmp.CardTitleProps{Order: 2}, "Title only")
		}
	}
}
```

## Example title.order-h3

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Card(cmp.CardProps{}) {
		@cmp.CardHeader(cmp.CardHeaderProps{}) {
			@cmp.CardTitle(cmp.CardTitleProps{Order: 3}, "Nested section title")
		}
	}
}
```
