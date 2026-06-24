package fixtures

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"
)

//go:embed locale/*.json
var localeFS embed.FS

// Locale holds embedded copy for the dev overlay widget.
type Locale struct {
	LauncherAriaLabel   string `json:"launcher_aria_label"`
	PanelAriaLabel      string `json:"panel_aria_label"`
	Title               string `json:"title"`
	Subtitle            string `json:"subtitle"`
	TablistAriaLabel    string `json:"tablist_aria_label"`
	HideButtonLabel     string `json:"hide_button_label"`
	HideButtonAriaLabel string `json:"hide_button_aria_label"`
	MobileReloadHint    string `json:"mobile_reload_hint"`
	MobileReloadButton  string `json:"mobile_reload_button"`
	LoadingTab          string `json:"loading_tab"`
	Tabs                Tabs   `json:"tabs"`
	Health              Health `json:"health"`
	Assets              Assets `json:"assets"`
	Request             Request `json:"request"`
	Errors              Errors  `json:"errors"`
}

// Tabs labels the overlay tab buttons.
type Tabs struct {
	Health  string `json:"health"`
	Assets  string `json:"assets"`
	Request string `json:"request"`
}

// Health copy for the health panel.
type Health struct {
	Intro         string `json:"intro"`
	RefreshButton string `json:"refresh_button"`
	OkBadge       string `json:"ok_badge"`
	HttpBadge     string `json:"http_badge"`
	DownLabel     string `json:"down_label"`
	LatencyUnit   string `json:"latency_unit"`
	Separator     string `json:"separator"`
}

// Assets copy for the assets panel.
type Assets struct {
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

// Request copy for the request panel.
type Request struct {
	Intro          string `json:"intro"`
	RequestIDLabel string `json:"request_id_label"`
	PathLabel      string `json:"path_label"`
	LocaleLabel    string `json:"locale_label"`
	CopyRequestID  string `json:"copy_request_id"`
	Copied         string `json:"copied"`
	CopyFailed     string `json:"copy_failed"`
	EmptyValue     string `json:"empty_value"`
	LocaleMissing  string `json:"locale_missing"`
}

// Errors copy for client-side error messages.
type Errors struct {
	FetchFailed      string `json:"fetch_failed"`
	StatusJSONFailed string `json:"status_json_failed"`
	DisableFailed    string `json:"disable_failed"`
}

// PanelJSON returns escaped-ready JSON for the overlay web component.
func (l Locale) PanelJSON() (string, error) {
	raw, err := json.Marshal(l)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

// LoadLocale reads locale/{code}.json (e.g. en, ru).
func LoadLocale(code string) (Locale, error) {
	raw, err := localeFS.ReadFile("locale/" + code + ".json")
	if err != nil {
		return Locale{}, fmt.Errorf("devoverlay fixtures: read locale %q: %w", code, err)
	}
	var out Locale
	if err := json.Unmarshal(raw, &out); err != nil {
		return Locale{}, fmt.Errorf("devoverlay fixtures: parse locale %q: %w", code, err)
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
