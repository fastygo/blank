package views

import (
	"github.com/fastygo/blank/internal/ui/layout"
	"github.com/fastygo/blank/internal/views/partials"
)

func shellProps(d LayoutData) layout.ShellProps {
	return layout.ShellProps{
		Title:          d.DocumentTitle(),
		Lang:           d.Lang,
		BrandName:      d.Brand,
		Active:         d.Active,
		NavItems:       d.NavItems,
		FooterText:     d.FooterText,
		Navigation:     d.Navigation,
		HeadExtra:      partials.ShellHead(d.Assets.CSS, d.Assets.ThemeJS, d.Assets.AppJS),
		HeaderTrailing: partials.HeaderTrailing(d.LanguageSwitch),
		ThemeToggle:    d.Theme,
	}
}
