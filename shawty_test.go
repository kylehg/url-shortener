package shawty

import "testing"

// Test that generating shortcodes from integers works.
func TestGenShortcode(t *testing.T) {

	testShortcode := func(n int, str string, errStr string) {
		code, err := s.genShortcode(n)
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

	code, err := genShortcode(-1)
	if err == nil {
		t.Errorf("TODO")
	}
}

// Test that getting a next shortcode works
func TestGetNextShortcode(t *testing.T) {
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
