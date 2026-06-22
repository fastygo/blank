package devoverlay

import (
	"bytes"
	"net/http"

	"github.com/fastygo/blank/internal/devoverlay/ui"
	"github.com/fastygo/blank/internal/fixtures"
)

func injectOverlay(html []byte, r *http.Request, cfg Config, requestID string) ([]byte, error) {
	lower := bytes.ToLower(html)
	idx := bytes.LastIndex(lower, []byte("</body>"))
	if idx < 0 {
		return html, nil
	}

	fix := fixtures.Resolve(cfg.resolveLocale(r), cfg.DefaultLocale)
	props, err := ui.NewHostProps(requestID, fix)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := ui.Host(props).Render(r.Context(), &buf); err != nil {
		return nil, err
	}

	out := make([]byte, 0, len(html)+buf.Len())
	out = append(out, html[:idx]...)
	out = append(out, buf.Bytes()...)
	out = append(out, html[idx:]...)
	return out, nil
}
