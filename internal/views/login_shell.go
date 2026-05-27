package views

import (
	"github.com/fastygo/blank/internal/ui/layout"
	"github.com/fastygo/blank/internal/views/partials"
)

// LoginShellProps builds shell props for the sign-in screen (no sidebar).
func LoginShellProps(d LoginPageData) layout.ShellProps {
	return layout.ShellProps{
		Title:          FormatDocumentTitle(d.Title, d.Brand),
		Lang:           d.Lang,
		BrandName:      d.Brand,
		HeadExtra:      partials.ShellHead(d.Assets.CSS, d.Assets.ThemeJS, d.Assets.AppJS),
		HeaderExtra:    partials.LoginHeaderExtra(d.LanguageSwitch),
		ThemeToggle:    d.Theme,
		MarketingShell: true,
	}
}
