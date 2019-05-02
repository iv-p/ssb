package site

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

// Context holds data about the resolved site context
type Context struct {
	SiteID string
}

// Resolver extracts the site id from the request url
type Resolver struct {
	// Map of site domain name to site id
	domainMapLoader IDomainMapLoader
}

// NewResolver creates a new site id resolver
func NewResolver(domainMapLoader IDomainMapLoader) *Resolver {
	return &Resolver{
		domainMapLoader: domainMapLoader,
	}
}

// Resolve takes a request url and finds the appropriate
// site id it is for
func (r *Resolver) Resolve(u *url.URL) (Context, error) {
	domainMap := r.domainMapLoader.Load()
	host := u.Host
	var err error
	if strings.Index(host, ":") != -1 {
		host, _, err = net.SplitHostPort(u.Host)
		if err != nil {
			return Context{}, err
		}
	}
	siteID, ok := domainMap[host]
	if !ok {
		return Context{}, fmt.Errorf("site id not found for request %s with host %s", u.String(), host)
	}
	return Context{SiteID: siteID}, nil
}
