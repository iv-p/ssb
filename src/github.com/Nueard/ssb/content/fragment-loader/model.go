package fragmentloader

// ID is the id of a fragment
type ID string

// Fragment is a the buiding block of a page
type Fragment struct {
	ID        ID
	Template  string
	Data      map[string]interface{}
	Fragments map[string][]Fragment
}

// Raw is the database representation of a fragment
type Raw struct {
	ID        ID
	Template  string
	Data      map[string]interface{}
	Fragments map[string][]ID
}
