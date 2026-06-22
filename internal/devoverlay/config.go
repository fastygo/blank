package devoverlay

import (
	"net"
	"os"
	"strings"

	"github.com/fastygo/framework/pkg/app"
)

const (
	EnvEnabled  = "APP_DEV_OVERLAY"
	CookieName  = "fastygo_dev"
	CookieOff   = "off"
	CookieOn    = "on"
	RoutePrefix = "/__fastygo/dev"
)

// AssetConfig describes a static asset tracked by the overlay status API.
type AssetConfig struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

// Config controls dev overlay availability and asset inspection.
type Config struct {
	Enabled          bool
	Bind             string
	StaticDir        string
	CookieName       string
	DefaultLocale    string
	AvailableLocales []string
	LangCookieName   string
	Assets           []AssetConfig
}

// Load resolves overlay config from app config and environment.
func Load(cfg app.Config) Config {
	out := Config{
		Enabled:          parseBoolEnv(EnvEnabled, false),
		Bind:             cfg.AppBind,
		StaticDir:        cfg.StaticDir,
		CookieName:       CookieName,
		DefaultLocale:    cfg.DefaultLocale,
		AvailableLocales: cfg.AvailableLocales,
		LangCookieName:   defaultLangCookieName,
		Assets: []AssetConfig{
			{ID: "app.css", Path: "css/app.css"},
			{ID: "ui8kit.js", Path: "js/ui8kit.js"},
			{ID: "theme.js", Path: "js/theme.js"},
		},
	}
	if out.Enabled && !isLoopbackBind(out.Bind) {
		out.Enabled = false
	}
	return out
}

func parseBoolEnv(key string, fallback bool) bool {
	raw := strings.TrimSpace(strings.ToLower(os.Getenv(key)))
	switch raw {
	case "1", "true", "yes", "on":
		return true
	case "0", "false", "no", "off":
		return false
	default:
		return fallback
	}
}

func isLoopbackBind(bind string) bool {
	host, _, err := net.SplitHostPort(strings.TrimSpace(bind))
	if err != nil {
		// Accept bare ":8080" for local dev.
		if strings.HasPrefix(strings.TrimSpace(bind), ":") {
			return true
		}
		return false
	}
	host = strings.TrimSpace(strings.ToLower(host))
	switch host {
	case "", "127.0.0.1", "localhost", "::1", "[::1]":
		return true
	default:
		return false
	}
}

func (c Config) optedOutValue(cookieValue string) bool {
	return strings.EqualFold(strings.TrimSpace(cookieValue), CookieOff)
}
