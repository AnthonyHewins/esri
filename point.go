package esri

type Point struct {
	X          float64    `json:"x"`
	Y          float64    `json:"y"`
	Z          float64    `json:"z"`
	M          float64    `json:"m"`
	SpatialRef SpatialRef `json:"spatialRef"'`
}

func (p Point) Type() string {
	return "esriGeometryPoint"
}
