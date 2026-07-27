package utils

import (
	"strconv"
	"strings"

	"github.com/a-h/templ"
)

// AriaExpanded sets aria-expanded.
func AriaExpanded(v bool) templ.Attributes {
	return templ.Attributes{"aria-expanded": strconv.FormatBool(v)}
}

// AriaControls sets aria-controls when id is non-empty.
func AriaControls(id string) templ.Attributes {
	if strings.TrimSpace(id) == "" {
		return templ.Attributes{}
	}
	return templ.Attributes{"aria-controls": id}
}

// AriaCurrent sets aria-current when value is non-empty.
func AriaCurrent(value string) templ.Attributes {
	if strings.TrimSpace(value) == "" {
		return templ.Attributes{}
	}
	return templ.Attributes{"aria-current": value}
}

// AriaLive sets aria-live when value is non-empty.
func AriaLive(value string) templ.Attributes {
	if strings.TrimSpace(value) == "" {
		return templ.Attributes{}
	}
	return templ.Attributes{"aria-live": value}
}

// AriaModal sets aria-modal.
func AriaModal(v bool) templ.Attributes {
	return templ.Attributes{"aria-modal": strconv.FormatBool(v)}
}

// AriaLabel sets aria-label when value is non-empty.
func AriaLabel(value string) templ.Attributes {
	if strings.TrimSpace(value) == "" {
		return templ.Attributes{}
	}
	return templ.Attributes{"aria-label": value}
}

// AriaPressed sets aria-pressed.
func AriaPressed(pressed bool) templ.Attributes {
	return templ.Attributes{"aria-pressed": strconv.FormatBool(pressed)}
}

// AriaHasPopup sets aria-haspopup for known kinds.
func AriaHasPopup(kind string) templ.Attributes {
	switch strings.TrimSpace(strings.ToLower(kind)) {
	case "true", "menu", "listbox", "tree", "grid", "dialog":
		return templ.Attributes{"aria-haspopup": strings.TrimSpace(strings.ToLower(kind))}
	default:
		return templ.Attributes{}
	}
}
