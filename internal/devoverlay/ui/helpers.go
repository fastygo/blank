package ui

import (
	"strings"

	"github.com/fastygo/blank/internal/devoverlay/fixtures"
)

// NewHostProps builds overlay markup props from resolved fixtures.
func NewHostProps(requestID string, fix fixtures.Locale) (HostProps, error) {
	panelJSON, err := fix.PanelJSON()
	if err != nil {
		return HostProps{}, err
	}
	return HostProps{
		RequestID: requestID,
		Copy:      fix,
		PanelJSON: panelJSON,
	}, nil
}

func loadingTabText(template, label string) string {
	if strings.Contains(template, "%s") {
		return strings.Replace(template, "%s", label, 1)
	}
	return template + " " + label
}
