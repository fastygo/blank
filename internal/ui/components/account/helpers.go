package account

import (
	"strings"

	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
	uiutils "github.com/fastygo/templ/utils"
)

func menuTriggerButtonProps(label string) ui.ButtonProps {
	return ui.ButtonProps{
		Type:    "button",
		Variant: "unstyled",
		Class: uiutils.Cn(
			"inline-flex h-8 w-8 shrink-0 items-center justify-center rounded-md bg-transparent p-0 text-muted-foreground transition-colors hover:bg-accent hover:text-accent-foreground",
		),
		Attrs: uiutils.MergeAttrs(
			templ.Attributes{
				"aria-haspopup": "menu",
			},
			uiutils.AriaLabel(menuTriggerAriaLabel(label)),
		),
	}
}

func menuTriggerAriaLabel(value string) string {
	if strings.TrimSpace(value) != "" {
		return value
	}
	return "Account menu"
}

func menuSignOutAction(value string) string {
	if strings.TrimSpace(value) != "" {
		return value
	}
	return "/auth/logout"
}

func menuSignOutLabel(value string) string {
	if strings.TrimSpace(value) != "" {
		return value
	}
	return "Sign out"
}

func menuProfileLabel(value string) string {
	if strings.TrimSpace(value) != "" {
		return value
	}
	return "Profile"
}
