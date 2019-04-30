package site

import (
	"fmt"
	"net/url"
)

// Resolver extracts the site id from the request url
type Resolver struct {
	// Map of site domain name to site id
	sites map[string]string
}

// NewResolver creates a new site id resolver
func NewResolver(sites map[string]string) *Resolver {
	return &Resolver{
		sites: sites,
	}
}

// Resolve takes a request url and finds the appropriate
// website id it is for
func (r *Resolver) Resolve(requestURL string) (string, error) {
	u, err := url.Parse(requestURL)
	if err != nil {
		return "", err
	}
	siteID, ok := r.sites[u.Hostname()]
	if !ok {
		return "", fmt.Errorf("siteID not found for request %s", requestURL)
	}
	return siteID, nil
}
