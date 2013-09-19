Shawty: The Hiptser URL Shortener
=====

A URL shortener written in Go, using Redis. Because why not.


API
-----

Unless otherwise specified, API responses are encoded as JSON and take the form:

    {
      data: {
        url: "http://google.com/",
        code: "ggl"
      },
      meta: {
        response_code: 200,
        response_text: "Created",
        error_text: "",
      }
    }

### `/`

#### GET

The main HTML app.

#### POST

**Parameters**

- `url`: A URL to shorten.

**Response**

- **200 OK**: The URL is already mapped by a code, which is returned as `code`.
- **201 Created**: The URL was successfully shortened to a new short code, returned as `code`.


### '/:shortcode'

#### GET

Lookup the URL at the given shortcode.

**Response**

- **302 Found**: The shortcode maps to the redirecting URL.
- **404 Not Found**: There is no URL for the given shortcode. Returns the HTML app.

#### PUT + POST

Save a new URL to the given shortcode.

**Parameters**

- `url`: A URL to shorten

**Response**

- **200 OK**: The shortcode already maps to that URL.

- **201 Created**: The URL was successfully shortened to the given short code, which is returned.

- **303 See Other**: The given shortcode is already taken, by the returned `url`.

