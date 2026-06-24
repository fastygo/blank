package site

import (
	"github.com/a-h/templ"
	"github.com/fastygo/blank/internal/fixtures"
	"github.com/fastygo/blank/internal/ui/layout"
	"github.com/fastygo/blank/internal/views"
)

type TitleResolver func(fixtures.Locale) string
type PageRenderer func(views.LayoutData, fixtures.Locale) templ.Component
type NavResolver func(fixtures.Locale) (layout.NavItem, bool)

// PageSpec is one runtime route: method, pattern, page renderer, and optional nav entry.
// The Body renderer returns a fully composed HTML document — the page template
// itself chooses its layout shell (layout.Shell or layout.SidebarShell), so the
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
		Body: func(d views.LayoutData, f fixtures.Locale) templ.Component {
			return views.HomePage(views.HomePageData{
				Shell:        views.ShellPropsFor(d),
				Welcome:      f.Home.Welcome,
				WelcomeBrand: f.Home.WelcomeBrand,
				Description:  f.Home.Description,
			})
		},
		Nav: func(f fixtures.Locale) (layout.NavItem, bool) {
			return layout.NavItem{Label: f.Home.NavLabel, Path: "/", Icon: "home"}, true
		},
	},
	{
		Method:  "GET",
		Pattern: "/sample",
		Active:  "/sample",
		Title:   func(f fixtures.Locale) string { return f.Sample.Title },
		Body: func(d views.LayoutData, f fixtures.Locale) templ.Component {
			return views.SamplePage(views.SamplePageData{
				Shell:       views.ShellPropsFor(d),
				Sidebar:     views.SidebarPropsFor(d, f.Sample.Title),
				Title:       f.Sample.Title,
				Description: f.Sample.Description,
				Body:        f.Sample.Body,
			})
		},
		Nav: func(f fixtures.Locale) (layout.NavItem, bool) {
			return layout.NavItem{Label: f.Sample.NavLabel, Path: "/sample", Icon: "box"}, true
		},
	},
}
