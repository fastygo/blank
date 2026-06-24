package views

import (
	"github.com/fastygo/blank/internal/ui/components/appsidebar"
	"github.com/fastygo/blank/internal/ui/components/toggles"
	"github.com/fastygo/blank/internal/ui/layout"
)

// AssetPaths are URLs for CSS and JS bundles.
type AssetPaths struct {
	CSS     string
	ThemeJS string
	AppJS   string
}

// LayoutData drives the document shell (assets, locale, nav, theme, language).
// Site/render builds it once per request and passes it into each page template.
type LayoutData struct {
	PageTitle      string
	Lang           string
	Brand          string
	FooterText     string
	Active         string
	NavItems       []layout.NavItem
	Navigation     layout.NavigationProps
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

// HomePageData is the home page payload (document shell + hero content).
type HomePageData struct {
	Shell        layout.ShellProps
	Welcome      string
	WelcomeBrand string
	Description  string
}

// SamplePageData is the sample page payload (document shell + sidebar + content).
type SamplePageData struct {
	Shell       layout.ShellProps
	Sidebar     appsidebar.Props
	Title       string
	Description string
	Body        string
}
