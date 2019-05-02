package content

import (
	"fmt"

	fragmentloader "github.com/Nueard/ssb/content/fragment-loader"
	pageloader "github.com/Nueard/ssb/content/page-loader"
	"github.com/Nueard/ssb/context/page"
	"github.com/Nueard/ssb/context/site"
)

// Loader loads all the information about a page by it's id
type Loader struct {
	fragmentLoader fragmentloader.IDeepLoader
	pageLoader     pageloader.ILoader
}

// NewLoader creates a new loader object
func NewLoader(pageLoader pageloader.ILoader, fragmentLoader fragmentloader.IDeepLoader) *Loader {
	return &Loader{
		fragmentLoader: fragmentLoader,
		pageLoader:     pageLoader,
	}
}

// Load returns the page contents by id
func (l *Loader) Load(siteContext site.Context, pageContext page.Context) Content {
	page := l.pageLoader.Load(pageloader.IDFromPageContext(pageContext))

	fragment := l.fragmentLoader.Load(page.FragmentID)

	fmt.Println(fragment)
	return Content{
		Fragment: fragment,
		// Fragment: fragment.Fragment{
		// 	Template: "layouts/simple-layout-one.html",
		// 	Fragments: map[string][]fragment.Fragment{
		// 		"main": []fragment.Fragment{
		// 			fragment.Fragment{
		// 				Template: "partials/1.html",
		// 				Fragments: map[string][]fragment.Fragment{
		// 					"what": []fragment.Fragment{
		// 						fragment.Fragment{
		// 							Template:  "partials/2.html",
		// 							Fragments: map[string][]fragment.Fragment{},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// },
	}
}

func (l *Loader) loadFragment(id fragmentloader.ID) fragmentloader.Fragment {
	return l.loadFragment(id)
}
