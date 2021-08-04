package bluemap

import (
	"github.com/unionofblackbean/api/common"
	"github.com/unionofblackbean/api/common/linalg"
)

type Marker struct {
	// all
	ID          *string         `json:"id"`
	Type        *string         `json:"type"`
	Map         *string         `json:"map"`
	Position    *linalg.Vec3f32 `json:"position"`
	Label       *string         `json:"label"`
	MinDistance *float32        `json:"minDistance"`
	MaxDistance *float32        `json:"maxDistance"`

	// poi
	Icon *string `json:"icon"`
	// html
	HTML *string `json:"html"`
	// poi, html
	Anchor *linalg.Vec2f32 `json:"anchor"`

	// extrude
	ShapeMinY *float32 `json:"shapeMinY"`
	ShapeMaxY *float32 `json:"shapeMaxY"`
	// shape
	ShapeY *float32 `json:"shapeY"`
	// extrude, shape
	Shape     []ShapeVertex `json:"shape"`
	FillColor *linalg.RGBA  `json:"fillColor"`
	// line
	Lines []linalg.Vec3i `json:"line"`
	// extrude, line, shape
	Detail    *string      `json:"detail"`
	DepthTest *bool        `json:"depthTest"`
	LineWidth *int         `json:"lineWidth"`
	LineColor *linalg.RGBA `json:"lineColor"`
}

func (m *Marker) IsValid() bool {
	if *m.ID == "" || *m.Map == "" || *m.Label == "" {
		return false
	}

	if common.Nil(m.Type, m.Position, m.MinDistance, m.MaxDistance) {
		return false
	}

	switch *m.Type {
	case MarkerTypePOI:
		if common.Nil(m.Icon, m.Anchor) {
			return false
		}
	case MarkerTypeHTML:
		if common.Nil(m.HTML, m.Anchor) {
			return false
		}
	case MarkerTypeExtrude:
		if common.Nil(
			m.ShapeMinY, m.ShapeMaxY,
			m.Shape,
			m.FillColor,
			m.Detail, m.DepthTest,
			m.LineWidth, m.LineColor) {
			return false
		}
	case MarkerTypeLine:
		if common.Nil(
			m.Lines,
			m.Detail, m.DepthTest,
			m.LineWidth, m.LineColor) {
			return false
		}
	case MarkerTypeShape:
		if common.Nil(
			m.ShapeY, m.Shape,
			m.FillColor,
			m.Detail, m.DepthTest,
			m.LineWidth, m.LineColor) {
			return false
		}
	default:
		return false
	}

	return true
}
