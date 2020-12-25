package esri

import (
	"fmt"
	"strings"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Form exposes some of the fields you need to query a feature layer
// using the esri.Query function
type Form struct {
	Token string
	Format         string
	Where          string
	OutFields      string
	BufferDistance float64
	Geometry       interface{}
}

type arcGISError struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func (e *arcGISError) Error() string {
	return e.Description
}

// Query is a generic query against an ArcGIS layer. Authentication is
// automatically handled using a global cached token.
//
// Defaults if you leave these form fields blank:
// form.Geometry -> No geometry in the query
// form.Format -> GeoJSON
// form.Where -> 1=1
// form.BufferDistance -> Use no buffer
func Query(endpoint string, form *Form) (*geojson.FeatureCollection, error) {
	if form.Where == "" {
		form.Where = "1=1"
	}

	if form.Format == "" {
		form.Format = "geojson"
	}

	req := url.Values{
		"token":     {form.Token},
		"outFields": {form.OutFields},
		"f":         {form.Format},
		"where":     {form.Where},
	}

	if form.Geometry != nil {
		geometryType, buf, err := geojsonToEsri(form.Geometry)
		if err != nil {
			return nil, err
		}

		req.Add("geometry", string(buf))
		req.Add("geometryType", geometryType)
	}

	if form.BufferDistance != 0 {
		req.Add("distance", fmt.Sprint(form.BufferDistance))
	}

	resp, err := request(endpoint, req)
	if err != nil {
		return nil, err
	}

	return geojson.UnmarshalFeatureCollection(resp)
}

func checkError(buf []byte) error {
	var data map[string]interface{}
	if err := json.Unmarshal(buf, &data); err != nil {
		return err
	}

	if data["error"] != nil {
		return fmt.Errorf("Error in AGOL request: %+v", string(buf))
	}

	return nil
}

func request(endpoint string, form url.Values) ([]byte, error) {
	httpResp, err := http.PostForm(endpoint, form)
	if err != nil {
		return nil, err
	}

	defer httpResp.Body.Close()
	resp, err := ioutil.ReadAll(httpResp.Body)

	if httpResp.StatusCode != http.StatusOK {
		var sb strings.Builder

		sb.WriteString(fmt.Sprintf("received HTTP %v from ArcGIS, something went wrong. ", httpResp.StatusCode))

		if err != nil {
			sb.WriteString("Tried extracting payload, but there was an error reading the response body")
		} else {
			sb.WriteString("Response body: " + string(resp))
		}

		return nil, fmt.Errorf(sb.String())
	}

	if err != nil {
		return nil, fmt.Errorf(
			"Failure making request to %v with HTTP form %+v: %v",
			endpoint,
			form,
			err,
		)
	} else if err = checkError(resp); err != nil {
		return nil, err
	}

	return resp, nil
}
