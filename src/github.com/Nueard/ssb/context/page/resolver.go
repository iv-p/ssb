package page

import (
	"fmt"
	"net/url"

	"github.com/Nueard/ssb/context/site"
)

// Context is the resolved page context
type Context struct {
	PageID string
}

// Resolver resolves the page context based on the request url
type Resolver struct {
	urlMapLoader IURLMapLoader
}

// NewResolver creates and returns a new page resolver object
func NewResolver(urlMapLoader IURLMapLoader) *Resolver {
	return &Resolver{
		urlMapLoader: urlMapLoader,
	}
}

// Resolve takes a request url and finds the appropriate
// site id it is for
func (r *Resolver) Resolve(context site.Context, u *url.URL) (Context, error) {
	domainMap := r.urlMapLoader.Load(context)
	pageID, ok := domainMap[u.Path]
	if !ok {
		return Context{}, fmt.Errorf("page id not found for request %s with path %s", u.String(), u.Path)
	}
	return Context{PageID: pageID}, nil
}
