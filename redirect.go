package main

import (
	"net/http"
	"regexp"
	"strings"
)

const (
	base      = "https://docs.djangoproject.com/ja/2.2/"
	prefix    = "/en/latest/"
	suffix    = ".html"
	srcPrefix = "_sources/"
	srcSuffix = ".txt"
)

func migrateUrl(s string) string {
	m := map[string]string{
		`howto/apache-auth/`:          "howto/deployment/wsgi/apache-auth/",
		`howto/deployment/fastcgi/`:   "internals/deprecation/#deprecation-removed-in-1-9",
		`howto/deployment/modpython/`: "internals/deprecation/#deprecation-removed-in-1-5",
		`internals/committers/`:       "internals/organization/#prerogatives",
		`internals/documentation/`:    "internals/contributing/writing-documentation/",
		`obsolete/.*`:                 "internals/deprecation/",
		`ref/authbackends/`:           "topics/auth/customizing/#authentication-backends",
		`ref/contrib/comments/.*`:     "releases/1.8/#features-removed-in-1-8",
		`ref/contrib/csrf/`:           "ref/csrf/",
		`ref/contrib/databrowse/`:     "releases/1.4/#django-contrib-databrowse",
		`ref/contrib/formtools/.*`:    "releases/1.8/#removal-of-django-contrib-formtools",
		`ref/contrib/localflavor/`:    "internals/deprecation/#deprecation-removed-in-1-6",
		`ref/contrib/webdesign/`:      "releases/1.8/#django-contrib-webdesign",
		`.*/generic-views.*/`:         "topics/class-based-views/",
		`releases/1.0-.*/`:            "releases/1.0/",
		`releases/1.1-.*/`:            "releases/1.1/",
	}

	for k, v := range m {
		if regexp.MustCompile(k).MatchString(s) {
			return v
		}
	}

	return s
}

func mapUrl(s string) string {
	if !strings.HasPrefix(s, prefix) {
		return base
	}

	// Remove .html and .txt
	for _, _suffix := range []string{suffix, srcSuffix} {
		if strings.HasSuffix(s, _suffix) {
			s = strings.TrimSuffix(s, _suffix) + "/"
		}
	}

	// Remove /en/latest and _sources
	for _, _prefix := range []string{prefix, srcPrefix} {
		if strings.HasPrefix(s, _prefix) {
			s = strings.TrimPrefix(s, _prefix)
		}
	}

	s = migrateUrl(s)

	return base + s
}

func redirect(w http.ResponseWriter, r *http.Request) {
	var s = mapUrl(r.URL.Path)

	http.Redirect(w, r, s, 301)
}

func main() {
	http.HandleFunc("/", redirect)
}
