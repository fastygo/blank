package ui

import "github.com/fastygo/blank/internal/fixtures"

// HostProps drives localized dev overlay markup.
type HostProps struct {
	RequestID string
	Copy      fixtures.DevOverlay
	PanelJSON string
}
