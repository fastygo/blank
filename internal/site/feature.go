package site

import (
	"net/http"

	"github.com/fastygo/framework/pkg/app"
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

// Routes implements app.Feature.
func (f *Feature) Routes(mux *http.ServeMux) {
	for _, page := range pages {
		page := page
		mux.HandleFunc(page.Method+" "+page.Pattern, f.handlePage(page))
	}
}
