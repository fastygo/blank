package devoverlay

import (
	"bytes"
	"context"

	"github.com/fastygo/blank/internal/devoverlay/ui"
)

func injectOverlay(html []byte, requestID string) ([]byte, error) {
	lower := bytes.ToLower(html)
	idx := bytes.LastIndex(lower, []byte("</body>"))
	if idx < 0 {
		return html, nil
	}

	var buf bytes.Buffer
	if err := ui.Host(requestID).Render(context.Background(), &buf); err != nil {
		return nil, err
	}

	out := make([]byte, 0, len(html)+buf.Len())
	out = append(out, html[:idx]...)
	out = append(out, buf.Bytes()...)
	out = append(out, html[idx:]...)
	return out, nil
}
