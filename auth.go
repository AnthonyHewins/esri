package esri

import (
	"fmt"

	"encoding/json"
	"net/url"
)

const (
	referer = "https://www.arcgis.com"
)

var (
	generateTokenURL = fmt.Sprintf("%v/sharing/rest/generateToken?f=json", referer)
)

type TokenResp struct {
	Token   string `json:"token"`
	Expires int64  `json:"expires"`
}

func GetToken(username, password string) (*TokenResp, error) {
	resp, err := request(generateTokenURL, url.Values{
		"username": {username},
		"password": {password},
		"referer":  {referer},
	})

	if err != nil {
		return nil, err
	}

	var tokenStruct TokenResp
	return &tokenStruct, json.Unmarshal(resp, &tokenStruct)
}
