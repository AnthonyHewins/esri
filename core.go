package esri

import (
	"fmt"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ToValues interface{
	ToValues() (url.Values, error)
}

// Query is a generic query against an ArcGIS layer. It will hit `endpoint` using a POST body
// supplied by `form`. Implement the `ToValues` interface to create your own form, or use `esri.Form`
// supplied by the package to make a request. Using `esri.Form` has the following defaults:
//
// form.Format -> GeoJSON
// form.Where -> 1=1
func Query(endpoint string, form ToValues) ([]byte, error) {
	if form == nil {
		return request(endpoint, nil)
	}

	req, err := form.ToValues()
	if err != nil {
		return nil, err
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
