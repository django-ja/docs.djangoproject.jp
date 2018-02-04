package redirect

import (
	"testing"
)

func TestMapUrlSimple(t *testing.T) {
	m := map[string]string{
		"":                                         "https://docs.djangoproject.com/ja/2.0/",
		"/":                                        "https://docs.djangoproject.com/ja/2.0/",
		"/favicon.ico":                             "https://docs.djangoproject.com/ja/2.0/",
		"/en/latest/faq/":                          "https://docs.djangoproject.com/ja/2.0/faq/",
		"/ja/latest/faq/":                          "https://docs.djangoproject.com/ja/2.0/",
		"/en/latest/contents.html":                 "https://docs.djangoproject.com/ja/2.0/contents/",
		"/en/latest/_sources/ref/forms/fields.txt": "https://docs.djangoproject.com/ja/2.0/ref/forms/fields/",
	}

	for k, v := range m {
		if r := mapUrl(k); r != v {
			t.Errorf("mapUrl(%q) is expected: %q, but actually %q", k, v, r)
		}
	}
}

func TestMapUrlMigration(t *testing.T) {
	m := map[string]string{
		"/en/latest/howto/apache-auth.html":                  "https://docs.djangoproject.com/ja/2.0/howto/deployment/wsgi/apache-auth/",
		"/en/latest/howto/deployment/fastcgi.html":           "https://docs.djangoproject.com/ja/2.0/internals/deprecation/#deprecation-removed-in-1-9",
		"/en/latest/howto/deployment/modpython.html":         "https://docs.djangoproject.com/ja/2.0/internals/deprecation/#deprecation-removed-in-1-5",
		"/en/latest/internals/committers.html":               "https://docs.djangoproject.com/ja/2.0/internals/organization/#prerogatives",
		"/en/latest/internals/documentation.html":            "https://docs.djangoproject.com/ja/2.0/internals/contributing/writing-documentation/",
		"/en/latest/obsolete/":                               "https://docs.djangoproject.com/ja/2.0/internals/deprecation/",
		"/en/latest/obsolete/admin-css.html":                 "https://docs.djangoproject.com/ja/2.0/internals/deprecation/",
		"/en/latest/ref/contrib/comments/":                   "https://docs.djangoproject.com/ja/2.0/releases/1.8/#features-removed-in-1-8",
		"/en/latest/ref/contrib/comments/custom.html":        "https://docs.djangoproject.com/ja/2.0/releases/1.8/#features-removed-in-1-8",
		"/en/latest/ref/contrib/csrf.html":                   "https://docs.djangoproject.com/ja/2.0/ref/csrf/",
		"/en/latest/ref/contrib/databrowse.html":             "https://docs.djangoproject.com/ja/2.0/releases/1.4/#django-contrib-databrowse",
		"/en/latest/ref/contrib/formtools/":                  "https://docs.djangoproject.com/ja/2.0/releases/1.8/#removal-of-django-contrib-formtools",
		"/en/latest/ref/contrib/formtools/form-preview.html": "https://docs.djangoproject.com/ja/2.0/releases/1.8/#removal-of-django-contrib-formtools",
		"/en/latest/ref/contrib/localflavor.html":            "https://docs.djangoproject.com/ja/2.0/internals/deprecation/#deprecation-removed-in-1-6",
		"/en/latest/ref/contrib/webdesign.html":              "https://docs.djangoproject.com/ja/2.0/releases/1.8/#django-contrib-webdesign",
		"/en/latest/ref/generic-views.html":                  "https://docs.djangoproject.com/ja/2.0/topics/class-based-views/",
		"/en/latest/topics/generic-views-migration.html":     "https://docs.djangoproject.com/ja/2.0/topics/class-based-views/",
		"/en/latest/topics/http/generic-views.html":          "https://docs.djangoproject.com/ja/2.0/topics/class-based-views/",
		"/en/latest/topics/generic-views.html":               "https://docs.djangoproject.com/ja/2.0/topics/class-based-views/",
		"/en/latest/releases/1.0-alpha-1.html":               "https://docs.djangoproject.com/ja/2.0/releases/1.0/",
		"/en/latest/releases/1.0-alpha-2.html":               "https://docs.djangoproject.com/ja/2.0/releases/1.0/",
		"/en/latest/releases/1.0-beta.html":                  "https://docs.djangoproject.com/ja/2.0/releases/1.0/",
		"/en/latest/releases/1.1-alpha-1.html":               "https://docs.djangoproject.com/ja/2.0/releases/1.1/",
	}

	for k, v := range m {
		if r := mapUrl(k); r != v {
			t.Errorf("mapUrl(%q) is expected with migration: %q, but actually %q", k, v, r)
		}
	}
}
