package layout

import (
	"net/http"

	"github.com/fastygo/blank/internal/fixtures"
)

// BuildParams collects request-scoped inputs for layout.Data assembly.
type BuildParams struct {
	PageTitle        string
	Active           string
	Lang             string
	Brand            string
	FooterText       string
	Fix              fixtures.Locale
	NavItems         []NavItem
	Request          *http.Request
	AvailableLocales []string
	DefaultLocale    string
}

// BuildData assembles layout.Data from resolved fixtures and request context.
func BuildData(p BuildParams) Data {
	return Data{
		PageTitle:  p.PageTitle,
		Lang:       p.Lang,
		Brand:      p.Brand,
		FooterText: p.FooterText,
		Active:     p.Active,
		NavItems:   p.NavItems,
		Navigation: navigationProps(p.Fix),
		Assets: AssetPaths{
			CSS:     assetCSS,
			ThemeJS: assetThemeJS,
			AppJS:   assetAppJS,
		},
		Theme: ThemeToggleProps{
			Label:              p.Fix.Theme.Label,
			SwitchToDarkLabel:  p.Fix.Theme.SwitchToDarkLabel,
			SwitchToLightLabel: p.Fix.Theme.SwitchToLight,
		},
		LanguageSwitch: languageSwitch(p.Request.Context(), p.Request, p.Fix, p.AvailableLocales, p.DefaultLocale),
	}
}
