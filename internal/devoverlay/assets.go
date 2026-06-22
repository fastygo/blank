package devoverlay

import (
	"os"
	"path/filepath"
	"time"

	"github.com/fastygo/blank/internal/devoverlay/fixtures"
)

// AssetStatus describes one tracked static asset on disk.
type AssetStatus struct {
	ID      string    `json:"id"`
	Path    string    `json:"path"`
	Exists  bool      `json:"exists"`
	Size    int64     `json:"size"`
	MTime   time.Time `json:"mtime"`
	AgeSec  int64     `json:"ageSec"`
	Stale   bool      `json:"stale"`
	Hint    string    `json:"hint,omitempty"`
}

// StatusPayload is returned by GET /__fastygo/dev/status.json.
type StatusPayload struct {
	Bind    string        `json:"bind"`
	Assets  []AssetStatus `json:"assets"`
	Hints   []string      `json:"hints,omitempty"`
	Overlay bool          `json:"overlay"`
}

func collectAssetStatus(cfg Config, copy fixtures.Assets) StatusPayload {
	now := time.Now()
	staleSeconds := cfg.StaleCSSSeconds
	if staleSeconds <= 0 {
		staleSeconds = defaultStaleCSSSeconds
	}

	payload := StatusPayload{
		Bind:    cfg.Bind,
		Overlay: cfg.Enabled,
	}
	for _, asset := range cfg.Assets {
		status := AssetStatus{
			ID:   asset.ID,
			Path: filepath.ToSlash(filepath.Join("web/static", asset.Path)),
		}
		full := filepath.Join(cfg.StaticDir, asset.Path)
		info, err := os.Stat(full)
		if err != nil {
			status.Exists = false
			if asset.ID == "app.css" {
				status.Hint = copy.MissingCSSHint
				payload.Hints = append(payload.Hints, status.Hint)
			}
			payload.Assets = append(payload.Assets, status)
			continue
		}
		status.Exists = true
		status.Size = info.Size()
		status.MTime = info.ModTime().UTC()
		status.AgeSec = int64(now.Sub(info.ModTime()).Seconds())
		if asset.ID == "app.css" && status.AgeSec > staleSeconds {
			status.Stale = true
			status.Hint = copy.StaleCSSHint
			payload.Hints = append(payload.Hints, status.Hint)
		}
		payload.Assets = append(payload.Assets, status)
	}
	return payload
}
