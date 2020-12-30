package esri

type MultiPoint struct {
	composedType

	Points [][]float64
}

func (m MultiPoint) Type() string {
	return "esriGeometryMultiPoint"
}
