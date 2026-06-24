package sidebarapp

import (
	"github.com/fastygo/blank/internal/ui/components/navigation"
	"github.com/fastygo/blank/internal/ui/layout"
)

// Props configures the sidebar app shell registry artifact (document chrome + sidebar + main slot).
type Props struct {
	Shell   layout.ShellProps
	Sidebar SidebarProps
}

// SidebarProps configures the desktop sidebar navigation region.
type SidebarProps struct {
	Title     string
	AriaLabel string
	Items     []navigation.Item
	Active    string
	Class     string
}
