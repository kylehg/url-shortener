package shawty

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

const (
	NETWORK                    = "tcp"
	ADDRESS                    = ":6379"
	DEFAULT_DB                 = 0
	TEST_DB                    = 2
	REDIS_KEY_PREFIX           = "shawty:"
	REDIS_URL_KEY_PREFIX       = REDIS_KEY_PREFIX + "url:"
	REDIS_SHORTCODE_KEY_PREFIX = REDIS_KEY_PREFIX + "code:"
)

var redisConn redis.Conn

// Get the Redis connection
func getConn() redis.Conn {
	var err error
	if redisConn == nil {
		if redisConn, err = redis.Dial(NETWORK, ADDRESS); err != nil {
			panic("Error connecting to Redis database: " + err.Error())
		}
		if _, err = redisConn.Do("SELECT", DEFAULT_DB); err != nil {
			panic("Cannot select default database")
		}
	}
	return redisConn
}

// Convert a shortcode to a shotcode key for Redis
func codeKey(code string) string {
	return REDIS_SHORTCODE_KEY_PREFIX + code
}

// Convert a URL to a URL key for Redis
func urlKey(url string) string {
	return REDIS_URL_KEY_PREFIX + url
}

func redisGet(key string) (string, error) {
	return redis.String(getConn().Do("GET", key))
}

// Lookup the URL for a given shortcode
func GetUrl(code string) (string, error) {
	return redisGet(codeKey(code))
}

// Lookup the default (random) shortcode for a given URL
func GetDefaultCode(url string) (string, error) {
	return redisGet(urlKey(url))
}

// Sets a custom, nondefault shortcode for a given URL
func SetCustomCode(url string, code string) error {
	resp, err := redis.String(getConn().Do("SET", codeKey(code), url, "NX"))
	if err != nil {
		return err
	}

	if resp != "OK" {
		return fmt.Errorf("Shortcode %s is already mapped to a URL", code)
	}

	return nil
}

// Set the defualt for a given shortcode it the code doesn't already exist
func SetDefaultCode(url string, code string) error {
	conn := getConn()

	conn.Send("MULTI")
	conn.Send("SET", codeKey(code), url, "NX")
	conn.Send("SET", urlKey(url), code, "NX")
	responses, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		return err
	}

	codeWasSet := responses[0] == "OK"
	urlWasSet := responses[1] == "OK"
	if codeWasSet && urlWasSet {
		return nil
	}

	// Undo partial set
	if codeWasSet && !urlWasSet {
		conn.Do("DEL", codeKey(code))
	}
	if !codeWasSet && urlWasSet {
		conn.Do("DEL", urlKey(url))
	}

	return fmt.Errorf("Failed to set default shortcode %s for %s", code, url)
}
