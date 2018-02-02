package redirect

import (
	"net/http"
	"strings"
)

const (
	domain = "https://docs.djangoproject.com/ja/2.0/"
	prefix = "/en/latest/"
	suffix = ".html"
)

func mapUrl(s string) (r string) {
	if strings.HasSuffix(s, suffix) {
		s = strings.TrimSuffix(s, suffix) + "/"
	}

	r = domain
	if strings.HasPrefix(s, prefix) {
		r += strings.TrimPrefix(s, prefix)
	}

	return r
}

func redirect(w http.ResponseWriter, r *http.Request) {
	var s = mapUrl(r.URL.Path)

	http.Redirect(w, r, s, 301)
}

func init() {
	http.HandleFunc("/", redirect)
}
