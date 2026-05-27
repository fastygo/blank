package views

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/fastygo/blank/internal/ui/components/toggles"
	"github.com/fastygo/blank/internal/ui/layout"
)

func TestCabinetLayout_dashboardRenders(t *testing.T) {
	d := LayoutData{
		PageTitle: "Dashboard",
		Lang:      "en",
		Brand:     "Blank Panel",
		Active:    "/cabinet",
		NavItems: []layout.NavItem{
			{Label: "Dashboard", Path: "/cabinet", Icon: "home"},
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
				{Locale: "en", Label: "En", Href: "/cabinet?lang=en", Active: true},
				{Locale: "ru", Label: "Ru", Href: "/cabinet?lang=ru"},
			},
		},
		AccountEmail:   "test@admin.dash",
		AccountSignOut: "Sign out",
	}
	body := DashboardPage(DashboardData{Title: "Dashboard", Description: "d", Body: "b"})
	var buf bytes.Buffer
	if err := CabinetLayout(d, body).Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	html := buf.String()
	if !strings.Contains(strings.ToLower(html), "<!doctype html>") {
		t.Fatal("expected full document with doctype")
	}
	if !strings.Contains(html, `data-ui8kit="sheet"`) {
		t.Fatal("expected shell mobile sheet markup")
	}
	if !strings.Contains(html, `id="ui8kit-theme-toggle"`) {
		t.Fatal("expected theme toggle control")
	}
	if !strings.Contains(html, "<title>Dashboard · Blank Panel</title>") {
		t.Fatal("expected brand in document title")
	}
	if !strings.Contains(html, `id="theme-toggle-icon"`) {
		t.Fatal("expected theme toggle icon host")
	}
}

func TestLoginPage_renders(t *testing.T) {
	d := LoginPageData{
		Title:    "Sign in",
		Lang:     "en",
		Brand:    "Blank Panel",
		Subtitle: "Private cabinet",
		Assets: AssetPaths{
			CSS:     "/static/css/app.css",
			ThemeJS: "/static/js/theme.js",
			AppJS:   "/static/js/ui8kit.js",
		},
		EmailLabel:    "Email",
		PasswordLabel: "Password",
		SubmitLabel:   "Sign in",
		Theme:         layout.ThemeToggleProps{},
		LanguageSwitch: toggles.LanguageSwitchProps{
			AriaLabel: "Language",
			Items: []toggles.LanguageSwitchItem{
				{Locale: "en", Label: "En", Href: "/cabinet/login?lang=en", Active: true},
			},
		},
	}
	var buf bytes.Buffer
	if err := LoginPage(d).Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	html := buf.String()
	if !strings.Contains(html, `id="login-email"`) {
		t.Fatal("expected login email field")
	}
	if strings.Contains(html, `data-ui8kit="sheet"`) {
		t.Fatal("login shell should not render sidebar mobile sheet")
	}
}
