package layout

import (
	"strings"

	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
	uiutils "github.com/fastygo/templ/utils"
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
		"max-md:has-[#"+MobileSheetPanelID+":not([hidden])]:overflow-hidden",
	)
}

func shellHasNavigation(props ShellProps) bool {
	return len(props.NavItems) > 0
}

func themeToggleLabel(value string) string {
	return strings.TrimSpace(value)
}

func themeToggleSwitchToDarkLabel(value string) string {
	return strings.TrimSpace(value)
}

func themeToggleSwitchToLightLabel(value string) string {
	return strings.TrimSpace(value)
}

func brandLogoButtonProps(brandName string, nav NavigationProps) ui.ButtonProps {
	name := shellBrand(brandName)
	return ui.ButtonProps{
		Href:    "/",
		Variant: "unstyled",
		Class:   "text-base font-semibold tracking-tight text-foreground transition-colors hover:text-foreground/80",
		Attrs: uiutils.MergeAttrs(
			templ.Attributes{"aria-label": name + nav.BrandHomeSuffix},
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
