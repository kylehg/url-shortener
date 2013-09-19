package shawty

import (
	"fmt"
	"testing"
)

// Test that generating
func TestGenRandShortcode(t *testing.T) {
	for i := 0; i < 25; i++ {
		fmt.Println(genRandShortcode(5))
	}
}

// Test URL shortening
func _TestShortenUrl(t *testing.T) {
	test := func(url string, expectedCode string) string {
		code, err := ShortenUrl(url)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if expectedCode != "" && code != expectedCode  {
			t.Errorf("Expected %s -> %s, was %s", url, expectedCode, code)
		}
		return code
	}

	url1 := "http://google.com"
	url2 := "http://medium.com"
	code1 := test(url1, "")
	code2 := test(url2, "")
	test(url2, code2)
	test(url1, code1)
	if code1 == code2 {
		t.Error("Duplicate codes")
	}
}
