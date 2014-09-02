package shawty

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	ALPHABET         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	ALPHABET_LEN     = len(ALPHABET)
	DEFAULT_CODE_LEN = 5
)

// Generate a string of random characters from ALPHABET
func GetRandomShortcode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]string, DEFAULT_CODE_LEN)
	for i := 0; i < DEFAULT_CODE_LEN; i++ {
		code[i] = string(ALPHABET[r.Intn(ALPHABET_LEN)])
	}
	return strings.Join(code, "")
}

// Check if a URL is valid
func isValidUrl(url string) bool {
	return true // TODO
}

// Shorten a URL to a default code
func ShortenDefault(url string) (string, error) {
	if !isValidUrl(url) {
		return "", fmt.Errorf("%s is not a valid URL", url)
	}

	code := GetRandomShortcode()
	if err := SetDefaultCode(url, code); err != nil {
		return ShortenDefault(url)
	}
	return code, nil
}

// Shorten a URL to a custom code
func ShortenCustom(url string, code string) error {
	if !isValidUrl(url) {
		return fmt.Errorf("%s is not a valid URL", url)
	}

	return SetCustomCode(url, code)
}
