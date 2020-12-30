package esri

// Polygon is an esri Polygon
type Polygon struct {
	composedType
	Rings [][][]float64 `json:"rings"`
}

func Type() string {
	return "esriGeometryPolygon"
}
