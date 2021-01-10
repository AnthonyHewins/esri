package esri

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type EsriGeom interface {
	Type() string
}

type Format string

const (
	HTML    Format = "html"
	PJSON   Format = "pjson"
	JSON    Format = "kmz"
	GeoJSON Format = "geojson"
	PBF     Format = "pbf"
)

// Form exposes some of the fields you need to query a feature layer
// using the esri.Query function
type Form struct {
	Where     string
	ObjectIDs []int
	Geometry  EsriGeom
	InSR      interface{}

	// Use WKID and WKD helper structs for this
	SpatialRel         interface{}
	RelationParam      string
	Time               time.Time
	BufferDistance     float64
	OutFields          []string
	ReturnNoGeometry         bool
	MaxAllowableOffset int
	geometryPrecision  int
	OutSR              string
	GDB                string
	Distinct           bool
	OnlyIDs            bool
	OnlyCount          bool
	OnlyExtent         bool
	OrderBy            []string
	GroupByForStats    []string
	ReturnZ            bool
	ReturnM            bool
	Offset             int
	Limit              int
	ReturnCentroid     bool
	ResultType         string
	HistoricMoment     time.Time
	ReturnTrueCurves   bool
	Token              string
	Format             Format
}

func (f Form) ToValues() (url.Values, error) {
	if f.Format == "" {
		f.Format = GeoJSON
	}

	if f.Where == "" {
		f.Where = "1=1"
	}

	req := url.Values{
		"token":     {f.Where},
		"f": {string(f.Format)},
	}

	if f.Token != "" {
		req.Add("token", f.Token)
	}

	if len(f.OutFields) > 0 {
		req.Add("outFields", strings.Join(f.OutFields, ","))
	}

	if f.Geometry != nil {
		buf, err := json.Marshal(f.Geometry)
		if err != nil {
			return nil, err
		}

		req.Add("geometry", string(buf))
		req.Add("geometryType", f.Geometry.Type())
	}

	if f.BufferDistance != 0 {
		req.Add("distance", fmt.Sprint(f.BufferDistance))
	}

	return req, nil
}

type SpatialRef struct {
	WKID          int `json:"wkid"`
	LatestWKID    int `json:"latestWkid"`
	VCSWKID       int `json:"vcsWkid"`
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
