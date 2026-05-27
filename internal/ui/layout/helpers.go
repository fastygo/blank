package layout

import (
	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
	uiutils "github.com/fastygo/templ/utils"
)

const (
	MobileSheetTriggerID = "ui8kit-mobile-sheet-trigger"
	MobileSheetPanelID   = "ui8kit-mobile-sheet-panel"
)

func shellBrand(name string) string {
	if name == "" {
		return "App"
	}
	return name
}

func shellLang(value string) string {
	if value == "" {
		return "en"
	}
	return value
}

func shellBodyClass(_ ShellProps) string {
	return uiutils.Cn(
		"min-h-screen overflow-x-hidden bg-background font-sans text-foreground",
		"max-md:has-[#ui8kit-mobile-sheet-panel:not([hidden])]:overflow-hidden",
	)
}

func shellHasNavigation(props ShellProps) bool {
	return len(props.NavItems) > 0
}

func themeToggleLabel(value string) string {
	if value == "" {
		return "Toggle theme"
	}
	return value
}

func themeToggleSwitchToDarkLabel(value string) string {
	if value == "" {
		return "Switch to dark theme"
	}
	return value
}

func themeToggleSwitchToLightLabel(value string) string {
	if value == "" {
		return "Switch to light theme"
	}
	return value
}

func navItemClasses(active, path string, vertical bool) string {
	if vertical {
		base := "flex w-full items-center gap-2 rounded-md px-4 py-2 text-sm"
		if active == path {
			return uiutils.Cn(base, "bg-accent text-accent-foreground")
		}
		return uiutils.Cn(base, "text-muted-foreground hover:bg-accent")
	}
	base := "inline-flex items-center rounded-md px-3 py-2 text-sm font-medium transition-colors"
	if active == path {
		return uiutils.Cn(base, "text-foreground")
	}
	return uiutils.Cn(base, "text-muted-foreground hover:text-foreground")
}

func brandLogoButtonProps(brandName string) ui.ButtonProps {
	name := shellBrand(brandName)
	return ui.ButtonProps{
		Href:    "/",
		Variant: "unstyled",
		Class:   "text-base font-semibold tracking-tight text-foreground transition-colors hover:text-foreground/80",
		Attrs: uiutils.MergeAttrs(
			templ.Attributes{"aria-label": name + " home"},
		),
	}
}

func headerMenuButtonProps() ui.ButtonProps {
	return ui.ButtonProps{
		ID:      MobileSheetTriggerID,
		Type:    "button",
		Variant: "unstyled",
		Class: uiutils.Cn(
			"inline-flex h-8 w-8 shrink-0 items-center justify-center rounded-md border border-border text-foreground md:hidden",
		),
		Attrs: uiutils.MergeAttrs(
			templ.Attributes{
				"data-ui8kit-dialog-open":   true,
				"data-ui8kit-dialog-target": MobileSheetPanelID,
			},
			uiutils.AriaLabel("Open navigation menu"),
			uiutils.AriaHasPopup("dialog"),
			uiutils.AriaControls(MobileSheetPanelID),
			uiutils.AriaExpanded(false),
		),
	}
}

func themeToggleButtonProps(props ThemeToggleProps) ui.ButtonProps {
	return ui.ButtonProps{
		ID:      "ui8kit-theme-toggle",
		Type:    "button",
		Variant: "unstyled",
		Class: uiutils.Cn(
			"inline-flex h-8 w-8 items-center justify-center rounded-md bg-transparent p-0 text-muted-foreground transition-colors hover:bg-accent hover:text-accent-foreground",
		),
		Attrs: uiutils.MergeAttrs(
			templ.Attributes{
				"data-switch-to-dark-label":  themeToggleSwitchToDarkLabel(props.SwitchToDarkLabel),
				"data-switch-to-light-label": themeToggleSwitchToLightLabel(props.SwitchToLightLabel),
				"title":                      themeToggleLabel(props.Label),
			},
			uiutils.AriaLabel(themeToggleLabel(props.Label)),
			uiutils.AriaPressed(false),
		),
	}
}

func mobileSheetRootBlock() ui.BlockProps {
	return ui.BlockProps{
		ID:    MobileSheetPanelID,
		Class: "fixed inset-y-0 left-0 z-50 w-full md:hidden",
		Attrs: templ.Attributes{
			"data-ui8kit":        "sheet",
			"data-ui8kit-dialog": true,
			"role":               "dialog",
			"aria-modal":         "true",
			"aria-label":         "Navigation menu",
			"aria-labelledby":    "ui8kit-mobile-sheet-title",
			"data-state":         "closed",
			"hidden":             true,
		},
	}
}

func mobileSheetOverlayBox() ui.BoxProps {
	return ui.BoxProps{
		Class: "absolute inset-0 cursor-pointer bg-card/50",
		Attrs: templ.Attributes{
			"data-ui8kit-dialog-overlay": true,
			"data-ui8kit-dialog-close":   true,
			"data-ui8kit-dialog-target":  MobileSheetPanelID,
		},
	}
}

func mobileSheetCloseButtonProps() ui.ButtonProps {
	return ui.ButtonProps{
		Type:      "button",
		Variant:   "unstyled",
		Class:     "inline-flex h-8 w-8 cursor-pointer items-center justify-center rounded-md border border-border text-muted-foreground",
		AriaLabel: "Close navigation menu",
		Attrs: templ.Attributes{
			"data-ui8kit-dialog-close":  true,
			"data-ui8kit-dialog-target": MobileSheetPanelID,
		},
	}
}

func navLinkButtonProps(active string, item NavItem, vertical bool) ui.ButtonProps {
	attrs := templ.Attributes{}
	if active == item.Path {
		attrs = uiutils.MergeAttrs(attrs, uiutils.AriaCurrent("page"))
	}
	return ui.ButtonProps{
		Href:    item.Path,
		Variant: "unstyled",
		Class:   navItemClasses(active, item.Path, vertical),
		Attrs:   attrs,
	}
}
