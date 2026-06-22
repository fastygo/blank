package devoverlay_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/fastygo/blank/internal/devoverlay"
	"github.com/fastygo/framework/pkg/app"
)

func testConfig(t *testing.T) devoverlay.Config {
	t.Helper()
	dir := t.TempDir()
	static := filepath.Join(dir, "static")
	if err := os.MkdirAll(filepath.Join(static, "css"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(static, "js"), 0o755); err != nil {
		t.Fatal(err)
	}
	writeFile(t, filepath.Join(static, "css", "app.css"), "body{}")
	writeFile(t, filepath.Join(static, "js", "ui8kit.js"), "console.log('ui8kit')")
	writeFile(t, filepath.Join(static, "js", "theme.js"), "console.log('theme')")

	return devoverlay.Config{
		Enabled:          true,
		Bind:             "127.0.0.1:8080",
		StaticDir:        static,
		CookieName:       devoverlay.CookieName,
		DefaultLocale:    "en",
		AvailableLocales: []string{"en", "ru"},
		LangCookieName:   "lang",
		StaleCSSSeconds:  300,
		Assets: []devoverlay.AssetConfig{
			{ID: "app.css", Path: "css/app.css"},
			{ID: "ui8kit.js", Path: "js/ui8kit.js"},
			{ID: "theme.js", Path: "js/theme.js"},
		},
	}
}

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestInjectMiddlewareUsesRequestLocale(t *testing.T) {
	cfg := testConfig(t)
	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = io.WriteString(w, `<!DOCTYPE html><html lang="ru"><body><main>Home</main></body></html>`)
	})

	handler := devoverlay.Wrap(inner, cfg)
	req := httptest.NewRequest(http.MethodGet, "/?lang=ru", nil)
	req.Header.Set("Accept", "text/html")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	body := rec.Body.String()
	if !strings.Contains(body, "Панель разработчика FastyGo") {
		t.Fatalf("expected russian overlay copy, got: %s", body)
	}
	if !strings.Contains(body, "Состояние") {
		t.Fatalf("expected russian health tab, got: %s", body)
	}
}

func TestInjectMiddlewareInjectsOverlayMarkup(t *testing.T) {
	cfg := testConfig(t)
	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("X-Request-ID", "req-test-123")
		_, _ = io.WriteString(w, "<!DOCTYPE html><html><body><main>Home</main></body></html>")
	})

	handler := devoverlay.Wrap(inner, cfg)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Accept", "text/html")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	body := rec.Body.String()
	if !strings.Contains(body, `id="fastygo-dev-overlay-root"`) {
		t.Fatalf("expected overlay root, got: %s", body)
	}
	if !strings.Contains(body, `data-request-id="req-test-123"`) {
		t.Fatalf("expected request id on script, got: %s", body)
	}
	if !strings.Contains(body, "/__fastygo/dev/overlay.js") {
		t.Fatalf("expected overlay script, got: %s", body)
	}
}

func TestInjectMiddlewareSkipsWhenCookieOff(t *testing.T) {
	cfg := testConfig(t)
	inner := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = io.WriteString(w, "<!DOCTYPE html><html><body>Clean</body></html>")
	})
	handler := devoverlay.Wrap(inner, cfg)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Accept", "text/html")
	req.AddCookie(&http.Cookie{Name: devoverlay.CookieName, Value: devoverlay.CookieOff})
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	body := rec.Body.String()
	if strings.Contains(body, "fastygo-dev-overlay") {
		t.Fatalf("expected clean html, got: %s", body)
	}
}

func TestInjectMiddlewareSkipsNonHTMLAndProbes(t *testing.T) {
	cfg := testConfig(t)
	inner := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"ok":true}`)
	})
	handler := devoverlay.Wrap(inner, cfg)

	for _, path := range []string{"/healthz", "/readyz", "/static/css/app.css"} {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		if strings.Contains(rec.Body.String(), "fastygo-dev-overlay") {
			t.Fatalf("path %s should not inject overlay", path)
		}
	}
}

func TestStatusJSONReportsAssets(t *testing.T) {
	cfg := testConfig(t)
	handler := devoverlay.Wrap(http.NotFoundHandler(), cfg)

	req := httptest.NewRequest(http.MethodGet, "/__fastygo/dev/status.json", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d", rec.Code)
	}
	body := rec.Body.String()
	if !strings.Contains(body, `"app.css"`) {
		t.Fatalf("expected assets in payload: %s", body)
	}
	if rec.Header().Get("Cache-Control") != "no-store" {
		t.Fatalf("expected no-store cache header")
	}
}

func TestDisableSetsOptOutCookie(t *testing.T) {
	cfg := testConfig(t)
	handler := devoverlay.Wrap(http.NotFoundHandler(), cfg)

	req := httptest.NewRequest(http.MethodPost, "/__fastygo/dev/disable", nil)
	req.Header.Set("Accept", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d", rec.Code)
	}
	cookies := rec.Result().Cookies()
	if len(cookies) == 0 || cookies[0].Value != devoverlay.CookieOff {
		t.Fatalf("expected opt-out cookie, got %#v", cookies)
	}
}

func TestLoadRequiresLoopback(t *testing.T) {
	t.Setenv(devoverlay.EnvEnabled, "1")
	cfg := devoverlay.Load(app.Config{
		AppBind:   "0.0.0.0:8080",
		StaticDir: "web/static",
	})
	if cfg.Enabled {
		t.Fatal("overlay must refuse non-loopback bind")
	}
}

func TestLoadEnabledOnLoopback(t *testing.T) {
	t.Setenv(devoverlay.EnvEnabled, "1")
	cfg := devoverlay.Load(app.Config{
		AppBind:   "127.0.0.1:8080",
		StaticDir: "web/static",
	})
	if !cfg.Enabled {
		t.Fatal("overlay should enable on loopback")
	}
}

func TestStatusJSONLocalizedHint(t *testing.T) {
	cfg := testConfig(t)
	path := filepath.Join(cfg.StaticDir, "css", "app.css")
	old := time.Now().Add(-10 * time.Minute)
	if err := os.Chtimes(path, old, old); err != nil {
		t.Fatal(err)
	}

	handler := devoverlay.Wrap(http.NotFoundHandler(), cfg)
	req := httptest.NewRequest(http.MethodGet, "/__fastygo/dev/status.json?lang=ru", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	body := rec.Body.String()
	if !strings.Contains(body, "CSS устарел") {
		t.Fatalf("expected russian stale css hint, got: %s", body)
	}
}

func TestAssetStatusStaleHint(t *testing.T) {
	cfg := testConfig(t)
	cfg.DefaultLocale = "en"
	path := filepath.Join(cfg.StaticDir, "css", "app.css")
	old := time.Now().Add(-10 * time.Minute)
	if err := os.Chtimes(path, old, old); err != nil {
		t.Fatal(err)
	}

	handler := devoverlay.Wrap(http.NotFoundHandler(), cfg)
	req := httptest.NewRequest(http.MethodGet, "/__fastygo/dev/status.json", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	body := rec.Body.String()
	if !strings.Contains(body, "watch:css") {
		t.Fatalf("expected stale css hint, got: %s", body)
	}
}
