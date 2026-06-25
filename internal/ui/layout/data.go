// Package layout provides document shells and request-scoped layout data.
// Two-layer model: RootLayout owns the document; TopnavLayout/DashboardLayout own app chrome.
// Pages compose both.
package layout

import (
	"github.com/fastygo/blank/internal/ui/components/toggles"
)

// AssetPaths are URLs for CSS and JS bundles.
type AssetPaths struct {
	CSS     string
	ThemeJS string
	AppJS   string
}

// Data drives the document shell (assets, locale, nav, theme, language).
// Site/render builds it once per request and passes it into each page template.
type Data struct {
	PageTitle      string
	Lang           string
	Brand          string
	FooterText     string
	Active         string
	NavItems       []NavItem
	Navigation     NavigationProps
	Assets         AssetPaths
	Theme          ThemeToggleProps
	LanguageSwitch toggles.LanguageSwitchProps
}

// DocumentTitle returns the SEO document title for <title>.
func (d Data) DocumentTitle() string {
	return FormatDocumentTitle(d.PageTitle, d.Brand)
}

// FormatDocumentTitle builds "Page · Brand" for the document head.
func FormatDocumentTitle(pageTitle, brand string) string {
	if brand == "" {
		brand = "FastyGo"
	}
	return pageTitle + " · " + brand
}

// Document maps request data into document-frame props for RootLayout.
func (d Data) Document() DocumentProps {
	return DocumentProps{
		Title:     d.DocumentTitle(),
		Lang:      d.Lang,
		BodyClass: documentBodyClass(),
		HeadExtra: ShellHead(d.Assets.CSS, d.Assets.ThemeJS, d.Assets.AppJS),
	}
}

// Topnav maps request data into topnav chrome props for TopnavLayout.
func (d Data) Topnav() TopnavLayoutProps {
	return TopnavLayoutProps{
		Brand:           d.Brand,
		Active:          d.Active,
		NavItems:        d.NavItems,
		Navigation:      d.Navigation,
		Theme:           d.Theme,
		Trailing:        HeaderTrailing(d.LanguageSwitch),
		FooterText:      d.FooterText,
		ShowMobileSheet: len(d.NavItems) > 0,
	}
}

// Dashboard maps request data into dashboard chrome props for DashboardLayout.
func (d Data) Dashboard(title string) DashboardLayoutProps {
	return DashboardLayoutProps{
		Topnav:  d.Topnav(),
		Sidebar: sidebarProps(d, title),
	}
}
