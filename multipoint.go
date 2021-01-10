package esri

type MultiPoint struct {
	HasM       bool       `json:"hasM"`
	HasZ       bool       `json:"hasZ"`
	SpatialRef SpatialRef `json:"spatialRef"`
	Points [][]float64
}

func (m MultiPoint) Type() string {
	return "esriGeometryMultiPoint"
}
