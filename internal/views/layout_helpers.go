package views

import (
	"github.com/fastygo/blank/internal/ui/components/appsidebar"
	"github.com/fastygo/blank/internal/ui/layout"
	"github.com/fastygo/blank/internal/views/partials"
)

// ShellPropsFor maps LayoutData into the document-shell props consumed by
// layout.Shell and layout.SidebarShell. Pages compose their own shell, so this
// is the one place that knows how request-derived data becomes shell props.
func ShellPropsFor(d LayoutData) layout.ShellProps {
	return layout.ShellProps{
		Title:          d.DocumentTitle(),
		Lang:           d.Lang,
		BrandName:      d.Brand,
		Active:         d.Active,
		NavItems:       d.NavItems,
		FooterText:     d.FooterText,
		Navigation:     d.Navigation,
		HeadExtra:      partials.ShellHead(d.Assets.CSS, d.Assets.ThemeJS, d.Assets.AppJS),
		HeaderTrailing: partials.HeaderTrailing(d.LanguageSwitch),
		ThemeToggle:    d.Theme,
	}
}

// SidebarPropsFor builds appsidebar.Props from LayoutData.
// Pass the title to display above the vertical nav (typically brand or section name);
// when empty, the sidebar falls back to "App".
func SidebarPropsFor(d LayoutData, title string) appsidebar.Props {
	resolvedTitle := title
	if resolvedTitle == "" {
		resolvedTitle = d.Brand
	}
	return appsidebar.Props{
		Title:     resolvedTitle,
		AriaLabel: d.Navigation.MainNavigation,
		Items:     d.NavItems,
		Active:    d.Active,
	}
}
