package bluemap

type MarkerSet struct {
	ID          string    `json:"id"`
	Label       string    `json:"label"`
	Toggleable  bool      `json:"toggleable"`
	DefaultHide bool      `json:"defaultHide"`
	Markers     []*Marker `json:"markers"`
}

func (set *MarkerSet) IsValid() bool {
	for _, marker := range set.Markers {
		if !marker.IsValid() {
			return false
		}
	}

	return set.ID != "" &&
		set.Label != ""
}
