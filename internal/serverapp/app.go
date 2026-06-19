package serverapp

import (
	"log/slog"
	"os"
	"strings"

	"github.com/fastygo/blank/internal/platform"
	"github.com/fastygo/blank/internal/site"
	"github.com/fastygo/framework/pkg/app"
	"github.com/fastygo/framework/pkg/web/locale"
	"github.com/fastygo/framework/pkg/web/security"
)

// New builds the assembled HTTP application.
func New() (*app.App, error) {
	cfg, err := platform.Load()
	if err != nil {
		return nil, err
	}

	logger := newLogger(cfg.LogLevel, cfg.LogFormat)
	slog.SetDefault(logger)

	feat := site.NewFeature(cfg.AvailableLocales, cfg.DefaultLocale)

	return app.New(cfg.Config).
		WithLogger(logger).
		WithSecurity(security.LoadConfig()).
		WithLocales(app.LocalesConfig{
			Default:   cfg.DefaultLocale,
			Available: cfg.AvailableLocales,
			Strategy:  nil,
			Cookie: locale.CookieOptions{
				Enabled: true,
				Name:    "lang",
			},
			SPA: true,
		}).
		WithHealthEndpoints("/healthz", "/readyz").
		WithFeature(feat).
		Build(), nil
}

func newLogger(level, format string) *slog.Logger {
	var lvl slog.Level
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "debug":
		lvl = slog.LevelDebug
	case "warn", "warning":
		lvl = slog.LevelWarn
	case "error":
		lvl = slog.LevelError
	default:
		lvl = slog.LevelInfo
	}
	opts := &slog.HandlerOptions{Level: lvl}
	var h slog.Handler
	if strings.ToLower(strings.TrimSpace(format)) == "json" {
		h = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		h = slog.NewTextHandler(os.Stdout, opts)
	}
	return slog.New(h)
}
