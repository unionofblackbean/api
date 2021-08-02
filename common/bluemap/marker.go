package bluemap

import "github.com/unionofblackbean/api/common/linalg"

type Marker struct {
	// all
	ID          string         `json:"id,omitempty"`
	Type        string         `json:"type,omitempty"`
	Map         string         `json:"map,omitempty"`
	Position    linalg.Vec3f32 `json:"position,omitempty"`
	Label       string         `json:"label,omitempty"`
	MinDistance float32        `json:"minDistance,omitempty"`
	MaxDistance float32        `json:"maxDistance,omitempty"`

	// poi
	Icon string `json:"icon,omitempty"`
	// html
	HTML string `json:"html,omitempty"`
	// poi, html
	Anchor linalg.Vec2f32 `json:"anchor,omitempty"`

	// extrude
	ShapeMinY float32 `json:"shapeMinY,omitempty"`
	ShapeMaxY float32 `json:"shapeMaxY,omitempty"`
	// shape
	ShapeY float32 `json:"shapeY,omitempty"`
	// extrude, shape
	Shape     []ShapeVertex `json:"shape,omitempty"`
	FillColor linalg.RGBA   `json:"fillColor,omitempty"`
	// line
	Lines []linalg.Vec3i `json:"line,omitempty"`
	// extrude, line, shape
	Detail    string      `json:"detail,omitempty"`
	DepthTest bool        `json:"depthTest,omitempty"`
	LineWidth int         `json:"lineWidth,omitempty"`
	LineColor linalg.RGBA `json:"lineColor,omitempty"`
}

func (m *Marker) IsValid() bool {
	if m.ID == "" || m.Map == "" || m.Label == "" {
		return false
	}

	switch m.Type {
	case MarkerTypePOI:
	case MarkerTypeHTML:
	case MarkerTypeExtrude:
	case MarkerTypeLine:
	case MarkerTypeShape:
	default:
		return false
	}

	return true
}
