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

//	wasSaved, err := s.ShortenUrlToCode(url, code)
//	tests := []struct {
//
//	}{
//		{url, code, wasSaved, shouldErr, msg},
//
//
//		{url1, code1, true, false, msg},
//		{url1, code1, true, false, msg},
//
//		{url2, code1, false, true, msg},
//		{url2, code2, true, false, msg},
//
//		{url1, code1, true, false, msg},
//		{url3, code1, true, false, msg},
//		{url, code, wasSaved, shouldErr, msg},
//
//		{url1, code, wasSaved, shouldErr, msg},
//		{url, code, wasSaved, shouldErr, msg},
//		{url, code, wasSaved, shouldErr, msg},
//		{url, code, wasSaved, shouldErr, msg},
//		{url, code, wasSaved, shouldErr, msg},
//		{url, code, wasSaved, shouldErr, msg},
//		{url, code, wasSaved, shouldErr, msg},
//	}
//
//
//
//
//
//
//
//
//
//
//
// 	testSetUrl := func(url string, inputCode string, expectedCode string, expectedErr bool, errStr string) {
// 		if inputCode == "" {
// 			code, err := s.ShortenUrl(url)
//
// 			// Getting a new random code should never throw an error
// 			if err != nil && !expectedErr {
// 				// TODO: Toss up error
// 				t.Errorf("TODO")
// 			}
//
// 			if code != expectedCode && !expectedErr {
// 				t.Errorf(errStr)
// 			}
// 		} else {
// 			saved, extantCode, err := s.ShortenUrlToCode(url, code)
//
// 			if code != nil && extantCode != nil && !expectedErr {
// 				t.Errorf("The code and extant code shouldn't both return")
// 			}
//
// 			if code != extantCode {
// 				t.Errorf(errStr)
// 			}
//
// 			}
// 		}
// 	}
// 	nextCode, err := s.genNextShortcode()
// }
