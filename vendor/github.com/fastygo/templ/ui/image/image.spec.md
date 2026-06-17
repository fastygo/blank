---
id: ui.image
layer: primitive
kind: media
package: github.com/fastygo/templ/ui/image
facade: github.com/fastygo/templ/ui
templ: Image
parts:
  - templ: Image
    props: [Variant, Size, Class, Src, SrcSet, Sizes, Alt, Width, Height, Fit, Position, Aspect, Loading, Decoding, FetchPriority, Decorative, ID, Attrs]
  - templ: Picture
    props: [Class, Attrs]
    slot: responsive-wrapper
  - templ: Source
    props: [SrcSet, Src, Media, Type, Sizes, Attrs]
    slot: source
api:
  Variant:
    role: appearance
    type: string
    enum: [default, rounded, avatar, thumbnail, logo, unstyled]
    allow-list-source: ui.image.ImageVariants
    default: default
  Size:
    role: density
    type: string
    enum: [auto, default, xs, sm, lg, xl, full]
    allow-list-source: ui.image.ImageVariants
    default: auto
  Fit:
    role: object-fit
    type: string
    enum: [contain, cover, fill, none, scale-down]
    default: cover
  Position:
    role: object-position
    type: string
    enum: [bottom, center, left, right, top, left-bottom, left-top, right-bottom, right-top]
    default: center
  Aspect:
    role: aspect-ratio
    type: string
    enum: [auto, square, video]
    default: auto
  Decorative:
    role: accessibility
    type: bool
    default: false
showcase:
  - id: variant.avatar
    props: { Variant: avatar, Size: md, Src: "/avatar.png", Alt: "User avatar" }
  - id: layout.picture
    props: { Variant: rounded, Src: "/hero.jpg", Alt: "Hero image" }
semantics:
  root: img | picture | source
  behavior: static
---

## Summary

Image renders accessible img, picture, and source primitives.
Use Decorative for purely visual images.

## Use Cases

- Render product screenshots, avatars, logos, and thumbnails
- Render responsive image sources with picture

## Semantics

- Decorative images render an empty alt and aria-hidden true
- Image defaults to loading lazy and decoding async

## Example variant.avatar

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Image(ui.ImageProps{Variant: "avatar", Size: "md", Src: "/avatar.png", Alt: "User avatar"})
}
```

## Example layout.picture

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Picture(ui.PictureProps{}) {
		@ui.Source(ui.SourceProps{SrcSet: "/hero.webp", Type: "image/webp"})
		@ui.Image(ui.ImageProps{Variant: "rounded", Src: "/hero.jpg", Alt: "Hero image"})
	}
}
```
