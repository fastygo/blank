package sidebarapp

import "github.com/fastygo/blank/internal/ui/layout"

// DefaultProps returns wireframe demo props for tests and showcase renders.
func DefaultProps() Props {
	return Props{
		Shell: layout.ShellProps{
			Title:      "Dashboard · FastyGo",
			Lang:       "en",
			BrandName:  "FastyGo",
			Active:     "/",
			FooterText: "Made with 🤍 in FastyGo",
			NavItems: []layout.NavItem{
				{Label: "Home", Path: "/", Icon: "home"},
				{Label: "Sample", Path: "/sample", Icon: "file"},
			},
			Navigation: layout.NavigationProps{
				BrandHomeSuffix:     " home",
				MainNavigation:      "Main navigation",
				OpenNavigationMenu:  "Open navigation menu",
				CloseNavigationMenu: "Close navigation menu",
				NavigationMenuLabel: "Navigation menu",
			},
			ThemeToggle: layout.ThemeToggleProps{
				Label:              "Theme",
				SwitchToDarkLabel:  "Switch to dark",
				SwitchToLightLabel: "Switch to light",
			},
		},
		Sidebar: SidebarProps{
			Title:     "FastyGo",
			AriaLabel: "App navigation",
		},
	}
}
