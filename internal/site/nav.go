package site

import (
	"github.com/fastygo/blank/internal/fixtures"
	"github.com/fastygo/blank/internal/ui/layout"
)

// siteNav builds header nav items from runtime route specs.
func (f *Feature) siteNav(fix fixtures.Locale) []layout.NavItem {
	var items []layout.NavItem
	for _, page := range pages {
		if page.Nav == nil {
			continue
		}
		item, ok := page.Nav(fix)
		if !ok {
			continue
		}
		items = append(items, item)
	}
	return items
}
