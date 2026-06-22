package serverapp

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/fastygo/blank/internal/devoverlay"
	"github.com/fastygo/blank/internal/platform"
	"github.com/fastygo/blank/internal/site"
	"github.com/fastygo/framework/pkg/app"
	"github.com/fastygo/framework/pkg/web/locale"
	"github.com/fastygo/framework/pkg/web/security"
)

// Application is the Blank composition root with optional dev overlay wiring.
type Application struct {
	app     *app.App
	overlay devoverlay.Config
	handler http.Handler
}

// New builds the assembled HTTP application.
func New() (*Application, error) {
	cfg, err := platform.Load()
	if err != nil {
		return nil, err
	}

	logger := newLogger(cfg.LogLevel, cfg.LogFormat)
	slog.SetDefault(logger)

	feat := site.NewFeature(cfg.AvailableLocales, cfg.DefaultLocale)
	application := app.New(cfg.Config).
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
		Build()

	overlayCfg := devoverlay.Load(cfg.Config)
	handler := application.Handler()
	if overlayCfg.Enabled {
		slog.Info("devoverlay:enabled", "bind", overlayCfg.Bind, "routes", devoverlay.RoutePrefix)
		handler = devoverlay.Wrap(handler, overlayCfg)
	} else {
		slog.Info("devoverlay:disabled", "env", devoverlay.EnvEnabled)
	}

	return &Application{
		app:     application,
		overlay: overlayCfg,
		handler: handler,
	}, nil
}

// Run starts the HTTP server and blocks until ctx is cancelled or the server exits.
func (a *Application) Run(ctx context.Context) error {
	if !a.overlay.Enabled {
		return a.app.Run(ctx)
	}
	return runWithHandler(ctx, a.app, a.handler)
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
