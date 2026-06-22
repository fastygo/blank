package site

import (
	"context"
	"net/http"
	"strings"

	"github.com/fastygo/blank/internal/fixtures"
	"github.com/fastygo/blank/internal/ui/components/toggles"
	"github.com/fastygo/blank/internal/ui/layout"
	"github.com/fastygo/blank/internal/views"
	"github.com/fastygo/framework/pkg/app"
	"github.com/fastygo/framework/pkg/web"
	"github.com/fastygo/framework/pkg/web/locale"
)

const (
	assetCSS     = "/static/css/app.css"
	assetThemeJS = "/static/js/theme.js"
	assetAppJS   = "/static/js/ui8kit.js"
)

// Feature wires public HTTP routes for the app shell (header nav, i18n, theme).
type Feature struct {
	available     []string
	defaultLocale string
}

// NewFeature constructs the site feature.
func NewFeature(available []string, defaultLocale string) *Feature {
	return &Feature{
		available:     available,
		defaultLocale: defaultLocale,
	}
}

// ID implements app.Feature.
func (f *Feature) ID() string {
	return "site"
}

// NavItems implements app.Feature.
// Navigation labels are resolved per request from fixtures (see siteNav).
func (f *Feature) NavItems() []app.NavItem {
	return nil
}

func (f *Feature) fixtureLocale(ctx context.Context) fixtures.Locale {
	return fixtures.Resolve(locale.From(ctx), f.defaultLocale)
}

func (f *Feature) siteNav(fix fixtures.Locale) []layout.NavItem {
	return []layout.NavItem{
		{Label: fix.Home.NavLabel, Path: "/", Icon: "home"},
		{Label: fix.Sample.NavLabel, Path: "/sample", Icon: "box"},
	}
}

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

// Routes implements app.Feature.
func (f *Feature) Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", f.getHome)
	mux.HandleFunc("GET /sample", f.getSample)
}

func (f *Feature) getHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fix := f.fixtureLocale(ctx)
	data := f.layoutData(ctx, r, fix, fix.Home.Title, "/")
	_ = web.Render(ctx, w, views.SiteShell(data, views.HomePage(views.HomeData{
		Welcome:      fix.Home.Welcome,
		WelcomeBrand: fix.Home.WelcomeBrand,
		Description:  fix.Home.Description,
	})))
}

func (f *Feature) getSample(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fix := f.fixtureLocale(ctx)
	data := f.layoutData(ctx, r, fix, fix.Sample.Title, "/sample")
	_ = web.Render(ctx, w, views.SiteShell(data, views.SamplePage(views.SampleData{
		Title:       fix.Sample.Title,
		Description: fix.Sample.Description,
		Body:        fix.Sample.Body,
	})))
}
