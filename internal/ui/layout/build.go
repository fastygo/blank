package layout

import (
	"context"
	"net/http"
	"strings"

	"github.com/fastygo/blank/internal/fixtures"
	"github.com/fastygo/blank/internal/ui/components/toggles"
	"github.com/fastygo/framework/pkg/web/locale"
)

const (
	assetCSS     = "/static/css/app.css"
	assetThemeJS = "/static/js/theme.js"
	assetAppJS   = "/static/js/ui8kit.js"
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

func navigationProps(fix fixtures.Locale) NavigationProps {
	return NavigationProps{
		BrandHomeSuffix:     fix.Layout.BrandHomeSuffix,
		MainNavigation:      fix.Layout.MainNavigation,
		OpenNavigationMenu:  fix.Layout.OpenNavigationMenu,
		CloseNavigationMenu: fix.Layout.CloseNavigationMenu,
		NavigationMenuLabel: fix.Layout.NavigationMenuLabel,
	}
}

func languageSwitch(ctx context.Context, r *http.Request, fix fixtures.Locale, available []string, defaultLocale string) toggles.LanguageSwitchProps {
	current := strings.ToLower(strings.TrimSpace(locale.From(ctx)))
	if current == "" {
		current = strings.ToLower(strings.TrimSpace(defaultLocale))
	}
	var items []toggles.LanguageSwitchItem
	for _, loc := range available {
		loc = strings.ToLower(strings.TrimSpace(loc))
		label := fix.Language[loc]
		if label == "" {
			label = strings.ToUpper(loc)
		}
		items = append(items, toggles.LanguageSwitchItem{
			Locale: loc,
			Label:  label,
			Href:   locale.BuildLangHref(r, loc, defaultLocale),
			Active: loc == current,
		})
	}
	return toggles.LanguageSwitchProps{
		AriaLabel: fix.LanguageToggleLabel,
		Items:     items,
	}
}
