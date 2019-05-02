package context

// Site holds the root information about a site
type Site struct {
	ID string

	Head Head

	PagesIDs []string
}

// Head holds information about what's in the html head of a website.
// These are the base values and will be overwritten by the head of
// specific pages.
type Head struct {
	// Meta is a map of meta fields in the head. Every key pair will be
	// added as an separate meta tag with key being the name and the
	// value being the content
	Title string

	Meta   map[string]Meta
	Script map[string]Script
	Style  map[string]Style
}

// Meta holds information for a single meta tag in the html head
type Meta struct {
	Name    string
	Content string
}

// Script holds information for a single script tag in the html head
type Script struct {
	Async   bool
	Charset string
	Defer   bool
	Src     string
	Type    string
}

// Style holds information for a single link (css) tag in the html head
type Style struct {
	CrossOrigin string
	Type        string
	Href        string
	Rel         string
	Media       string
}
