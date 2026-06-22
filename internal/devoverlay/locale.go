package devoverlay

import (
	"net/http"

	"github.com/fastygo/framework/pkg/web/locale"
)

const defaultLangCookieName = "lang"

// resolveLocale mirrors site negotiation because devoverlay.Wrap sits outside
// app locale middleware and the outer *http.Request context is never updated.
func (c Config) resolveLocale(r *http.Request) string {
	defaultLocale := c.DefaultLocale
	if defaultLocale == "" {
		defaultLocale = "en"
	}
	available := c.AvailableLocales
	if len(available) == 0 {
		available = []string{"en", "ru"}
	}

	n := locale.New(defaultLocale, available)
	n.CookieName = c.LangCookieName
	if n.CookieName == "" {
		n.CookieName = defaultLangCookieName
	}
	return n.Resolve(r)
}
