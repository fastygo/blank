---
id: ui.text
layer: primitive
kind: typography
package: github.com/fastygo/templ/ui/text
facade: github.com/fastygo/templ/ui
templ: Text
api:
  Tag:
    role: prose-tag
    type: string
    enum: [p, blockquote, figcaption, address, pre, span, em, strong, small, abbr, cite, code, kbd, mark, time, data, var, samp, sub, sup, b, i, u, s, q, dfn, bdo, bdi, ins, del]
    allow-list-source: utils.tags.TagGroupText
    default: p
  Class:
    role: style-extension
    type: string
  value:
    role: copy
    type: string
    position: argument
  Attrs:
    role: html-attrs
    type: templ.Attributes
showcase:
  - id: tag.p
    props: { Tag: p }
  - id: tag.small
    props: { Tag: small, Class: "text-muted-foreground" }
  - id: tag.code
    props: { Tag: code, Class: "text-sm" }
semantics:
  root: resolved from Tag
  role: none
  behavior: static

---
## Summary

Text renders plain copy in a semantic tag.
Pass copy as the second value argument.

## Use Cases

- Render body paragraph copy
- Render muted helper line with Tag small

## Semantics

- ResolveTag uses TagGroupText from utils
- Invalid Tag falls back to p

## Example tag.p

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Text(ui.TextProps{}, "Body paragraph text.")
}
```

## Example tag.small

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Text(ui.TextProps{Tag: "small", Class: "text-muted-foreground"}, "Helper line.")
}
```

## Example tag.code

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Text(ui.TextProps{Tag: "code", Class: "text-sm"}, "npm install")
}
```
