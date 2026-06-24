package navigation

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func sampleItems() []Item {
	return []Item{
		{Label: "Home", Path: "/", Icon: "home"},
		{Label: "Sample", Path: "/sample", Icon: "box"},
	}
}

func TestMobileSheet_rendersUi8kitContract(t *testing.T) {
	var buf bytes.Buffer
	if err := MobileSheet(MobileSheetProps{
		Title:      "FastyGo",
		AriaLabel:  "Navigation menu",
		CloseLabel: "Close navigation menu",
		Items:      sampleItems(),
		Active:     "/",
	}).Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	html := buf.String()

	if !strings.Contains(html, `id="`+MobileSheetPanelID+`"`) {
		t.Fatal("expected mobile sheet panel id")
	}
	if !strings.Contains(html, `id="`+MobileSheetTitleID+`"`) {
		t.Fatal("expected mobile sheet title id")
	}
	if !strings.Contains(html, `data-ui8kit="sheet"`) {
		t.Fatal("expected sheet ui8kit hook")
	}
	if !strings.Contains(html, `role="dialog"`) {
		t.Fatal("expected dialog role")
	}
	if !strings.Contains(html, `aria-modal="true"`) {
		t.Fatal("expected aria-modal")
	}
	if !strings.Contains(html, `aria-label="Navigation menu"`) {
		t.Fatal("expected sheet aria-label")
	}
	if !strings.Contains(html, `aria-current="page"`) {
		t.Fatal("expected active nav link")
	}
	if !strings.Contains(html, "Home") || !strings.Contains(html, "Sample") {
		t.Fatal("expected nav labels")
	}
}

func TestMobileSheetTrigger_rendersUi8kitContract(t *testing.T) {
	var buf bytes.Buffer
	if err := MobileSheetTrigger(MobileSheetTriggerProps{
		AriaLabel: "Open navigation menu",
	}).Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	html := buf.String()

	if !strings.Contains(html, `id="`+MobileSheetTriggerID+`"`) {
		t.Fatal("expected mobile sheet trigger id")
	}
	if !strings.Contains(html, `aria-controls="`+MobileSheetPanelID+`"`) {
		t.Fatal("expected aria-controls")
	}
	if !strings.Contains(html, `aria-haspopup="dialog"`) {
		t.Fatal("expected aria-haspopup dialog")
	}
	if !strings.Contains(html, `aria-expanded="false"`) {
		t.Fatal("expected aria-expanded false")
	}
	if !strings.Contains(html, `data-ui8kit-dialog-open`) {
		t.Fatal("expected ui8kit dialog open hook")
	}
	if !strings.Contains(html, `data-ui8kit-dialog-target="`+MobileSheetPanelID+`"`) {
		t.Fatal("expected ui8kit dialog target")
	}
}
