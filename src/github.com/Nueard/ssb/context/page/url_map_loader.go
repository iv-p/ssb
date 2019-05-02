package page

import "github.com/Nueard/ssb/context/site"

// URLMap holds the url -> page id for a site
type URLMap map[string]string

// DummyData holds a dummy map for local development
var DummyData = URLMap{
	"/":              "1",
	"/something/one": "2",
	"/another":       "3",
}

// IURLMapLoader is the url page loader interface
type IURLMapLoader interface {
	Load(site.Context) URLMap
}

// URLMapLoader resolves page id based on the url request and site context
type URLMapLoader struct {
}

// NewURLMapLoader creates and returns a new URLMapLoader object
func NewURLMapLoader() *URLMapLoader {
	return &URLMapLoader{}
}

// Load takes a site context and loads all the urls for that site in a url map
func (upl *URLMapLoader) Load(context site.Context) URLMap {
	return DummyData
}
