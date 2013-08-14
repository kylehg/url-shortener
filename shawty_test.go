package shawty

import (
	"fmt"
	"math"
	"testing"
)

// Test that generating shortcodes from integers works.
// func TestGenShortcode(t *testing.T) {
// 
// 	testShortcode := func(n int, str string) {
// 		code := genShortcode(n)
// 		if code != str {
// 			t.Errorf(fmt.Sprintf("Expected %d -> %s, was %s", n, str, code))
// 		}
// 	}
// 
// 	p := func(n float64) int { return int(math.Pow(63, n)) }
// 
// 	tests := []struct {
// 		n   int
// 		str string
// 	}{
// 		{0, "a"},           // 0, 0
// 		{1, "b"},           // 0, 1
// 		{26, "A"},          // 0, 26
// 		{26 + 26, "0"},     // 0, 52
// 		{26 + 26 + 9, "9"}, // 0, 61
// 		{62, "aa"},         // 0, 62
// 		{63, "ba"},         // 1, 0
// 		{p(1) + 1, "bb"},   // 1, 1
// 		{p(2) - 1, "aaa"},  // 0, 62, 62
// 		{p(2), "baa"},      // 1, 0, 0
// 		{p(3), "baaa"},     // 1, 0, 0, 0
// 	}
// 
// 	for _, test := range tests {
// 		testShortcode(test.n, test.str)
// 	}
// }

// Test URL shortening
// TODO
// func TestGetNextShortcode(t *testing.T) {
// 	url1 := "http://google.com"
// 	url2 := "http://medium.com"
// 
// 	code, err := s.ShortenUrl(url)
// 	tests := []struct {
// 		
// 	}{
// 		{url, code, shouldErr, msg},
// 
// 
// 		{url1, code1, false, msg},
// 		{url1, code1, false, msg},
// 		{url2, code2, false, msg},
// 		{url3, code3, false, msg},
// 		{url2, code2, false, msg},
// 		{url1, code1, false msg}
// 	}
// 
// 	wasSaved, err := s.ShortenUrlToCode(url, code)
// 	tests := []struct {
// 
// 	}{
// 		{url, code, wasSaved, shouldErr, msg},
// 
// 
// 		{url1, code1, true, false, msg},
// 		{url1, code1, true, false, msg},
// 
// 		{url2, code1, false, true, msg},
// 		{url2, code2, true, false, msg},
// 
// 		{url1, code1, true, false, msg},
// 		{url3, code1, true, false, msg},
// 		{url, code, wasSaved, shouldErr, msg},
// 
// 		{url1, code, wasSaved, shouldErr, msg},
// 		{url, code, wasSaved, shouldErr, msg},
// 		{url, code, wasSaved, shouldErr, msg},
// 		{url, code, wasSaved, shouldErr, msg},
// 		{url, code, wasSaved, shouldErr, msg},
// 		{url, code, wasSaved, shouldErr, msg},
// 		{url, code, wasSaved, shouldErr, msg},
// 	}
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
