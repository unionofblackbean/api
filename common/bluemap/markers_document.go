package bluemap

type MarkersDocument struct {
	MarkerSets []MarkerSet `json:"markerSets"`
}

func (doc *MarkersDocument) IsValid() bool {
	for _, set := range doc.MarkerSets {
		if !set.IsValid() {
			return false
		}
	}

	return true
}
