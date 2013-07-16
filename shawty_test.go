package shawty_test

import (
	"testing"
	. "shawty"
)

// Test that generating shortcodes from integers works.
func TestGenShortcode(t *testing.T) {

	testShortcode := func(n int, str string, errStr string) {
		code, err := genShortcode(n)
		if code != str || err != nil {
			t.Errorf(errStr)
		}
	}

	tests := []struct {
		n int
		str string
		errStr string
	}{
		{0, "aaaa", "TODO"},
		{1, "aaab", "TODO"},
		{26, "aaaA", "TODO"},
		{52, "aaa1", "TODO"},
		{63, "aaa_", "TODO"},
		{64, "aaaaa", "TODO"},
	}

	// TODO iter tests
}

// Test URL shortening
// TODO
func TestGetNextShortcode(t *testing.T) {
	url1 := "http://google.com"
	url2 := "http://medium.com"

	code, err := s.ShortenUrl(url)
	tests := []struct {

	}{
		{url, code, shouldErr, msg},


		{url1, code1, false, msg},
		{url1, code1, false, msg},
		{url2, code2, false, msg},
		{url3, code3, false, msg},
		{url2, code2, false, msg},
		{url1, code1, false msg},
	}

	wasSaved, err := s.ShortenUrlToCode(url, code)
	tests := []struct {

	}{
		{url, code, wasSaved, shouldErr, msg},


		{url1, code1, true, false, msg},
		{url1, code1, true, false, msg},

		{url2, code1, false, true, msg},
		{url2, code2, true, false, msg},

		{url1, code1, true, false, msg},
		{url3, code1, true, false, msg},
		{url, code, wasSaved, shouldErr, msg},

		{url1, code, wasSaved, shouldErr, msg},
		{url, code, wasSaved, shouldErr, msg},
		{url, code, wasSaved, shouldErr, msg},
		{url, code, wasSaved, shouldErr, msg},
		{url, code, wasSaved, shouldErr, msg},
		{url, code, wasSaved, shouldErr, msg},
		{url, code, wasSaved, shouldErr, msg},
	}










	testSetUrl := func(url string, inputCode string, expectedCode string, expectedErr bool, errStr string) {
		if inputCode == "" {
			code, err := s.ShortenUrl(url)

			// Getting a new random code should never throw an error
			if err != nil && !expectedErr {
				// TODO: Toss up error
				t.Errorf("TODO")
			}

			if code != expectedCode && !expectedErr {
				t.Errorf(errStr)
			}
		} else {
			saved, extantCode, err := s.ShortenUrlToCode(url, code)

			if code != nil && extantCode != nil && !expectedErr {
				t.Errorf("The code and extant code shouldn't both return")
			}

			if code != extantCode {
				t.Errorf(errStr)
			}

			}
		}
	}
	nextCode, err := s.genNextShortcode()
}
