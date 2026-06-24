package views

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/fastygo/blank/internal/ui/components/toggles"
	"github.com/fastygo/blank/internal/ui/layout"
)

func homeLayoutData() LayoutData {
	return LayoutData{
		PageTitle:  "Home",
		Lang:       "en",
		Brand:      "FastyGo",
		FooterText: "Made with 🤍 in FastyGo",
		Active:     "/",
		NavItems: []layout.NavItem{
			{Label: "Home", Path: "/", Icon: "home"},
		},
		Assets: AssetPaths{
			CSS:     "/static/css/app.css",
			ThemeJS: "/static/js/theme.js",
			AppJS:   "/static/js/ui8kit.js",
		},
		Navigation: layout.NavigationProps{
			BrandHomeSuffix:     " home",
			MainNavigation:      "Main navigation",
			OpenNavigationMenu:  "Open navigation menu",
			CloseNavigationMenu: "Close navigation menu",
			NavigationMenuLabel: "Navigation menu",
		},
		Theme: layout.ThemeToggleProps{
			Label:              "Theme",
			SwitchToDarkLabel:  "Switch to dark",
			SwitchToLightLabel: "Switch to light",
		},
		LanguageSwitch: toggles.LanguageSwitchProps{
			AriaLabel: "Language",
			Items: []toggles.LanguageSwitchItem{
				{Locale: "en", Label: "En", Href: "/?lang=en", Active: true},
				{Locale: "ru", Label: "Ru", Href: "/?lang=ru"},
			},
		},
	}
}

func renderHomePage(t *testing.T, d LayoutData) string {
	t.Helper()
	body := HomePage(HomePageData{
		Shell:        ShellPropsFor(d),
		Welcome:      "Welcome",
		WelcomeBrand: "to FastyGo",
		Description:  "Minimal starter.",
	})
	var buf bytes.Buffer
	if err := body.Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	return buf.String()
}

func TestHomePage_composesTopnavShell(t *testing.T) {
	html := renderHomePage(t, homeLayoutData())

	if !strings.Contains(strings.ToLower(html), "<!doctype html>") {
		t.Fatal("expected full document with doctype")
	}
	if strings.Contains(html, "<aside") {
		t.Fatal("expected no sidebar aside on topnav page")
	}
	if !strings.Contains(html, `data-ui8kit="sheet"`) {
		t.Fatal("expected shell mobile sheet markup")
	}
	if !strings.Contains(html, `id="ui8kit-theme-toggle"`) {
		t.Fatal("expected theme toggle control")
	}
	if !strings.Contains(html, `id="ui8kit-mobile-sheet-trigger"`) {
		t.Fatal("expected mobile menu trigger")
	}
	if !strings.Contains(html, "Made with 🤍 in FastyGo") {
		t.Fatal("expected footer copy")
	}
	if !strings.Contains(html, "<title>Home · FastyGo</title>") {
		t.Fatal("expected brand in document title")
	}
	if !strings.Contains(html, "Welcome") || !strings.Contains(html, "to FastyGo") {
		t.Fatal("expected hero welcome copy")
	}
	if !strings.Contains(html, `role="group"`) || !strings.Contains(html, "En") {
		t.Fatal("expected language switch control")
	}
}

func sampleLayoutData() LayoutData {
	d := homeLayoutData()
	d.PageTitle = "Sample"
	d.Active = "/sample"
	d.NavItems = []layout.NavItem{
		{Label: "Home", Path: "/", Icon: "home"},
		{Label: "Sample", Path: "/sample", Icon: "box"},
	}
	return d
}

func renderSamplePage(t *testing.T, d LayoutData) string {
	t.Helper()
	body := SamplePage(SamplePageData{
		Shell:       ShellPropsFor(d),
		Sidebar:     SidebarPropsFor(d, "Sample page"),
		Title:       "Sample page",
		Description: "Second route for onboarding.",
		Body:        "Add routes in router.go.",
	})
	var buf bytes.Buffer
	if err := body.Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	return buf.String()
}

func TestSamplePage_composesSidebarShell(t *testing.T) {
	html := renderSamplePage(t, sampleLayoutData())

	if !strings.Contains(strings.ToLower(html), "<!doctype html>") {
		t.Fatal("expected full document with doctype")
	}
	if !strings.Contains(html, "<title>Sample · FastyGo</title>") {
		t.Fatal("expected document title")
	}
	if !strings.Contains(html, "<aside") {
		t.Fatal("expected desktop sidebar aside")
	}
	if !strings.Contains(html, `aria-label="Main navigation"`) {
		t.Fatal("expected sidebar aria label from shell navigation props")
	}
	if !strings.Contains(html, `aria-current="page"`) {
		t.Fatal("expected active nav item aria-current")
	}
	if !strings.Contains(html, `data-ui8kit="sheet"`) {
		t.Fatal("expected mobile sheet markup")
	}
	if !strings.Contains(html, `id="ui8kit-mobile-sheet-panel"`) {
		t.Fatal("expected mobile sheet panel id")
	}
	if !strings.Contains(html, `id="ui8kit-mobile-sheet-trigger"`) {
		t.Fatal("expected mobile sheet trigger id")
	}
	if !strings.Contains(html, "Sample page") {
		t.Fatal("expected page body slot copy")
	}
}
