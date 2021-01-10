package esri

import (
	"fmt"
	"time"

	"encoding/json"
	"net/url"
)

const (
	referer = "https://www.arcgis.com"
)

var (
	generateTokenURL = fmt.Sprintf("%v/sharing/rest/generateToken", referer)
)

type tokenResp struct {
	Token   string `json:"token"`
	Expires int64  `json:"expires"`
}

// An AGOL token response
type Token struct {
	Token string
	Expires time.Time
}

// Expired is a shorthand for telling whether or not a token is expired
func (t Token) Expired() bool {
	return !time.Now().Before(t.Expires)
}

func GetToken(username, password string) (*Token, error) {
	resp, err := request(generateTokenURL, url.Values{
		"username": {username},
		"password": {password},
		"referer":  {referer},
		"f": {"json"},
	})

	if err != nil {
		return nil, err
	}

	var t tokenResp
	err = json.Unmarshal(resp, &t)
	if err != nil {
		return nil, err
	}

	return &Token{Token: t.Token, Expires: parseMilliseconds(t.Expires)}, nil
}

// parseMilliseconds parses a time in milliseconds and spits out a time.Time
func parseMilliseconds(expiration int64) time.Time {
	// Milliseconds elapsed since the epoch
	//        1,000 micro sec     1,000 ns
	// X ms * --------------- * -------------- = 1,000,000 * X ns
	//           1 ms             1 micro sec

	return time.Unix(
		0,
		expiration*1000000,
	)
}
