package shawty

import (
	"fmt"
//	"math"
	"strings"
)

func genShortcode(n int) string {
	const ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789"
	const BASE = 63

	code := make([]string, 0)
	
	if r == 1 {
		code = append(code, "a")
	}
	for n > 0 {
		r := n % BASE
		n = n / BASE
		char := ALPHABET[r]
		code = append(code, string(char))
		fmt.Printf("r=%d, n=%d, code=%s\n", r, n, code)
		r = n % BASE
	}

//	n += int(math.Pow(BASE, 3))
//	code := make([]string, 0)
//	for n > 0 {
//		r := n % BASE
//		n = n / BASE
//		char := ALPHABET[r]
//		if n == 0 {
//			char = ALPHABET[r-1]
//		}
//		code = append(code, string(char))
//		fmt.Printf("r=%d, n=%d, code=%s\n", r, n, code)
//	}

	// Reverse the built code
	for i, j := 0, len(code)-1; i < j; i, j = i+1, j-1 {
		code[i], code[j] = code[j], code[i]
	}
	res := strings.Join(code, "")
	fmt.Printf("-> %s\n", res)
	return res
}

// func main() {
//     fmt.Println(genShortcode(0))
// }

// Shorten the given URL and return the shortcode
// func ShortenUrl(url string) (code string, err error) {
// 	var code string
// 
// 	for exists := true; exists; {
// 		// Get the newest unique integer
// 		n, err := redis.Int(conn.Do("INCR", URL_COUNT_KEY))
// 		if err != nil {
// 			return
// 		}
// 
// 		// Generate its shortcode and make sure its not taken
// 		code, err := genShortcode(n)
// 
// 		exists, err := redis.Bool(conn.DO("EXISTS", codeKey(code)))
// 		if err != nil {
// 			return
// 		}
// 	}
// 
// 	return
// }
// 
// // Shorten the given URL to the given code. Return true if the URL was
// // successfully saved to the code, false if the code was taken.
// func ShortenUrlToCode(url string, code string) (success bool, err error) {
// 	success := false
// 
// 	// Attempt to set the shortcode
// 	wasSet, err := redis.Bool(conn.Do("SETNX", codeKey(code), url))
// 	if err != nil {
// 		return
// 	}
// 
// 	if !wasSet {
// 		setUrl, err := conn.Do("GET", codeKey(code))
// 		success = setUrl == url
// 		return
// 	}
// 
// 	return true, nil
// }
// 
// // Return the URL for the given shortcode.
// func LookupUrl(code sting) (url string, err error) {
// 	return redis.String(conn.Do("GET", codeKey(code)))
// }
