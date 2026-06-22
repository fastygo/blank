package serverapp

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/fastygo/framework/pkg/app"
)

func runWithHandler(ctx context.Context, application *app.App, handler http.Handler) error {
	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	logger := slog.Default()
	cfg := application.Config()

	if workers := application.Workers(); workers != nil {
		workers.Start(runCtx)
	}

	server := &http.Server{
		Addr:              cfg.AppBind,
		Handler:           handler,
		ReadTimeout:       cfg.HTTPReadTimeout,
		ReadHeaderTimeout: cfg.HTTPReadHeaderTimeout,
		WriteTimeout:      cfg.HTTPWriteTimeout,
		IdleTimeout:       cfg.HTTPIdleTimeout,
		MaxHeaderBytes:    cfg.HTTPMaxHeaderBytes,
	}

	shutdownTimeout := cfg.HTTPShutdownTimeout
	if shutdownTimeout <= 0 {
		shutdownTimeout = 5 * time.Second
	}

	errCh := make(chan error, 1)
	go func() {
		logger.Info("app:listen", "addr", cfg.AppBind)
		errCh <- server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancelShutdown := context.WithTimeout(context.WithoutCancel(ctx), shutdownTimeout)
		defer cancelShutdown()
		shutdownErr := server.Shutdown(shutdownCtx)
		if workers := application.Workers(); workers != nil {
			if err := workers.Stop(shutdownCtx); err != nil {
				logger.Warn("workers stop", "error", err)
			}
		}
		return shutdownErr
	case err := <-errCh:
		cancel()
		stopParent := context.WithoutCancel(ctx)
		if workers := application.Workers(); workers != nil {
			stopCtx, cancelStop := context.WithTimeout(stopParent, shutdownTimeout)
			if stopErr := workers.Stop(stopCtx); stopErr != nil {
				logger.Warn("workers stop", "error", stopErr)
			}
			cancelStop()
		}
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("server listen: %w", err)
		}
		return err
	}
}
