package esri

// Polyline is an esri Polyline
//
// https://developers.arcgis.com/documentation/common-data-types/geometry-objects.htm
type Polyline struct {
	HasM             bool             `json:"hasM"`
	HasZ             bool             `json:"hasZ"`
	Paths            []Point `json:"paths"`
	SpatialReference SpatialRef       `json:"spatialReference"`
}
