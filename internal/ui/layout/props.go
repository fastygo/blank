package layout

import (
	"github.com/a-h/templ"
	"github.com/fastygo/blank/internal/ui/components/navigation"
)

const (
	MobileSheetTriggerID = navigation.MobileSheetTriggerID
	MobileSheetPanelID   = navigation.MobileSheetPanelID
)

type NavItem = navigation.Item
type NavProps = navigation.NavProps

// HeaderProps configures the top header bar.
type HeaderProps struct {
	BrandName       string
	NavItems        []NavItem
	Active          string
	ShowMenuTrigger bool
	Navigation      NavigationProps
	Trailing        templ.Component
	ThemeToggle     ThemeToggleProps
}

// ThemeToggleProps configures copy for the theme toggle button.
type ThemeToggleProps struct {
	Label              string
	SwitchToDarkLabel  string
	SwitchToLightLabel string
}

// NavigationProps configures shell navigation accessible names.
type NavigationProps struct {
	BrandHomeSuffix     string
	MainNavigation      string
	OpenNavigationMenu  string
	CloseNavigationMenu string
	NavigationMenuLabel string
}

// ShellProps configures the full page shell (header + main + footer).
type ShellProps struct {
	Title          string
	Lang           string
	BrandName      string
	Active         string
	NavItems       []NavItem
	FooterText     string
	Navigation     NavigationProps
	HeadExtra      templ.Component
	HeaderTrailing templ.Component
	ThemeToggle    ThemeToggleProps
}
