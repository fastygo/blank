package layout

import (
	"strings"

	"github.com/fastygo/blank/internal/ui/components/appsidebar"
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

// ShellProps maps request data into document-shell props consumed by
// Shell and SidebarShell.
func (d Data) ShellProps() ShellProps {
	return ShellProps{
		Title:          d.DocumentTitle(),
		Lang:           d.Lang,
		BrandName:      d.Brand,
		Active:         d.Active,
		NavItems:       d.NavItems,
		FooterText:     d.FooterText,
		Navigation:     d.Navigation,
		HeadExtra:      ShellHead(d.Assets.CSS, d.Assets.ThemeJS, d.Assets.AppJS),
		HeaderTrailing: HeaderTrailing(d.LanguageSwitch),
		ThemeToggle:    d.Theme,
	}
}

// SidebarProps builds appsidebar.Props from request data.
// Pass the title to display above the vertical nav (typically brand or section name);
// when empty, the sidebar falls back to the brand name or "App".
func (d Data) SidebarProps(title string) appsidebar.Props {
	resolvedTitle := strings.TrimSpace(title)
	if resolvedTitle == "" {
		resolvedTitle = strings.TrimSpace(d.Brand)
	}
	if resolvedTitle == "" {
		resolvedTitle = "App"
	}
	return appsidebar.Props{
		Title:     resolvedTitle,
		AriaLabel: d.Navigation.MainNavigation,
		Items:     d.NavItems,
		Active:    d.Active,
	}
}
