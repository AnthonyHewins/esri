package esri

// Polygon is an esri Polygon
type Polygon struct {
	Rings      [][]Point `json:"rings"`
	HasM       bool       `json:"hasM"`
	HasZ       bool       `json:"hasZ"`
	SpatialRef SpatialRef `json:"spatialRef"`
}
