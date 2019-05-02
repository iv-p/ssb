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
		Template: "layouts/simple-layout-one.html",
		Fragments: map[string][]Fragment{
			"main": []Fragment{
				Fragment{
					Template: "partials/1.html",
					Fragments: map[string][]Fragment{
						"what": []Fragment{
							Fragment{
								Template:  "partials/2.html",
								Fragments: map[string][]Fragment{},
							},
						},
					},
				},
			},
		},
	}
}
