package fragmentloader

import "fmt"

// IDeepLoader is Loader's interface
type IDeepLoader interface {
	Load(ID) Fragment
}

// DeepLoader loads fragments from a datastore
type DeepLoader struct {
	rawLoader IRawLoader
}

// NewDeepLoader creates and returns a new fragment loader
func NewDeepLoader(rawLoader IRawLoader) *DeepLoader {
	return &DeepLoader{
		rawLoader: rawLoader,
	}
}

// Load retrieves one fragmet from the datastore by its id
func (dl *DeepLoader) Load(id ID) Fragment {
	return dl.load(id)
}

func (dl *DeepLoader) load(id ID) Fragment {
	fmt.Println(id)
	raw := dl.rawLoader.Load(id)
	fragment := fragmentFromRaw(raw)
	for segmentName, segmentFragmentIDs := range raw.Fragments {
		var segmentFragments []Fragment
		for _, segmentFragmentID := range segmentFragmentIDs {
			segmentFragment := dl.load(segmentFragmentID)
			segmentFragments = append(segmentFragments, segmentFragment)
		}
		fragment.Fragments[segmentName] = segmentFragments
	}
	return fragment
}

func fragmentFromRaw(raw Raw) Fragment {
	return Fragment{
		ID:        raw.ID,
		Template:  raw.Template,
		Data:      raw.Data,
		Fragments: make(map[string][]Fragment),
	}
}
