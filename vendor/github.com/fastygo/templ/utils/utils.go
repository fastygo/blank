// Package uiutils holds shared Go helpers for registry bricks.
// See utils.spec.md for machine-readable export catalog.
package uiutils

import (
	"strconv"
	"strings"

	"github.com/a-h/templ"
)

// Variants groups a base utility chain and keyed variant maps (cva-style).
type Variants struct {
	Base     string
	Keys     []string
	ByKey    map[string]map[string]string
	Defaults map[string]string
}

// Compose merges Base, variant selections, and optional tail classes.
func Compose(v Variants, selection map[string]string, extra ...string) string {
	var parts []string
	if strings.TrimSpace(v.Base) != "" {
		parts = append(parts, strings.TrimSpace(v.Base))
	}
	for _, key := range v.Keys {
		choices, ok := v.ByKey[key]
		if !ok {
			continue
		}
		choice := strings.TrimSpace(selection[key])
		if choice == "" {
			choice = strings.TrimSpace(v.Defaults[key])
		}
		if choice == "" {
			continue
		}
		if cls, ok2 := choices[choice]; ok2 {
			if strings.TrimSpace(cls) != "" {
				parts = append(parts, cls)
			}
		} else {
			parts = append(parts, choice)
		}
	}
	parts = append(parts, extra...)
	return Cn(parts...)
}

// Cn joins non-empty class fragments with a single space.
func Cn(classes ...string) string {
	var parts []string
	for _, c := range classes {
		if t := strings.TrimSpace(c); t != "" {
			parts = append(parts, t)
		}
	}
	return strings.Join(parts, " ")
}

// MergeAttrs merges attribute maps left to right.
func MergeAttrs(parts ...templ.Attributes) templ.Attributes {
	out := templ.Attributes{}
	for _, attrs := range parts {
		for k, v := range attrs {
			out[k] = v
		}
	}
	return out
}

// DOMAttrs merges ID, role, tabindex, and custom attrs.
func DOMAttrs(id, role, tabIndex string, attrs templ.Attributes) templ.Attributes {
	out := templ.Attributes{}
	for k, v := range attrs {
		out[k] = v
	}
	if strings.TrimSpace(id) != "" {
		out["id"] = id
	}
	if strings.TrimSpace(role) != "" {
		out["role"] = role
	}
	if strings.TrimSpace(tabIndex) != "" {
		out["tabindex"] = tabIndex
	}
	return out
}

// ControlAttrs spreads field attrs without duplicating id on the element.
func ControlAttrs(id, role, tabIndex, ariaLabel string, attrs templ.Attributes) templ.Attributes {
	out := DOMAttrs(id, role, tabIndex, attrs)
	delete(out, "id")
	if strings.TrimSpace(ariaLabel) != "" {
		out["aria-label"] = ariaLabel
	}
	return out
}

// SwitchAttrs adds switch semantics on checkbox controls.
func SwitchAttrs(id, role, tabIndex, ariaLabel string, checked bool, attrs templ.Attributes) templ.Attributes {
	out := ControlAttrs(id, role, tabIndex, ariaLabel, attrs)
	out["role"] = "switch"
	out["aria-checked"] = strconv.FormatBool(checked)
	return out
}

// ARIA helpers for static components.
func AriaExpanded(v bool) templ.Attributes {
	return templ.Attributes{"aria-expanded": strconv.FormatBool(v)}
}

func AriaControls(id string) templ.Attributes {
	if strings.TrimSpace(id) == "" {
		return templ.Attributes{}
	}
	return templ.Attributes{"aria-controls": id}
}

func AriaCurrent(value string) templ.Attributes {
	if strings.TrimSpace(value) == "" {
		return templ.Attributes{}
	}
	return templ.Attributes{"aria-current": value}
}

func AriaLive(value string) templ.Attributes {
	if strings.TrimSpace(value) == "" {
		return templ.Attributes{}
	}
	return templ.Attributes{"aria-live": value}
}

func AriaModal(v bool) templ.Attributes {
	return templ.Attributes{"aria-modal": strconv.FormatBool(v)}
}

func AriaLabel(value string) templ.Attributes {
	if strings.TrimSpace(value) == "" {
		return templ.Attributes{}
	}
	return templ.Attributes{"aria-label": value}
}

func AriaPressed(pressed bool) templ.Attributes {
	return templ.Attributes{"aria-pressed": strconv.FormatBool(pressed)}
}

func AriaHasPopup(kind string) templ.Attributes {
	switch strings.TrimSpace(strings.ToLower(kind)) {
	case "true", "menu", "listbox", "tree", "grid", "dialog":
		return templ.Attributes{"aria-haspopup": strings.TrimSpace(strings.ToLower(kind))}
	default:
		return templ.Attributes{}
	}
}

// Form control class recipes (shared by input, textarea, select, checkbox, radio, switch).
var InputChrome = Variants{
	Base: "w-full rounded-md border px-3 py-2 text-sm outline-none transition-colors focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50",
	Keys: []string{"variant"},
	Defaults: map[string]string{"variant": ""},
	ByKey: map[string]map[string]string{
		"variant": {
			"unstyled": "",
			"":         "border-border bg-background",
			"default":  "border-border bg-background",
			"outline":  "border-border bg-background",
			"ghost":    "border-transparent bg-muted",
		},
	},
}

