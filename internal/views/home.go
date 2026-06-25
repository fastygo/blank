package views

import (
	"github.com/a-h/templ"
	"github.com/fastygo/blank/internal/fixtures"
	"github.com/fastygo/blank/internal/ui/blocks/marketing/hero"
	"github.com/fastygo/blank/internal/ui/layout"
)

// HomePageFrom maps request data and locale copy into the home page component.
func HomePageFrom(d layout.Data, f fixtures.Locale) templ.Component {
	return HomePage(d.ShellProps(), hero.Props{
		Welcome:      f.Home.Welcome,
		WelcomeBrand: f.Home.WelcomeBrand,
		Description:  f.Home.Description,
	})
}
