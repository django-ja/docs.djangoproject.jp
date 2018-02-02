package redirect

import (
	"testing"
)

func TestMapUrl(t *testing.T) {
	m := map[string]string{
		"":                         "https://docs.djangoproject.com/ja/2.0/",
		"/":                        "https://docs.djangoproject.com/ja/2.0/",
		"/favicon.ico":             "https://docs.djangoproject.com/ja/2.0/",
		"/en/latest/faq/":          "https://docs.djangoproject.com/ja/2.0/faq/",
		"/ja/latest/faq/":          "https://docs.djangoproject.com/ja/2.0/",
		"/en/latest/contents.html": "https://docs.djangoproject.com/ja/2.0/contents/",
	}

	for k, v := range m {
		if r := mapUrl(k); r != v {
			t.Errorf("mapUrl(%q) is expected: %q, but actually %q", k, v, r)
		}
	}
}