var InputSize = Variants{
	Keys:     []string{"size"},
	Defaults: map[string]string{"size": "default"},
	ByKey: map[string]map[string]string{
		"size": {
			"":        "h-10 px-3 py-2 text-sm",
			"default": "h-10 px-3 py-2 text-sm",
			"md":      "h-10 px-3 py-2 text-sm",
			"xs":      "h-8 px-2 text-xs",
			"sm":      "h-9 px-3 text-sm",
			"lg":      "h-11 px-4 text-base",
		},
	},
}

var ControlChrome = Variants{
	Base: "rounded border border-primary text-primary focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50",
	Keys: []string{"variant"},
	Defaults: map[string]string{"variant": ""},
	ByKey: map[string]map[string]string{
		"variant": {
			"unstyled": "",
			"":         "border-primary",
			"default":  "border-primary",
			"outline":  "border-primary",
			"ghost":    "border-primary bg-muted/30",
		},
	},
}

var ControlSize = Variants{
	Keys:     []string{"size"},
	Defaults: map[string]string{"size": "default"},
	ByKey: map[string]map[string]string{
		"size": {
			"":        "h-4 w-4",
			"default": "h-4 w-4",
			"md":      "h-4 w-4",
			"xs":      "h-3 w-3",
			"sm":      "h-3.5 w-3.5",
			"lg":      "h-5 w-5",
		},
	},
}

// InputClasses composes input/textarea/select classes.
func InputClasses(variant, size string, extra ...string) string {
	chrome := Compose(InputChrome, map[string]string{"variant": strings.TrimSpace(variant)})
	sz := Compose(InputSize, map[string]string{"size": strings.TrimSpace(size)})
	return Cn(append([]string{chrome, sz}, extra...)...)
}

// ControlClasses composes checkbox/radio/switch classes.
func ControlClasses(variant, size string, extra ...string) string {
	chrome := Compose(ControlChrome, map[string]string{"variant": strings.TrimSpace(variant)})
	sz := Compose(ControlSize, map[string]string{"size": strings.TrimSpace(size)})
	return Cn(append([]string{chrome, sz}, extra...)...)
}

// CardVariants is the card surface recipe.
var CardVariants = Variants{
	Base: "rounded-md border border-border bg-card text-card-foreground shadow-sm",
	Keys: []string{"variant"},
	Defaults: map[string]string{"variant": ""},
	ByKey: map[string]map[string]string{
		"variant": {
			"":        "",
			"default": "",
			"raised":  "",
			"kpi":     "bg-card/50",
			"muted":   "bg-muted/40",
			"ghost":   "border-dashed bg-background shadow-none",
			"compact": "p-3",
			"flat":    "shadow-none",
			"accent":  "border-transparent bg-accent/30",
		},
	},
}

// CardClasses merges card variant presets.
func CardClasses(variant, extra string) string {
	return Compose(CardVariants, map[string]string{"variant": strings.TrimSpace(variant)}, extra)
}

// AlertVariants is the alert surface recipe.
var AlertVariants = Variants{
	Base: "rounded-lg border px-4 py-3 shadow-sm outline-none ring-offset-background gap-6",
	Keys: []string{"variant"},
	Defaults: map[string]string{"variant": ""},
	ByKey: map[string]map[string]string{
		"variant": {
			"":            "bg-card text-card-foreground",
			"default":     "bg-card text-card-foreground",
			"destructive": "border-destructive bg-destructive/10 text-destructive",
			"success":     "border-primary bg-primary/10 text-primary",
			"warning":     "border-border bg-accent text-accent-foreground",
		},
	},
}

// AlertClasses merges alert variant presets.
func AlertClasses(variant, extra string) string {
	return Compose(AlertVariants, map[string]string{"variant": strings.TrimSpace(variant)}, extra)
}

// AlertAttrs returns static alert semantics.
func AlertAttrs() templ.Attributes {
	return MergeAttrs(
		templ.Attributes{"role": "status"},
		AriaLive("polite"),
	)
}

// BreadcrumbRootAttrs returns nav root attrs.
func BreadcrumbRootAttrs() templ.Attributes {
	return templ.Attributes{
		"data-ui8kit": "breadcrumb",
		"aria-label":  "Breadcrumb",
	}
}

// BreadcrumbItemAttrs returns per-item attrs.
func BreadcrumbItemAttrs(current bool) templ.Attributes {
	if current {
		return AriaCurrent("page")
	}
	return templ.Attributes{}
}

// TextareaRows returns a safe row count.
func TextareaRows(rows int) int {
	if rows <= 0 {
		return 4
	}
	return rows
}

// DefaultInputType returns text when empty.
func DefaultInputType(t string) string {
	if strings.TrimSpace(t) == "" {
		return "text"
	}
	return t
}

// DefaultButtonType returns button when empty.
func DefaultButtonType(t string) string {
	if strings.TrimSpace(t) == "" {
		return "button"
	}
	return t
}

// TitleTag returns h1-h6 from order.
func TitleTag(order int) string {
	switch order {
	case 1:
		return "h1"
	case 3:
		return "h3"
	case 4:
		return "h4"
	case 5:
		return "h5"
	case 6:
		return "h6"
	default:
		return "h2"
	}
}
