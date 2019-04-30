package site

import (
	"testing"
)

var sites = map[string]string{
	"dummydomain.com":           "site1",
	"subdomain.dummydomain.com": "site2",
}

var tables = []struct {
	url    string
	siteID string
	err    bool
}{
	{"https://dummydomain.com/some/path", "site1", false},
	{"https://subdomain.dummydomain.com/", "site2", false},
	{"https://subdomain.dummydomain.com/#asdf", "site2", false},

	{"https://nonexistingdomain.com/query123", "", true},
	{"notadomain", "", true},
	{"", "", true},
}

func TestResolver(t *testing.T) {
	resolver := NewResolver(sites)
	for _, table := range tables {
		siteID, err := resolver.Resolve(table.url)
		if siteID != table.siteID {
			t.Errorf("Expected siteID %s for url %s, got siteID %s", table.siteID, table.url, siteID)
		}
		if table.err && err == nil {
			t.Errorf("Expected error for url '%s', but got nil", table.url)
		}
		if !table.err && err != nil {
			t.Errorf("Did not expect error for url '%s', but got '%s'", table.url, err)
		}
	}
}
