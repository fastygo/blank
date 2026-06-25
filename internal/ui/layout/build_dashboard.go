package layout

import (
	"strings"

	"github.com/fastygo/blank/internal/ui/components/appsidebar"
)

// sidebarProps builds appsidebar.Props from request data.
// Pass the title to display above the vertical nav (typically brand or section name);
// when empty, the sidebar falls back to the brand name or "App".
func sidebarProps(d Data, title string) appsidebar.Props {
	resolvedTitle := strings.TrimSpace(title)
	if resolvedTitle == "" {
		resolvedTitle = strings.TrimSpace(d.Brand)
	}
	if resolvedTitle == "" {
		resolvedTitle = "App"
	}
	return appsidebar.Props{
		Title:     resolvedTitle,
		AriaLabel: d.Navigation.MainNavigation,
		Items:     d.NavItems,
		Active:    d.Active,
	}
}
