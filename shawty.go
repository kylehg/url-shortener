package shawty

import (
//	"bytes"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"math"
//	"math/big"
	"strings"
)

// TODO
var network = ""
var address = ""
var conn, err = redis.Dial(network, address)

func codeKey(code string) string {
	return "code:" + code
}

const URL_COUNT_KEY = "global:count"

func genShortcode(n int) string {
	const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-"
	const BASE = 64

	n += int(math.Pow(BASE, 3))
	code := make([]string, 0)
	for n > 0 {
		r := n % BASE
		n = n / BASE
		code = append(code, string(ALPHABET[r]))
		fmt.Printf("r=%d, n=%d, code=%s\n", r, n, code)
	}
	for i, j := 0, len(code)-1; i < j; i, j = i+1, j-1 {
		code[i], code[j] = code[j], code[i]
	}
	return strings.Join(code, "")
}

// Shorten the given URL and return the shortcode
func ShortenUrl(url string) (code string, err error) {
	var code string

	for exists := true; exists; {
		// Get the newest unique integer
		n, err := redis.Int(conn.Do("INCR", URL_COUNT_KEY))
		if err != nil {
			return
		}

		// Generate its shortcode and make sure its not taken
		code, err := genShortcode(n)

		exists, err := redis.Bool(conn.DO("EXISTS", codeKey(code)))
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
