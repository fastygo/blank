package views

import (
	"github.com/a-h/templ"
	"github.com/fastygo/blank/internal/fixtures"
	"github.com/fastygo/blank/internal/ui/layout"
)

// SamplePageFrom maps request data and locale copy into the sample page component.
func SamplePageFrom(d layout.Data, f fixtures.Locale) templ.Component {
	return SamplePage(
		d.ShellProps(),
		d.SidebarProps(f.Sample.Title),
		SampleContent{
			Title:       f.Sample.Title,
			Description: f.Sample.Description,
			Body:        f.Sample.Body,
		},
	)
}

// SampleContent is the page body copy for /sample.
type SampleContent struct {
	Title       string
	Description string
	Body        string
}
