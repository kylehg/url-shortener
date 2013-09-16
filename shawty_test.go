package shawty

import (
	"fmt"
	"math"
	"testing"
)

// Test URL shortening
func TestGetNextShortcode(t *testing.T) {
	testUrlCode := func(url string, expectedCode string) string {
		code, err := ShortenUrl(url)
		if err != nil {
			t.Error("Error: " + err.Error())
		}
		if code != expectedCode {
			t.Errorf("Expected %s -> %s, was %s")
		}
	}

	url1 := "http://google.com"
	url2 := "http://medium.com"
	tests := []struct {
		url  string
		code string
	}{
		{url1, code1},
		{url2, code2},
		{url1, code1},
		{url2, code2},
		{url2, code2},
		{url1, code1},
	}
	for _, test := range tests {
		testUrlCode(test.url, test.code)
	}
}
