package esri

func geojsonToEsri(gj orb.Geometry) string{
	switch geometry := gj.(type) {
	case Polyline:
		return "esriGeometryPolyline"
	case Polygon:
		return "esriGeometryMultipoint"
	default:
		return ""
	}
}
