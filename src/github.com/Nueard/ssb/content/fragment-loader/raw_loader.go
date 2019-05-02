package fragmentloader

var dummyData = map[ID]Raw{
	"layout": Raw{
		Template: "layouts/simple-layout-one.html",
		Fragments: map[string][]ID{
			"main": []ID{"1"},
		},
	},
	"1": Raw{
		Template: "partials/1.html",
		Fragments: map[string][]ID{
			"what": []ID{"2"},
		},
	},
	"2": Raw{
		Template:  "partials/2.html",
		Fragments: map[string][]ID{},
	},
}

// IRawLoader is Loader's interface
type IRawLoader interface {
	Load(ID) Raw
}

// RawLoader loads fragments from a datastore
type RawLoader struct {
}

// NewRawLoader creates and returns a new fragment loader
func NewRawLoader() *RawLoader {
	return &RawLoader{}
}

// Load retrieves one fragmet from the datastore by its id
func (l *RawLoader) Load(id ID) Raw {
	return dummyData[id]
}
