package shawty

import "github.com/garyburd/redigo/redis"

const (
	NETWORK                    = "tcp"
	ADDRESS                    = ":6379"
	DEFAULT_DB                 = 0
	TEST_DB                    = 2
	REDIS_KEY_PREFIX           = "shawty:"
	REDIS_URL_KEY_PREFIX       = REDIS_KEY_PREFIX + "url:"
	REDIS_SHORTCODE_KEY_PREFIX = REDIS_KEY_PREFIX + "code:"
)

var redisConn

// Get the Redis connection
func getConn() redis.Conn {
	var err
	if redisConn != nil {
		if redisConn, err = redis.Dial(NETWORK, ADDRESS); err != nil {
			panic("Error connecting to Redis database: " + err.Error())
		}
		if _, err = conn.Do("SELECT", DEFAULT_DB); err != nil {
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

func redisSetIfNotExists(key string, val string) (bool, error) {
	return redis.Bool(getConn().Do("SET", key, val, "NX"))
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
	wasSet, err := redisSetIfNotExists(codeKey(code), url)
	if err != nil {
		return err
	}

	if !wasSet {
		existingUrl, err := GetUrl(code)
		if err != nil {
			return err
		} else if existingUrl == url {
			return nil
		}
		return // TODO ERROR code is already taken
	}

	return nil
}

// Set the defua for a given shortcode it the code doesn't already exist
func SetDefaultCode(url string, code string) error {
	conn := getConn()

	conn.Send("MULTI")
	conn.Send("SET", codeKey(code), url, "NX")
	conn.Send("SET", urlKey(url), code, "NX")
	responses, err := conn.DO("EXEC")
	if err != nil {
		return err
	}

	// Check if default code is used
	// If used and maps to same url, return nil
	// If used and maps to different url, return error
	existingUrl, err := GetUrl(code)
	if err != redis.ErrNil {
		if existingUrl == url {
			return nil
		}
		return err
	}

	// Set code -> url
	// If fails, return error
	wasSet, err := redisSetIfNotExists(codeKey(code), url)
	if !wasSet {
		return // TODO Race error
	} else if err != nil {
		return err
	}

	// Set url -> code
	// If fails:
	// - unset code -> url
	// - return error
	wasSet, err := redisSetIfNotExists(urlKey(url), code)
	if !wasSet || err != nil {
		// TODO
		return err
	}

	// Else return nil
}

// Get a random, unused shortcode for a URL
func CreateRandomShortcode() (string, error) {
	// TODO
	return "", nil
}
