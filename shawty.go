package shawty

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"strings"
	"time"
)

// TODO
const network = ""
const address = ""

var conn, err = redis.Dial(network, address)

const ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"0123456789"
const ALPHABET_LEN = len(ALPHABET)
const DEFAULT_LEN = 5

// Generate a string of n random chars from ALPHABET
func genRandShortcode(n int) string {
	code := make([]string, n)
	for i := 0; i < n; i++ {
		code[i] = rand.Intn(ALPHABET_LEN)
	}
	res := strings.Join(code, "")
	fmt.Printf("-> %s\n", res)
	return res
}

// Convert a shortcode to a shotcode key in Redis
func codeKey(code string) string {
	return "code:" + code
}

// Shorten the given URL and return the shortcode
func ShortenUrl(url string) (code string, err error) {
	var code string

	for exists := true; exists; {
		code = genRandShortcode(DEFAULT_LEN)
		key = codeKey(code)

		exists, err := conn.Do("SETNX", key, url)
		if err != nil {
			return
		}
	}

	return
}

// Shorten the given URL to the given code. Return true if the URL was
// successfully saved to the code, false if the code was taken.
func ShortenUrlToCode(url string, code string) (success bool, err error) {
	success := false

	// Attempt to set the shortcode
	wasSet, err := redis.Bool(conn.Do("SETNX", codeKey(code), url))
	if err != nil {
		return
	}

	if !wasSet {
		setUrl, err := conn.Do("GET", codeKey(code))
		success = setUrl == url
		return
	}

	return true, nil
}

// Return the URL for the given shortcode.
func LookupUrl(code sting) (url string, err error) {
	return redis.String(conn.Do("GET", codeKey(code)))
}
