package esri

// Polygon is an esri Polygon
type Polygon struct {
	HasM       bool       `json:"hasM"`
	HasZ       bool       `json:"hasZ"`
	SpatialRef SpatialRef `json:"spatialRef"`
	Rings [][][]float64 `json:"rings"`
}

func (p Polygon) Type() string {
	return "esriGeometryPolygon"
}
