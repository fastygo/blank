package layout

import (
	"github.com/a-h/templ"
	"github.com/fastygo/blank/internal/ui/components/appsidebar"
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

// DocumentProps configures the document frame (html, head, body).
type DocumentProps struct {
	Title     string
	Lang      string
	BodyClass string
	HeadExtra templ.Component
}

// TopnavLayoutProps configures the topnav app chrome (header, main, footer, mobile sheet).
type TopnavLayoutProps struct {
	Brand           string
	Active          string
	NavItems        []NavItem
	Navigation      NavigationProps
	Theme           ThemeToggleProps
	Trailing        templ.Component
	FooterText      string
	ShowMobileSheet bool
}

// DashboardLayoutProps configures the dashboard app chrome (topnav + desktop aside).
type DashboardLayoutProps struct {
	Topnav  TopnavLayoutProps
	Sidebar appsidebar.Props
}
