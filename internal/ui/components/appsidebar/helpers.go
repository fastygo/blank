package appsidebar

import (
	"strings"

	"github.com/a-h/templ"
	uiutils "github.com/fastygo/blank/internal/kit/utils"
)

func (p Props) title() string {
	if t := strings.TrimSpace(p.Title); t != "" {
		return t
	}
	return "App"
}

func (p Props) ariaLabel() string {
	if l := strings.TrimSpace(p.AriaLabel); l != "" {
		return l
	}
	return "Sidebar navigation"
}

func asideClass(extra string) string {
	return uiutils.Cn(
		"hidden shrink-0 border-r border-border bg-background md:flex md:w-full md:max-w-sm md:flex-col",
		extra,
	)
}

func asideAttrs(p Props) templ.Attributes {
	return uiutils.MergeAttrs(
		templ.Attributes{"aria-label": p.ariaLabel()},
	)
}

func navClass() string {
	return "flex w-full flex-col gap-1 p-2"
}
