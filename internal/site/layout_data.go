package site

import (
	"context"
	"net/http"

	"github.com/fastygo/blank/internal/fixtures"
	"github.com/fastygo/blank/internal/ui/layout"
	"github.com/fastygo/framework/pkg/web/locale"
)

func (f *Feature) layoutData(ctx context.Context, r *http.Request, fix fixtures.Locale, title, active string) layout.Data {
	return layout.BuildData(layout.BuildParams{
		PageTitle:        title,
		Active:           active,
		Lang:             locale.From(ctx),
		Brand:            fix.Brand,
		FooterText:       fix.Footer,
		Fix:              fix,
		NavItems:         f.siteNav(fix),
		Request:          r,
		AvailableLocales: f.available,
		DefaultLocale:    f.defaultLocale,
	})
}
