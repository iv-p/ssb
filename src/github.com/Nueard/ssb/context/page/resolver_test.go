package page

import (
	"net/url"
	"testing"

	"github.com/Nueard/ssb/context/site"
)

var pages = URLMap{
	"/":              "1",
	"/something/one": "2",
	"/another":       "3",
}

type DummyURLMapLoader struct {
	IURLMapLoader
}

func (ddml *DummyURLMapLoader) Load(context site.Context) URLMap {
	return pages
}

var tables = []struct {
	url     string
	context Context
	err     bool
}{
	{"/", Context{PageID: "1"}, false},
	{"/something/one", Context{PageID: "2"}, false},
	{"/something/one?query=params", Context{PageID: "2"}, false},
	{"/another", Context{PageID: "3"}, false},

	{"another", Context{}, true},
	{"/missing/url", Context{}, true},
	{"", Context{}, true},
}

func TestResolver(t *testing.T) {
	resolver := NewResolver(&DummyURLMapLoader{})
	for _, table := range tables {
		u, _ := url.Parse(table.url)
		context, err := resolver.Resolve(site.Context{}, u)
		if context.PageID != table.context.PageID {
			t.Errorf("Expected PageID %s for url %s, got PageID %s", table.context.PageID, table.url, context.PageID)
		}
		if table.err && err == nil {
			t.Errorf("Expected error for url '%s', but got nil", table.url)
		}
		if !table.err && err != nil {
			t.Errorf("Did not expect error for url '%s', but got '%s'", table.url, err)
		}
	}
}
