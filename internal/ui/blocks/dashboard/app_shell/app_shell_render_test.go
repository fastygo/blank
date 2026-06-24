package appshell

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/fastygo/blank/internal/ui/layout"
	"github.com/fastygo/templ/ui"
)

func minimalShellProps() Props {
	return Props{
		Shell: layout.ShellProps{
			Title:      "Home · FastyGo",
			Lang:       "en",
			BrandName:  "FastyGo",
			Active:     "/",
			FooterText: "Made with 🤍 in FastyGo",
			NavItems: []layout.NavItem{
				{Label: "Home", Path: "/", Icon: "home"},
			},
			Navigation: layout.NavigationProps{
				BrandHomeSuffix:     " home",
				MainNavigation:      "Main navigation",
				OpenNavigationMenu:  "Open navigation menu",
				CloseNavigationMenu: "Close navigation menu",
				NavigationMenuLabel: "Navigation menu",
			},
			ThemeToggle: layout.ThemeToggleProps{
				Label:              "Theme",
				SwitchToDarkLabel:  "Switch to dark",
				SwitchToLightLabel: "Switch to light",
			},
		},
	}
}

func TestAppShell_rendersDocumentShell(t *testing.T) {
	var buf bytes.Buffer
	body := ui.Text(ui.TextProps{}, "Page body")
	if err := AppShell(minimalShellProps(), body).Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	html := buf.String()

	if !strings.Contains(strings.ToLower(html), "<!doctype html>") {
		t.Fatal("expected full document with doctype")
	}
	if !strings.Contains(html, "<title>Home · FastyGo</title>") {
		t.Fatal("expected document title")
	}
	if strings.Contains(html, "<aside") {
		t.Fatal("expected no sidebar aside")
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
	if !strings.Contains(html, "Page body") {
		t.Fatal("expected page body slot")
	}
}
