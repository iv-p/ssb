package site

import (
	"net/url"
	"testing"
)

var sites = map[string]string{
	"dummydomain.com":           "site1",
	"subdomain.dummydomain.com": "site2",
}

type DummyDomainMapLoader struct {
	IDomainMapLoader
}

func (ddml *DummyDomainMapLoader) Load() DomainMap {
	return sites
}

var tables = []struct {
	url     string
	context Context
	err     bool
}{
	{"https://dummydomain.com/some/path", Context{SiteID: "site1"}, false},
	{"https://subdomain.dummydomain.com/", Context{SiteID: "site2"}, false},
	{"https://subdomain.dummydomain.com/#asdf", Context{SiteID: "site2"}, false},
	{"https://dummydomain.com:8080/", Context{SiteID: "site1"}, false},

	{"https://nonexistingdomain.com/query123", Context{}, true},
	{"subdomain.dummydomain.com", Context{}, true},
	{"notadomain", Context{}, true},
	{"", Context{}, true},
}

func TestResolver(t *testing.T) {
	resolver := NewResolver(&DummyDomainMapLoader{})
	for _, table := range tables {
		u, _ := url.Parse(table.url)
		context, err := resolver.Resolve(u)
		if context.SiteID != table.context.SiteID {
			t.Errorf("Expected siteID %s for url %s, got siteID %s", table.context.SiteID, table.url, context.SiteID)
		}
		if table.err && err == nil {
			t.Errorf("Expected error for url '%s', but got nil", table.url)
		}
		if !table.err && err != nil {
			t.Errorf("Did not expect error for url '%s', but got '%s'", table.url, err)
		}
	}
}
