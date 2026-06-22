package fixtures_test

import (
	"encoding/json"
	"reflect"
	"slices"
	"testing"

	"github.com/fastygo/blank/internal/fixtures"
)

func TestLoadLocale_enAndRuShareKeys(t *testing.T) {
	en, err := fixtures.LoadLocale("en")
	if err != nil {
		t.Fatal(err)
	}
	ru, err := fixtures.LoadLocale("ru")
	if err != nil {
		t.Fatal(err)
	}

	enKeys := jsonKeyPaths(t, en)
	ruKeys := jsonKeyPaths(t, ru)
	slices.Sort(enKeys)
	slices.Sort(ruKeys)
	if !reflect.DeepEqual(enKeys, ruKeys) {
		t.Fatalf("locale key mismatch\nen: %#v\nru: %#v", enKeys, ruKeys)
	}
}

func TestResolve_fallsBackToDefault(t *testing.T) {
	loc := fixtures.Resolve("zz", "en")
	if loc.Home.Title != "Home" {
		t.Fatalf("expected english fallback, got %q", loc.Home.Title)
	}
}

func jsonKeyPaths(t *testing.T, value any) []string {
	t.Helper()
	raw, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	var out any
	if err := json.Unmarshal(raw, &out); err != nil {
		t.Fatal(err)
	}
	return collectKeyPaths(out, "")
}

func collectKeyPaths(value any, prefix string) []string {
	switch typed := value.(type) {
	case map[string]any:
		var paths []string
		for key, child := range typed {
			path := key
			if prefix != "" {
				path = prefix + "." + key
			}
			paths = append(paths, path)
			paths = append(paths, collectKeyPaths(child, path)...)
		}
		return paths
	default:
		return nil
	}
}
