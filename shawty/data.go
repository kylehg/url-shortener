package shawty

import "github.com/garyburd/redigo/redis"

const (
	NETWORK = "tcp"
	ADDRESS = ":6379"
	REDIS_KEY_PREFIX = "shawty:"
	REDIS_URL_KEY_PREFIX = REDIS_KEY_PREFIX + "url:"
	REDIS_SHORTCODE_KEY_PREFIX = REDIS_KEY_PREFIX + "code:"
)

var conn

// Get the Redis connection
func getConn() redis.Conn {
	var err
	if conn != nil {
		conn, err = redis.Dial(NETWORK, ADDRESS)
		if err != nil {
			panic("Error connecting to Redis database: " + err.Error())
		}
	}
	return conn
}

// Convert a shortcode to a shotcode key for Redis
func codeKey(code string) string {
	return REDIS_SHORTCODE_KEY_PREFIX + code
}

// Convert a URL to a URL key for Redis
func urlKey(url string) string {
	return REDIS_URL_KEY_PREFIX + url
}

// Lookup the URL for a given shortcode
func GetUrlFromCode(code string) (string, error) {
	return redis.String(getConn().Do("GET", codeKey(code)))
}

// Lookup the shortcode for a given URL
func GetCodeFromUrl(url string) (string, error) {
	return redis.String(getConn().Do("GET", urlKey(url)))
}

// Set the URL for a given shortcode it the code doesn't already exist
func SetUrlForCode(code string, url string) error {
	conn := getConn()
	existingUrl, err := GetUrlFromCode(code)
	if url == existingUrl {
		return nil
	}
	if err != redis.ErrNil {
		return // TODO return error "Code %s already maps to URL %s"
	}

	wasSet, err := redis.Bool(conn.Do("SETNX", codeKey(code), url))
	if err != nil {
		return err
	}

	urlMapped, err := redis.Bool(conn.Do("SETNX", urlKey(url), code))
	if !urlMapped || err != nil {
		return // TODO return error "URL %salready mapped by code %s" (may be race condition)
	}

	return nil
}

// Get a random, unused shortcode for a URL
func CreateRandomShortcode() (string, error) {
	// TODO
	return "", nil
}