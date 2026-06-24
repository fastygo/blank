package site

import (
	"context"
	"net/http"

	"github.com/fastygo/blank/internal/fixtures"
	"github.com/fastygo/framework/pkg/web"
	"github.com/fastygo/framework/pkg/web/locale"
)

func (f *Feature) fixtureLocale(ctx context.Context) fixtures.Locale {
	return fixtures.Resolve(locale.From(ctx), f.defaultLocale)
}

func (f *Feature) handlePage(page PageSpec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		fix := f.fixtureLocale(ctx)
		data := f.layoutData(ctx, r, fix, page.Title(fix), page.Active)
		_ = web.Render(ctx, w, page.Layout(data, page.Body(fix)))
	}
}
