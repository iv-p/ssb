package content

// Content is the data for a single page
type Content struct {
	Layout        Layout
	RootFragments []FragmentID
	Fragments     map[FragmentID]Fragment
}

// LayoutID is the id of a fragment
type LayoutID string

// Layout is the root layout of a page
type Layout struct {
	ID       LayoutID
	Template string
	Data     map[string]interface{}
}

// FragmentID is the id of a fragment
type FragmentID string

// Fragment is a the buiding block of a page
type Fragment struct {
	ID        FragmentID
	Template  string
	Data      map[string]interface{}
	Fragments []FragmentID
}
