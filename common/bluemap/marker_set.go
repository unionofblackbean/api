package bluemap

type MarketSet struct {
	ID          string   `json:"id"`
	Label       string   `json:"label"`
	Toggleable  bool     `json:"toggleable"`
	DefaultHide bool     `json:"defaultHide"`
	Markers     []Marker `json:"markers"`
}
