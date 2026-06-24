package site

import (
	"github.com/a-h/templ"
	"github.com/fastygo/blank/internal/fixtures"
	"github.com/fastygo/blank/internal/ui/layout"
	"github.com/fastygo/blank/internal/views"
)

type LayoutRenderer func(views.LayoutData, templ.Component) templ.Component
type TitleResolver func(fixtures.Locale) string
type BodyRenderer func(fixtures.Locale) templ.Component
type NavResolver func(fixtures.Locale) (layout.NavItem, bool)

// PageSpec is one runtime route: method, pattern, layout adapter, and page body.
type PageSpec struct {
	Method  string
	Pattern string
	Active  string
	Layout  LayoutRenderer
	Title   TitleResolver
	Body    BodyRenderer
	Nav     NavResolver
}

var pages = []PageSpec{
	{
		Method:  "GET",
		Pattern: "/{$}",
		Active:  "/",
		Layout:  views.AppShell,
		Title:   func(f fixtures.Locale) string { return f.Home.Title },
		Body: func(f fixtures.Locale) templ.Component {
			return views.HomePage(views.HomeData{
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
		Layout:  views.SidebarAppShell,
		Title:   func(f fixtures.Locale) string { return f.Sample.Title },
		Body: func(f fixtures.Locale) templ.Component {
			return views.SamplePage(views.SampleData{
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
