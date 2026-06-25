package layout

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/a-h/templ"
	"github.com/fastygo/blank/internal/ui/components/toggles"
)

func minimalLayoutData() Data {
	return Data{
		PageTitle:  "Home",
		Lang:       "en",
		Brand:      "FastyGo",
		FooterText: "Made with FastyGo",
		Active:     "/",
		NavItems: []NavItem{
			{Label: "Home", Path: "/", Icon: "home"},
		},
		Assets: AssetPaths{
			CSS:     "/static/css/app.css",
			ThemeJS: "/static/js/theme.js",
			AppJS:   "/static/js/ui8kit.js",
		},
		Navigation: NavigationProps{
			BrandHomeSuffix:     " home",
			MainNavigation:      "Main navigation",
			OpenNavigationMenu:  "Open navigation menu",
			CloseNavigationMenu: "Close navigation menu",
			NavigationMenuLabel: "Navigation menu",
		},
		Theme: ThemeToggleProps{
			Label:              "Theme",
			SwitchToDarkLabel:  "Switch to dark",
			SwitchToLightLabel: "Switch to light",
		},
		LanguageSwitch: toggles.LanguageSwitchProps{
			AriaLabel: "Language",
			Items: []toggles.LanguageSwitchItem{
				{Locale: "en", Label: "En", Href: "/?lang=en", Active: true},
			},
		},
	}
}

func renderComponent(t *testing.T, c templ.Component) string {
	t.Helper()
	var buf bytes.Buffer
	if err := c.Render(context.Background(), io.Writer(&buf)); err != nil {
		t.Fatal(err)
	}
	return buf.String()
}

func TestRootLayout_rendersDocumentFrame(t *testing.T) {
	d := minimalLayoutData()
	html := renderComponent(t, RootLayout(d.Document()))

	if !strings.Contains(strings.ToLower(html), "<!doctype html>") {
		t.Fatal("expected doctype")
	}
	if !strings.Contains(html, `<html lang="en">`) {
		t.Fatal("expected html lang")
	}
	if !strings.Contains(html, "<title>Home · FastyGo</title>") {
		t.Fatal("expected document title")
	}
	if !strings.Contains(html, `href="/static/css/app.css"`) {
		t.Fatal("expected css asset link")
	}
	if strings.Contains(html, "<header") {
		t.Fatal("RootLayout should not render header chrome")
	}
}

func TestTopnavLayout_rendersHeaderAndMobileSheet(t *testing.T) {
	d := minimalLayoutData()
	html := renderComponent(t, TopnavLayout(d.Topnav()))

	if !strings.Contains(html, "<header") {
		t.Fatal("expected header")
	}
	if !strings.Contains(html, "<footer") || !strings.Contains(html, "Made with FastyGo") {
		t.Fatal("expected footer")
	}
	if !strings.Contains(html, `data-ui8kit="sheet"`) {
		t.Fatal("expected mobile sheet markup")
	}
	if !strings.Contains(html, `id="ui8kit-mobile-sheet-trigger"`) {
		t.Fatal("expected mobile sheet trigger")
	}
	if !strings.Contains(html, `id="ui8kit-theme-toggle"`) {
		t.Fatal("expected theme toggle")
	}
}

func TestDashboardLayout_rendersAside(t *testing.T) {
	d := minimalLayoutData()
	d.Active = "/sample"
	d.NavItems = append(d.NavItems, NavItem{Label: "Sample", Path: "/sample", Icon: "box"})
	html := renderComponent(t, DashboardLayout(d.Dashboard("Sample page")))

	if !strings.Contains(html, "<aside") {
		t.Fatal("expected desktop sidebar aside")
	}
	if !strings.Contains(html, `aria-label="Main navigation"`) {
		t.Fatal("expected sidebar aria label")
	}
	if !strings.Contains(html, "<header") {
		t.Fatal("expected inherited topnav header")
	}
	if !strings.Contains(html, `data-ui8kit="sheet"`) {
		t.Fatal("expected inherited mobile sheet")
	}
}
