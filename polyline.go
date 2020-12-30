package esri

// Polyline is an esri Polyline
//
// https://developers.arcgis.com/documentation/common-data-types/geometry-objects.htm
type Polyline struct {
	Polygon
}

func (p Polyline) Type() string {
	return "esriGeometryPolyline"
}
