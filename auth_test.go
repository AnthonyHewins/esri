package esri

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetToken(t *testing.T) {
	testCases := []struct {
		name        string
		before      func()
		expected    string
		expectedErr error
	}{
		{
			"Returns tokenHolder when not the empty string, the expiration is nonzero and before expiration",
			func() {
				tokenHolder = "non empty"
				tokenExpiration = time.Now().Add(time.Hour)
			},
			"non empty",
			nil,
		},
	}
	for _, tc := range testCases {
		tc.before()
		actual, actualErr := getToken()
		assert.Equal(t, tc.expected, actual)
		assert.Equal(t, tc.expectedErr, actualErr)
	}
}

func TestParseMilliseconds(t *testing.T) {
	testCases := []struct {
		expected time.Time
		arg      int64
	}{
		{
			time.Unix(0, 1592498776848000000),
			1592498776848,
		},
	}
	for _, tc := range testCases {
		assert.InDelta(t, tc.expected.UnixNano(), parseMilliseconds(tc.arg).UnixNano(), 1000000)
	}
}
