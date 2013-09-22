package shawty

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"strings"
	"time"
)

const NETWORK = "tcp"
const ADDRESS = ":6379"

const ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const ALPHABET_LEN = len(ALPHABET)
const DEFAULT_LEN = 5

var conn, err = redis.Dial(NETWORK, ADDRESS)
// TODO: Better way to refer to connection
// if err != nil {
// 	panic("Error connecting to Redis database: " + err.Error())
// }

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func ShortenUrl(w http.ResponseWriter, r *http.Request, code string) (code string, err error) {
}


// Shorten the given URL and return the shortcode
func ShortenUrl(url string, code string) (code string, err error) {
	// First, check that URL isn't already mapped by a random shortcode
	code, err = redis.String(conn.Do("GET", urlKey(url)))
	if err != redis.ErrNil {
		return
	}

	// If not, try generating shortcodes until we find one that's not taken
	for wasSet := false; wasSet == false; {
		code = genRandShortcode(DEFAULT_LEN)
		wasSet, err = redis.Bool(conn.Do("SETNX", codeKey(code), url))
		if err != nil {
			return
		}
	}

	// Set that the URL is mapped by the code
	var urlMapped bool
	urlMapped, err = redis.Bool(conn.Do("SETNX", urlKey(url), code))
	if urlMapped {
		return "", errors.New(fmt.Sprintf("Attempted to add already-mapped URL %s to %s", url, code))
	}

	return
}


// Shorten the given URL to the given code. Return true if the URL was
// successfully saved or already mapped to the URL, false if it was taken
func ShortenUrlToCode(url string, code string) (success bool, err error) {
	// Attempt to set the shortcode
	success, err := redis.Bool(conn.Do("SETNX", codeKey(code), url))
	if err != nil {
		return
	}

	// If failed, check that the code doesn't already map to the shortcode
	if !success {
		var setUrl string
		setUrl, err = redis.String(conn.Do("GET", codeKey(code)))
		success = setUrl == url
		return
	}

	return
}


// Return the URL for the given shortcode.
func LookupUrl(code string) (url string, err error) {
	return redis.String(conn.Do("GET", codeKey(code)))
}


// Generate a string of n random chars from ALPHABET
func genRandShortcode(n int) string {
	code := make([]string, n)
	for i := 0; i < n; i++ {
		code[i] = string(ALPHABET[r.Intn(ALPHABET_LEN)])
	}
	res := strings.Join(code, "")
	fmt.Println("Random code " + res)
	return res
}


// Convert a shortcode to a shotcode key for Redis
func codeKey(code string) string {
	return "code:" + code
}


// Convert a URL to a URL key for Redis
func urlKey(code string) string {
	return "url:" + code
}
