package devoverlay

import (
	"bytes"
	"net/http"
	"strings"
)

type injectMiddleware struct {
	cfg    Config
	inner  http.Handler
}

func (m injectMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !shouldInject(m.cfg, r) {
		m.inner.ServeHTTP(w, r)
		return
	}

	rec := &responseCapture{ResponseWriter: w}
	m.inner.ServeHTTP(rec, r)

	if rec.status == 0 {
		rec.status = http.StatusOK
	}
	if rec.status != http.StatusOK || !isHTMLResponse(rec.header) {
		writeCaptured(w, rec)
		return
	}

	requestID := rec.header.Get("X-Request-ID")
	body, err := injectOverlay(rec.body.Bytes(), r, m.cfg, requestID)
	if err != nil {
		writeCaptured(w, rec)
		return
	}

	copyHeaders(w, rec.header)
	w.Header().Del("Content-Length")
	w.WriteHeader(rec.status)
	_, _ = w.Write(body)
}

type responseCapture struct {
	http.ResponseWriter
	status int
	header http.Header
	body   bytes.Buffer
}

func (r *responseCapture) Header() http.Header {
	if r.header == nil {
		r.header = make(http.Header)
	}
	return r.header
}

func (r *responseCapture) WriteHeader(statusCode int) {
	r.status = statusCode
}

func (r *responseCapture) Write(p []byte) (int, error) {
	if r.status == 0 {
		r.status = http.StatusOK
	}
	return r.body.Write(p)
}

func shouldInject(cfg Config, r *http.Request) bool {
	if !cfg.Enabled {
		return false
	}
	if cfg.optedOut(r) {
		return false
	}
	if r.Method != http.MethodGet {
		return false
	}
	path := r.URL.Path
	if strings.HasPrefix(path, "/static/") {
		return false
	}
	if strings.HasPrefix(path, RoutePrefix) {
		return false
	}
	if path == "/healthz" || path == "/readyz" {
		return false
	}
	accept := r.Header.Get("Accept")
	if accept != "" && !strings.Contains(accept, "text/html") && !strings.Contains(accept, "*/*") {
		return false
	}
	return true
}

func (cfg Config) optedOut(r *http.Request) bool {
	if c, err := r.Cookie(cfg.CookieName); err == nil {
		return cfg.optedOutValue(c.Value)
	}
	return false
}

func isHTMLResponse(header http.Header) bool {
	ct := header.Get("Content-Type")
	return strings.Contains(strings.ToLower(ct), "text/html")
}

func copyHeaders(w http.ResponseWriter, src http.Header) {
	for key, values := range src {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
}

func writeCaptured(w http.ResponseWriter, rec *responseCapture) {
	copyHeaders(w, rec.header)
	if rec.status == 0 {
		rec.status = http.StatusOK
	}
	w.WriteHeader(rec.status)
	_, _ = w.Write(rec.body.Bytes())
}
