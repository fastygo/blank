package devoverlay

import (
	_ "embed"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

//go:embed static/overlay.js
var overlayJS []byte

// Service serves dev overlay routes and wraps the application handler.
type Service struct {
	cfg Config
}

// Wrap returns a handler that serves dev routes and optionally injects overlay markup.
func Wrap(inner http.Handler, cfg Config) http.Handler {
	if !cfg.Enabled {
		return inner
	}
	svc := &Service{cfg: cfg}
	mux := http.NewServeMux()
	mux.HandleFunc("GET "+RoutePrefix+"/overlay.js", svc.handleOverlayJS)
	mux.HandleFunc("GET "+RoutePrefix+"/status.json", svc.handleStatusJSON)
	mux.HandleFunc("POST "+RoutePrefix+"/disable", svc.handleDisable)
	mux.HandleFunc("GET "+RoutePrefix+"/enable", svc.handleEnable)
	mux.Handle("/", injectMiddleware{cfg: cfg, inner: inner})
	return mux
}

func (s *Service) handleOverlayJS(w http.ResponseWriter, _ *http.Request) {
	if len(overlayJS) == 0 {
		http.NotFound(w, nil)
		return
	}
	setDevHeaders(w)
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	_, _ = w.Write(overlayJS)
}

func (s *Service) handleStatusJSON(w http.ResponseWriter, _ *http.Request) {
	setDevHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	payload := collectAssetStatus(s.cfg)
	_ = json.NewEncoder(w).Encode(payload)
}

func (s *Service) handleDisable(w http.ResponseWriter, r *http.Request) {
	setDevHeaders(w)
	http.SetCookie(w, &http.Cookie{
		Name:     s.cfg.CookieName,
		Value:    CookieOff,
		Path:     "/",
		MaxAge:   31536000,
		SameSite: http.SameSiteLaxMode,
	})
	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"ok":true}`))
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func (s *Service) handleEnable(w http.ResponseWriter, _ *http.Request) {
	setDevHeaders(w)
	http.SetCookie(w, &http.Cookie{
		Name:     s.cfg.CookieName,
		Value:    CookieOn,
		Path:     "/",
		MaxAge:   31536000,
		SameSite: http.SameSiteLaxMode,
	})
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write([]byte(`{"ok":true}`))
}

func setDevHeaders(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("X-Robots-Tag", "noindex, nofollow")
}

// FormatAge returns a human-readable age string for tests and UI helpers.
func FormatAge(seconds int64) string {
	if seconds < 60 {
		return strconv.FormatInt(seconds, 10) + "s"
	}
	if seconds < 3600 {
		return strconv.FormatInt(seconds/60, 10) + "m"
	}
	return strconv.FormatInt(seconds/3600, 10) + "h"
}

// NowUTC is a test seam for asset age calculations.
var NowUTC = func() time.Time { return time.Now().UTC() }
