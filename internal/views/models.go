package views

import (
	"github.com/fastygo/blank/internal/ui/components/toggles"
	"github.com/fastygo/blank/internal/ui/layout"
)

// AssetPaths are URLs for CSS and JS bundles.
type AssetPaths struct {
	CSS     string
	ThemeJS string
	AppJS   string
}

// LayoutData drives the app shell (header, main, footer).
type LayoutData struct {
	PageTitle      string
	Lang           string
	Brand          string
	FooterText     string
	Active         string
	NavItems       []layout.NavItem
	Assets         AssetPaths
	Theme          layout.ThemeToggleProps
	LanguageSwitch toggles.LanguageSwitchProps
}

// DocumentTitle returns the SEO document title for <title>.
func (d LayoutData) DocumentTitle() string {
	return FormatDocumentTitle(d.PageTitle, d.Brand)
}

// FormatDocumentTitle builds "Page · Brand" for the document head.
func FormatDocumentTitle(pageTitle, brand string) string {
	if brand == "" {
		brand = "FastyGo"
	}
	return pageTitle + " · " + brand
}

// HomeData is the home page hero body inside the shell.
type HomeData struct {
	Welcome      string
	WelcomeBrand string
	Description  string
}

// SampleData is a second stub route for onboarding new pages.
type SampleData struct {
	Title       string
	Description string
	Body        string
}
