package views

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/fastygo/blank/internal/ui/components/toggles"
	"github.com/fastygo/blank/internal/ui/layout"
)

func TestSiteShell_homeRenders(t *testing.T) {
	d := LayoutData{
		PageTitle: "Home",
		Lang:      "en",
		Brand:     "Blank",
		Active:    "/",
		NavItems: []layout.NavItem{
			{Label: "Home", Path: "/", Icon: "home"},
		},
		Assets: AssetPaths{
			CSS:     "/static/css/app.css",
			ThemeJS: "/static/js/theme.js",
			AppJS:   "/static/js/ui8kit.js",
		},
		Theme: layout.ThemeToggleProps{},
		LanguageSwitch: toggles.LanguageSwitchProps{
			AriaLabel: "Language",
			Items: []toggles.LanguageSwitchItem{
				{Locale: "en", Label: "En", Href: "/?lang=en", Active: true},
				{Locale: "ru", Label: "Ru", Href: "/?lang=ru"},
			},
		},
	}
	body := HomePage(HomeData{Title: "Home", Description: "d", Body: "b"})
	var buf bytes.Buffer
	if err := SiteShell(d, body).Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	html := buf.String()
	if !strings.Contains(strings.ToLower(html), "<!doctype html>") {
		t.Fatal("expected full document with doctype")
	}
	if !strings.Contains(html, `data-ui8kit="sheet"`) {
		t.Fatal("expected shell mobile sheet markup")
	}
	if !strings.Contains(html, "<aside") {
		t.Fatal("expected desktop sidebar aside landmark")
	}
	if !strings.Contains(html, `id="ui8kit-theme-toggle"`) {
		t.Fatal("expected theme toggle control")
	}
	if !strings.Contains(html, "<title>Home · Blank</title>") {
		t.Fatal("expected brand in document title")
	}
	if !strings.Contains(html, `role="group"`) || !strings.Contains(html, "En") {
		t.Fatal("expected language switch control")
	}
}
