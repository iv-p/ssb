package content

import (
	"github.com/Nueard/ssb/context/page"
	"github.com/Nueard/ssb/context/site"
)

// Loader loads all the information about a page by it's id
type Loader struct {
}

// NewLoader creates a new loader object
func NewLoader() *Loader {
	return &Loader{}
}

// Load returns the page contents by id
func (l *Loader) Load(siteContext site.Context, pageContext page.Context) Content {
	return Content{
		Layout: Layout{
			Template: "layouts/simple-layout-one.html",
		},
		RootFragments: []FragmentID{
			"1",
		},
		Fragments: map[FragmentID]Fragment{
			"1": Fragment{
				Template: "partials/1.html",
				Fragments: []FragmentID{
					"2",
				},
			},
			"2": Fragment{
				Template: "partials/2.html",
			},
		},
	}
}
