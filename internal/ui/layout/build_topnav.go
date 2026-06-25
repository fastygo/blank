package layout

import (
	"context"
	"net/http"
	"strings"

	"github.com/fastygo/blank/internal/fixtures"
	"github.com/fastygo/blank/internal/ui/components/toggles"
	"github.com/fastygo/framework/pkg/web/locale"
)

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
