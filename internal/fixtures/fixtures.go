package fixtures

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"
)

//go:embed locale/*.json
var localeFS embed.FS

// Locale holds embedded copy for one language.
type Locale struct {
	Brand               string            `json:"brand"`
	Footer              string            `json:"footer"`
	Theme               Theme             `json:"theme"`
	LanguageToggleLabel string            `json:"language_toggle_label"`
	Language            map[string]string `json:"language"`
	Layout              Layout            `json:"layout"`
	Home                Home              `json:"home"`
	Sample              Sample            `json:"sample"`
}

// Theme copy for the theme toggle control.
type Theme struct {
	Label             string `json:"label"`
	SwitchToDarkLabel string `json:"switch_to_dark"`
	SwitchToLight     string `json:"switch_to_light"`
}

// Layout copy for shell navigation and chrome.
type Layout struct {
	BrandHomeSuffix     string `json:"brand_home_suffix"`
	MainNavigation      string `json:"main_navigation_label"`
	OpenNavigationMenu  string `json:"open_navigation_menu"`
	CloseNavigationMenu string `json:"close_navigation_menu"`
	NavigationMenuLabel string `json:"navigation_menu_label"`
}

// Home page copy.
type Home struct {
	NavLabel     string `json:"nav_label"`
	Title        string `json:"title"`
	Welcome      string `json:"welcome"`
	WelcomeBrand string `json:"welcome_brand"`
	Description  string `json:"description"`
}

// Sample page copy.
type Sample struct {
	NavLabel    string `json:"nav_label"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
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

// Resolve loads locale by code with fallback to defaultCode.
func Resolve(code, defaultCode string) Locale {
	code = strings.TrimSpace(code)
	if code == "" {
		code = strings.TrimSpace(defaultCode)
	}
	loc, err := LoadLocale(code)
	if err != nil {
		loc, _ = LoadLocale(defaultCode)
	}
	return loc
}
