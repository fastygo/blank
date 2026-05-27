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

// Feature wires public HTTP routes for the app shell (sidebar, i18n, theme).
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

// SetNavItems implements app.NavProvider.
func (f *Feature) SetNavItems(_ []app.NavItem) {}

// ID implements app.Feature.
func (f *Feature) ID() string {
	return "site"
}

// NavItems implements app.Feature.
func (f *Feature) NavItems() []app.NavItem {
	return nil
}

func (f *Feature) fixtureLocale(ctx context.Context) fixtures.Locale {
	code := locale.From(ctx)
	if code == "" {
		code = f.defaultLocale
	}
	loc, err := fixtures.LoadLocale(code)
	if err != nil {
		loc, _ = fixtures.LoadLocale(f.defaultLocale)
	}
	return loc
}

func (f *Feature) assetPaths() views.AssetPaths {
	return views.AssetPaths{
		CSS:     "/static/css/app.css",
		ThemeJS: "/static/js/theme.js",
		AppJS:   "/static/js/ui8kit.js",
	}
}

func (f *Feature) siteNav(fix fixtures.Locale) []layout.NavItem {
	return []layout.NavItem{
		{Label: fix.Home.NavLabel, Path: "/", Icon: "home"},
		{Label: fix.Sample.NavLabel, Path: "/sample", Icon: "box"},
	}
}

func (f *Feature) languageSwitch(ctx context.Context, r *http.Request, fix fixtures.Locale) toggles.LanguageSwitchProps {
	current := strings.ToLower(strings.TrimSpace(locale.From(ctx)))
	if current == "" {
		current = strings.ToLower(strings.TrimSpace(f.defaultLocale))
	}
	labels := map[string]string{"en": "En", "ru": "Ru"}
	var items []toggles.LanguageSwitchItem
	for _, loc := range f.available {
		loc = strings.ToLower(strings.TrimSpace(loc))
		label := labels[loc]
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

func (f *Feature) layoutData(ctx context.Context, r *http.Request, title, active string) views.LayoutData {
	fix := f.fixtureLocale(ctx)
	return views.LayoutData{
		PageTitle:      title,
		Lang:           locale.From(ctx),
		Brand:          fix.Brand,
		Active:         active,
		NavItems:       f.siteNav(fix),
		Assets:         f.assetPaths(),
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
	layout := f.layoutData(ctx, r, fix.Home.Title, "/")
	_ = web.Render(ctx, w, views.SiteShell(layout, views.HomePage(views.HomeData{
		Title:       fix.Home.Title,
		Description: fix.Home.Description,
		Body:        fix.Home.Body,
	})))
}

func (f *Feature) getSample(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fix := f.fixtureLocale(ctx)
	layout := f.layoutData(ctx, r, fix.Sample.Title, "/sample")
	_ = web.Render(ctx, w, views.SiteShell(layout, views.SamplePage(views.SampleData{
		Title:       fix.Sample.Title,
		Description: fix.Sample.Description,
		Body:        fix.Sample.Body,
	})))
}
