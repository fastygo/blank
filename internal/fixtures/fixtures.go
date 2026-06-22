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
	Brand               string `json:"brand"`
	Footer              string `json:"footer"`
	Theme               Theme  `json:"theme"`
	LanguageToggleLabel string `json:"language_toggle_label"`
	Language            map[string]string `json:"language"`
	Layout              Layout            `json:"layout"`
	Home                Home              `json:"home"`
	Sample              Sample            `json:"sample"`
	DevOverlay          DevOverlay        `json:"dev_overlay"`
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
	MainNavigation    string `json:"main_navigation_label"`
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

// DevOverlay copy for the local dev widget (SSR + client panels).
type DevOverlay struct {
	LauncherAriaLabel   string          `json:"launcher_aria_label"`
	PanelAriaLabel      string          `json:"panel_aria_label"`
	Title               string          `json:"title"`
	Subtitle            string          `json:"subtitle"`
	TablistAriaLabel    string          `json:"tablist_aria_label"`
	HideButtonLabel     string          `json:"hide_button_label"`
	HideButtonAriaLabel string          `json:"hide_button_aria_label"`
	LoadingTab          string          `json:"loading_tab"`
	Tabs                DevOverlayTabs  `json:"tabs"`
	Health              DevOverlayHealth `json:"health"`
	Assets              DevOverlayAssets `json:"assets"`
	Request             DevOverlayRequest `json:"request"`
	Errors              DevOverlayErrors  `json:"errors"`
}

// DevOverlayTabs labels the overlay tab buttons.
type DevOverlayTabs struct {
	Health  string `json:"health"`
	Assets  string `json:"assets"`
	Request string `json:"request"`
}

// DevOverlayHealth copy for the health panel.
type DevOverlayHealth struct {
	Intro             string `json:"intro"`
	RefreshButton     string `json:"refresh_button"`
	OkBadge           string `json:"ok_badge"`
	HttpBadge         string `json:"http_badge"`
	DownLabel         string `json:"down_label"`
	LatencyUnit       string `json:"latency_unit"`
	Separator         string `json:"separator"`
}

// DevOverlayAssets copy for the assets panel.
type DevOverlayAssets struct {
	Intro          string `json:"intro"`
	Present        string `json:"present"`
	Missing        string `json:"missing"`
	Stale          string `json:"stale"`
	AgePrefix      string `json:"age_prefix"`
	Separator      string `json:"separator"`
	LoadFailed     string `json:"load_failed"`
	MissingCSSHint string `json:"missing_css_hint"`
	StaleCSSHint   string `json:"stale_css_hint"`
}

// DevOverlayRequest copy for the request panel.
type DevOverlayRequest struct {
	Intro           string `json:"intro"`
	RequestIDLabel  string `json:"request_id_label"`
	PathLabel       string `json:"path_label"`
	LocaleLabel     string `json:"locale_label"`
	CopyRequestID   string `json:"copy_request_id"`
	Copied          string `json:"copied"`
	CopyFailed      string `json:"copy_failed"`
	EmptyValue      string `json:"empty_value"`
	LocaleMissing   string `json:"locale_missing"`
}

// DevOverlayErrors copy for client-side error messages.
type DevOverlayErrors struct {
	FetchFailed      string `json:"fetch_failed"`
	StatusJSONFailed string `json:"status_json_failed"`
	DisableFailed    string `json:"disable_failed"`
}

// DevOverlayPanelJSON is serialized into the overlay script tag for client panels.
type DevOverlayPanelJSON struct {
	Health  DevOverlayHealth  `json:"health"`
	Assets  DevOverlayAssets  `json:"assets"`
	Request DevOverlayRequest `json:"request"`
	Errors  DevOverlayErrors  `json:"errors"`
}

// PanelJSON returns escaped-ready JSON for overlay client panels.
func (l Locale) PanelJSON() (string, error) {
	payload := DevOverlayPanelJSON{
		Health:  l.DevOverlay.Health,
		Assets:  l.DevOverlay.Assets,
		Request: l.DevOverlay.Request,
		Errors:  l.DevOverlay.Errors,
	}
	raw, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	return string(raw), nil
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
