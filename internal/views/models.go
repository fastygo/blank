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

// LayoutData drives the cabinet shell (sidebar + header).
type LayoutData struct {
	PageTitle      string
	Lang           string
	Brand          string
	Active         string
	NavItems       []layout.NavItem
	Assets         AssetPaths
	Theme          layout.ThemeToggleProps
	LanguageSwitch toggles.LanguageSwitchProps
	AccountEmail   string
	AccountSignOut string
}

// DocumentTitle returns the SEO document title for <title>.
func (d LayoutData) DocumentTitle() string {
	return FormatDocumentTitle(d.PageTitle, d.Brand)
}

// FormatDocumentTitle builds "Page · Brand" for the document head.
func FormatDocumentTitle(pageTitle, brand string) string {
	if brand == "" {
		brand = "Blank Panel"
	}
	return pageTitle + " · " + brand
}

// LoginPageData is the sign-in screen.
type LoginPageData struct {
	Title          string
	Lang           string
	Brand          string
	Subtitle       string
	Error          string
	ReturnTo       string
	Assets         AssetPaths
	EmailLabel     string
	PasswordLabel  string
	SubmitLabel    string
	Theme          layout.ThemeToggleProps
	LanguageSwitch toggles.LanguageSwitchProps
}

// DashboardData is the cabinet home page body.
type DashboardData struct {
	Title       string
	Description string
	Body        string
}

// SampleData is a second stub route for copy-paste onboarding.
type SampleData struct {
	Title       string
	Description string
	Body        string
}
