package content

// Content is the data for a single page
type Content struct {
	Template  string
	Data      map[string]interface{}
	Fragments map[string][]Fragment
}

// FragmentID is the id of a fragment
type FragmentID string

// Fragment is a the buiding block of a page
type Fragment struct {
	ID        FragmentID
	Template  string
	Data      map[string]interface{}
	Fragments map[string][]Fragment
}
