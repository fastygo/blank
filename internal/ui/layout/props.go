package layout

import "github.com/a-h/templ"

// NavItem describes a single sidebar navigation link.
type NavItem struct {
	Path  string
	Label string
	Icon  string
}

// SidebarProps configures the sidebar navigation.
type SidebarProps struct {
	Items  []NavItem
	Active string
}

// HeaderProps configures the top header bar.
type HeaderProps struct {
	ShowMenuTrigger bool
	Title           string
	Trailing        templ.Component
	ThemeToggle     ThemeToggleProps
}

// ThemeToggleProps configures copy for the theme toggle button.
type ThemeToggleProps struct {
	Label              string
	SwitchToDarkLabel  string
	SwitchToLightLabel string
}

// ShellProps configures the full page shell (sidebar + header + main).
type ShellProps struct {
	Title          string
	HeaderTitle    string
	Lang           string
	BrandName      string
	Active         string
	NavItems       []NavItem
	HeadExtra      templ.Component
	HeaderTrailing templ.Component
	ThemeToggle    ThemeToggleProps
}
