Shawty: The Hiptser URL Shortener
=================================

A URL shortener written in Go. Because why not.


API
---

### Routes

#### `/`

##### GET

The main page, and HTML app.

##### POST

**Parameters**

- `url`: A URL to shorten.
- `code` _(optional)_: A code to shorten it too.

**Response**

- **201 Created**: The URL+code pair (or just the URL, if no code was provided) had not been used before and was created.
  - `url`
  - `code`
- **200 OK**: The URL already existed was the previously contained code is returned
  - `url`
  - `code`


