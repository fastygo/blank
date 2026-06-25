package hero

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestHero_rendersWelcomeCopy(t *testing.T) {
	props := DefaultProps()
	var buf bytes.Buffer
	if err := Hero(props).Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	html := buf.String()
	for _, want := range []string{props.Welcome, props.WelcomeBrand, props.Description} {
		if !strings.Contains(html, want) {
			t.Fatalf("expected hero copy %q in output", want)
		}
	}
}
