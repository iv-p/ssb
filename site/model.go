package site

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
	Meta map[string]string
}
