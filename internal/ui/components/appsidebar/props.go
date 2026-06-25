// Package appsidebar provides the local sidebar component used by route pages that
// compose layout.DashboardLayout. It is the templ analogue of shadcn's
// components/app-sidebar.tsx: a small, props-only, copy-paste-friendly aside.
package appsidebar

import "github.com/fastygo/blank/internal/ui/components/navigation"

// Props configures the local sidebar aside (title + vertical navigation).
type Props struct {
	Title     string
	AriaLabel string
	Items     []navigation.Item
	Active    string
	Class     string
}
