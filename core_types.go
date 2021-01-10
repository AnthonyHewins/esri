package esri

import "time"

type EsriGeom interface {
	Type() string
}

// Form exposes some of the fields you need to query a feature layer
// using the esri.Query function
type Form struct {
	Where          string
	ObjectIDs []int
	Geometry       EsriGeom
	InSR interface{}

	// Use WKID and WKD helper structs for this
	SpatialRel interface{}
	RelationParam string
	Time time.Time
	BufferDistance float64
	OutFields      []string
	NoGeometry bool
	MaxAllowableOffset int
	geometryPrecision int
	OutSR string
	GDB string
	Distinct bool
	OnlyIDs bool
	OnlyCount bool
	OnlyExtent bool
	OrderBy []string
	GroupByForStats []string
	ReturnZ bool
	ReturnM bool
	Offset int
	Limit int
	ReturnCentroid bool
	ResultType string
	HistoricMoment time.Time
	ReturnTrueCurves bool
	Token          string
	Format         string
}

type SpatialRef struct {
 WKID int`json:"wkid"`
 LatestWKID int `json:"latestWkid"`
 VCSWKID int `json:"vcsWkid"`
 LatestVCSWKID int `json:"latestVcsWkid"`
}

//type WKID struct {
//	WKID int `json:"wkid"`
//	LatestWKID int `json:"latestWkid"`
//	VCSWKID int `json:"vcsWkid"`
//	LatestVCSWKID int `json:"latestVcsWkid"`
//}
//
//type WKT struct {
//	WKT string `json:"wkt"`
//}
