package bluemap

import "github.com/unionofblackbean/api/common/linalg"

type Marker struct {
	// all
	ID          string         `json:"id"`
	Type        string         `json:"type"`
	Map         string         `json:"map"`
	Position    linalg.Vec3f32 `json:"position"`
	Label       string         `json:"label"`
	MinDistance float32        `json:"minDistance"`
	MaxDistance float32        `json:"maxDistance"`

	// poi
	Icon string `json:"icon"`

	// html
	HTML string `json:"html"`

	// poi, html
	Anchor linalg.Vec2f32 `json:"anchor"`

	// extrude, shape
	Shape []ShapeVertex `json:"shape"`

	// extrude
	ShapeMinY float32 `json:"shapeMinY"`
	ShapeMaxY float32 `json:"shapeMaxY"`

	// shape
	ShapeY    float32     `json:"shapeY"`
	FillColor linalg.RGBA `json:"fillColor"`

	// line
	Lines []linalg.Vec3i `json:"line"`

	// extrude, line, shape
	Detail    string      `json:"detail"`
	DepthTest bool        `json:"depthTest"`
	LineWidth int         `json:"lineWidth"`
	LineColor linalg.RGBA `json:"lineColor"`
}
