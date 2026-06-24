package sidebarapp

import (
	"strings"

	"github.com/a-h/templ"
	"github.com/fastygo/blank/internal/ui/components/navigation"
	uiutils "github.com/fastygo/templ/utils"
)

func (p Props) sidebarItems() []navigation.Item {
	if len(p.Sidebar.Items) > 0 {
		return p.Sidebar.Items
	}
	return p.Shell.NavItems
}

func (p Props) sidebarActive() string {
	if p.Sidebar.Active != "" {
		return p.Sidebar.Active
	}
	return p.Shell.Active
}

func (p Props) sidebarTitle() string {
	if strings.TrimSpace(p.Sidebar.Title) != "" {
		return p.Sidebar.Title
	}
	if strings.TrimSpace(p.Shell.BrandName) != "" {
		return p.Shell.BrandName
	}
	return "App"
}

func (p Props) sidebarAriaLabel() string {
	if strings.TrimSpace(p.Sidebar.AriaLabel) != "" {
		return p.Sidebar.AriaLabel
	}
	return p.Shell.Navigation.MainNavigation
}

func sidebarAsideClass(extra string) string {
	return uiutils.Cn(
		"hidden shrink-0 border-r border-border bg-background md:flex md:w-full md:max-w-sm md:flex-col",
		extra,
	)
}

func sidebarAsideAttrs(p Props) templ.Attributes {
	return uiutils.MergeAttrs(
		templ.Attributes{"aria-label": p.sidebarAriaLabel()},
	)
}

func sidebarNavClass() string {
	return "flex w-full flex-col gap-1 p-2"
}
