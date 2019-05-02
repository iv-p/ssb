package pageloader

import "github.com/Nueard/ssb/context/page"

// ILoader is page loaders interface
type ILoader interface {
	Load(ID) Raw
}

// Loader loads the data for a single page
type Loader struct {
}

// NewLoader creates and returns a new page loader
func NewLoader() *Loader {
	return &Loader{}
}

// Load fetches and returns a page by its id
func (l *Loader) Load(id ID) Raw {
	return Raw{
		ID:         id,
		FragmentID: "layout",
	}
}

// IDFromPageContext gets the page id from the page context
func IDFromPageContext(context page.Context) ID {
	return ID(context.PageID)
}
