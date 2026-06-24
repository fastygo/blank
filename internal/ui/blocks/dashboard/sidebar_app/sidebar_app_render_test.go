package sidebarapp

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/fastygo/templ/ui"
)

func TestSidebarApp_rendersDocumentShell(t *testing.T) {
	var buf bytes.Buffer
	body := ui.Text(ui.TextProps{}, "Page body")
	if err := SidebarApp(DefaultProps(), body).Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	html := buf.String()

	if !strings.Contains(strings.ToLower(html), "<!doctype html>") {
		t.Fatal("expected full document with doctype")
	}
	if !strings.Contains(html, "<title>Dashboard · FastyGo</title>") {
		t.Fatal("expected document title")
	}
	if !strings.Contains(html, "Page body") {
		t.Fatal("expected page body slot")
	}
}

func TestSidebarApp_rendersDesktopAsideNav(t *testing.T) {
	var buf bytes.Buffer
	body := ui.Text(ui.TextProps{}, "Page body")
	if err := SidebarApp(DefaultProps(), body).Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	html := buf.String()

	if !strings.Contains(html, "<aside") {
		t.Fatal("expected desktop sidebar aside")
	}
	if !strings.Contains(html, `aria-label="App navigation"`) {
		t.Fatal("expected sidebar aria label")
	}
	if !strings.Contains(html, ">Home<") || !strings.Contains(html, ">Sample<") {
		t.Fatal("expected sidebar nav links")
	}
	if !strings.Contains(html, `aria-current="page"`) {
		t.Fatal("expected active nav item aria-current")
	}
}

func TestSidebarApp_rendersMobileSheetContract(t *testing.T) {
	var buf bytes.Buffer
	body := ui.Text(ui.TextProps{}, "Page body")
	if err := SidebarApp(DefaultProps(), body).Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	html := buf.String()

	if !strings.Contains(html, `data-ui8kit="sheet"`) {
		t.Fatal("expected mobile sheet markup")
	}
	if !strings.Contains(html, `id="ui8kit-mobile-sheet-panel"`) {
		t.Fatal("expected mobile sheet panel id")
	}
	if !strings.Contains(html, `id="ui8kit-mobile-sheet-trigger"`) {
		t.Fatal("expected mobile sheet trigger id")
	}
}
