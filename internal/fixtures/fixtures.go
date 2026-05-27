package fixtures

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed locale/*.json
var localeFS embed.FS

// Locale holds embedded copy for one language.
type Locale struct {
	Brand               string `json:"brand"`
	Footer              string `json:"footer"`
	Theme               struct {
		Label             string `json:"label"`
		SwitchToDarkLabel string `json:"switch_to_dark"`
		SwitchToLight     string `json:"switch_to_light"`
	} `json:"theme"`
	LanguageToggleLabel string `json:"language_toggle_label"`
	Home                struct {
		NavLabel     string `json:"nav_label"`
		Title        string `json:"title"`
		Welcome      string `json:"welcome"`
		WelcomeBrand string `json:"welcome_brand"`
		Description  string `json:"description"`
	} `json:"home"`
	Sample struct {
		NavLabel    string `json:"nav_label"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Body        string `json:"body"`
	} `json:"sample"`
}

// LoadLocale reads locale/{code}.json (e.g. en, ru).
func LoadLocale(code string) (Locale, error) {
	raw, err := localeFS.ReadFile("locale/" + code + ".json")
	if err != nil {
		return Locale{}, fmt.Errorf("fixtures: read locale %q: %w", code, err)
	}
	var out Locale
	if err := json.Unmarshal(raw, &out); err != nil {
		return Locale{}, fmt.Errorf("fixtures: parse locale %q: %w", code, err)
	}
	return out, nil
}
