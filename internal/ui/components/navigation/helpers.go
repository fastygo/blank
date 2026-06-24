package navigation

import (
	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
	uiutils "github.com/fastygo/templ/utils"
)

func (p MobileSheetProps) panelID() string {
	if p.ID != "" {
		return p.ID
	}
	return MobileSheetPanelID
}

func (p MobileSheetProps) titleID() string {
	if p.TitleID != "" {
		return p.TitleID
	}
	return MobileSheetTitleID
}

func (p MobileSheetTriggerProps) triggerID() string {
	if p.ID != "" {
		return p.ID
	}
	return MobileSheetTriggerID
}

func (p MobileSheetTriggerProps) panelID() string {
	if p.For != "" {
		return p.For
	}
	return MobileSheetPanelID
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

func navLinkButtonProps(active string, item Item, vertical bool) ui.ButtonProps {
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

func mobileSheetTriggerClass(extra string) string {
	return uiutils.Cn(
		"inline-flex h-8 w-8 shrink-0 items-center justify-center rounded-md border border-border text-foreground md:hidden",
		extra,
	)
}

func mobileSheetPanelClass() string {
	return "md:hidden max-h-dvh overflow-y-auto p-0 w-full"
}
