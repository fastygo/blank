package ui

import "github.com/fastygo/blank/internal/devoverlay/fixtures"

// HostProps drives the dev overlay loader script.
type HostProps struct {
	RequestID string
	PanelJSON string
}

// NewHostProps serializes localized overlay copy for the web component.
func NewHostProps(requestID string, fix fixtures.Locale) (HostProps, error) {
	panelJSON, err := fix.PanelJSON()
	if err != nil {
		return HostProps{}, err
	}
	return HostProps{
		RequestID: requestID,
		PanelJSON: panelJSON,
	}, nil
}
