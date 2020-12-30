package esri

type EsriGeom interface {
	Type() string
}

// Form exposes some of the fields you need to query a feature layer
// using the esri.Query function
type Form struct {
	Token          string
	Format         string
	Where          string
	OutFields      string
	BufferDistance float64
	Geometry       EsriGeom
}

type arcGISError struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func (e *arcGISError) Error() string {
	return e.Description
}

type composedType struct {
	HasM       bool       `json:"hasM"`
	HasZ       bool       `json:"hasZ"`
	SpatialRef SpatialRef `json:"spatialRef"`
}
