package esri

import (
	"fmt"
	"strings"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Query is a generic query against an ArcGIS layer
//
// Defaults if you leave these form fields blank:
// form.Format -> GeoJSON
// form.Where -> 1=1
func Query(endpoint string, form *Form) ([]byte, error) {
	if form.Where == "" {
		form.Where = "1=1"
	}

	if form.Format == "" {
		form.Format = "geojson"
	}

	req := url.Values{
		"token":     {form.Token},
		"outFields": {strings.Join(form.OutFields, ",")},
		"f":         {form.Format},
		"where":     {form.Where},
	}

	if form.Geometry != nil {
		buf, err := json.Marshal(form.Geometry)
		if err != nil {
			return nil, err
		}

		req.Add("geometry", string(buf))
		req.Add("geometryType", form.Geometry.Type())
	}

	if form.BufferDistance != 0 {
		req.Add("distance", fmt.Sprint(form.BufferDistance))
	}

	return request(endpoint, req)
}

func checkError(buf []byte) error {
	var data map[string]interface{}
	if err := json.Unmarshal(buf, &data); err != nil {
		return fmt.Errorf("While checking if the ArcGIS response has any errors, there was an error encountered: %v.\nRaw response: %v", err, string(buf))
	}

	if data["error"] != nil {
		return fmt.Errorf("Error in ArcGIS response: %+v", string(buf))
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

	if err != nil {
		return nil, fmt.Errorf(
			"Failure reading response from %v with HTTP form %+v: %v.\nRaw response: %+v",
			endpoint,
			form,
			err,
			string(resp),
		)
	}

	if httpResp.StatusCode != http.StatusOK {
		var sb strings.Builder

		sb.WriteString(fmt.Sprintf("received HTTP %v from ArcGIS:", httpResp.StatusCode))

		if err != nil {
			sb.WriteString(" tried extracting payload, but there was an error reading the response body")
		} else {
			sb.WriteString(" got response body: " + string(resp))
		}

		return nil, fmt.Errorf(sb.String())
	}

	return resp, checkError(resp)
}
