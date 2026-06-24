package site

import (
	"context"
	"net/http"
	"strings"

	"github.com/fastygo/blank/internal/fixtures"
	"github.com/fastygo/blank/internal/ui/components/toggles"
	"github.com/fastygo/blank/internal/ui/layout"
	"github.com/fastygo/blank/internal/views"
	"github.com/fastygo/framework/pkg/web/locale"
)

const (
	assetCSS     = "/static/css/app.css"
	assetThemeJS = "/static/js/theme.js"
	assetAppJS   = "/static/js/ui8kit.js"
)

func (f *Feature) navigationProps(fix fixtures.Locale) layout.NavigationProps {
	return layout.NavigationProps{
		BrandHomeSuffix:     fix.Layout.BrandHomeSuffix,
		MainNavigation:      fix.Layout.MainNavigation,
		OpenNavigationMenu:  fix.Layout.OpenNavigationMenu,
		CloseNavigationMenu: fix.Layout.CloseNavigationMenu,
		NavigationMenuLabel: fix.Layout.NavigationMenuLabel,
	}
}

func (f *Feature) languageSwitch(ctx context.Context, r *http.Request, fix fixtures.Locale) toggles.LanguageSwitchProps {
	current := strings.ToLower(strings.TrimSpace(locale.From(ctx)))
	if current == "" {
		current = strings.ToLower(strings.TrimSpace(f.defaultLocale))
	}
	var items []toggles.LanguageSwitchItem
	for _, loc := range f.available {
		loc = strings.ToLower(strings.TrimSpace(loc))
		label := fix.Language[loc]
		if label == "" {
			label = strings.ToUpper(loc)
		}
		items = append(items, toggles.LanguageSwitchItem{
			Locale: loc,
			Label:  label,
			Href:   locale.BuildLangHref(r, loc, f.defaultLocale),
			Active: loc == current,
		})
	}
	return toggles.LanguageSwitchProps{
		AriaLabel: fix.LanguageToggleLabel,
		Items:     items,
	}
}

func (f *Feature) layoutData(ctx context.Context, r *http.Request, fix fixtures.Locale, title, active string) views.LayoutData {
	return views.LayoutData{
		PageTitle:  title,
		Lang:       locale.From(ctx),
		Brand:      fix.Brand,
		FooterText: fix.Footer,
		Active:     active,
		NavItems:   f.siteNav(fix),
		Navigation: f.navigationProps(fix),
		Assets: views.AssetPaths{
			CSS:     assetCSS,
			ThemeJS: assetThemeJS,
			AppJS:   assetAppJS,
		},
		Theme: layout.ThemeToggleProps{
			Label:              fix.Theme.Label,
			SwitchToDarkLabel:  fix.Theme.SwitchToDarkLabel,
			SwitchToLightLabel: fix.Theme.SwitchToLight,
		},
		LanguageSwitch: f.languageSwitch(ctx, r, fix),
	}
}
