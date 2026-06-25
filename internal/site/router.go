package site

import (
	"github.com/a-h/templ"
	"github.com/fastygo/blank/internal/fixtures"
	"github.com/fastygo/blank/internal/ui/layout"
	"github.com/fastygo/blank/internal/views"
)

type TitleResolver func(fixtures.Locale) string
type PageRenderer func(layout.Data, fixtures.Locale) templ.Component
type NavResolver func(fixtures.Locale) (layout.NavItem, bool)

// PageSpec is one runtime route: method, pattern, page renderer, and optional nav entry.
// The Body renderer returns a fully composed HTML document — the page template
// itself chooses its layout layers (RootLayout + TopnavLayout or DashboardLayout), so the
// runtime router does not own layout selection.
type PageSpec struct {
	Method  string
	Pattern string
	Active  string
	Title   TitleResolver
	Body    PageRenderer
	Nav     NavResolver
}

var pages = []PageSpec{
	{
		Method:  "GET",
		Pattern: "/{$}",
		Active:  "/",
		Title:   func(f fixtures.Locale) string { return f.Home.Title },
		Body:    views.HomePage,
		Nav: func(f fixtures.Locale) (layout.NavItem, bool) {
			return layout.NavItem{Label: f.Home.NavLabel, Path: "/", Icon: "home"}, true
		},
	},
	{
		Method:  "GET",
		Pattern: "/sample",
		Active:  "/sample",
		Title:   func(f fixtures.Locale) string { return f.Sample.Title },
		Body:    views.SamplePage,
		Nav: func(f fixtures.Locale) (layout.NavItem, bool) {
			return layout.NavItem{Label: f.Sample.NavLabel, Path: "/sample", Icon: "box"}, true
		},
	},
}
