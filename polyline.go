package esri

// Polyline is an esri Polyline
//
// https://developers.arcgis.com/documentation/common-data-types/geometry-objects.htm
type Polyline struct {
	HasM       bool       `json:"hasM"`
	HasZ       bool       `json:"hasZ"`
	SpatialRef SpatialRef `json:"spatialRef"`
	Rings [][][]float64 `json:"rings"`
}

func (p Polyline) Type() string {
	return "esriGeometryPolyline"
}
