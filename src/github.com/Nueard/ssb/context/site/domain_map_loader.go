package site

// DomainMap is a map of domains to site ids
type DomainMap map[string]string

// DummyData holds a dummy map for local development
var DummyData = DomainMap{
	"localhost":                 "1",
	"dummydomain.com":           "2",
	"subdomain.dummydomain.com": "3",
}

// IDomainMapLoader is the interface of a domain map loader
type IDomainMapLoader interface {
	Load() DomainMap
}

// DomainMapLoader resolves site url data to a dummy object
type DomainMapLoader struct {
}

// NewDomainMapLoader creates and returns a new dummy loader object
func NewDomainMapLoader() *DomainMapLoader {
	return &DomainMapLoader{}
}

// Load fetches the DomainMap from an appropriate data source and returns it
func (dl *DomainMapLoader) Load() DomainMap {
	return DummyData
}
