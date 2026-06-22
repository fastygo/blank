package ui

import "github.com/fastygo/blank/internal/devoverlay/fixtures"

// HostProps drives localized dev overlay markup.
type HostProps struct {
	RequestID string
	Copy      fixtures.Locale
	PanelJSON string
}
