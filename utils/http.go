package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	utils "github.com/ShevaXu/web-utils"
)

// DefaultClient is a web-utils.SafeClient instance
// for the whole project's usage.
var DefaultClient = utils.StdClient()

// HTTPRequester is one of the "reversed" interface of
// web-utils.SafeClient for better unit tests;
// use it for GET requests.
type HTTPRequester interface {
	RequestWithRetry(req *http.Request, maxTries int) (tries, status int, body []byte, err error)
}

// HTTPJsonPoster is one of the "reversed" interface of
// web-utils.SafeClient for better unit tests;
// use it for POST requests with JSON payload.
type HTTPJsonPoster interface {
	PostJsonWithRetry(url string, v interface{}, maxTries int, f utils.RequestHook) (tries, status int, body []byte, err error)
}

// HTTPFormPoster is one of the "reversed" interface of
// web-utils.SafeClient for better unit tests;
// use it for POST requests (default to plain/text).
type HTTPFormPoster interface {
	PostFormWithRetry(url string, v url.Values, maxTries int, f utils.RequestHook) (tries, status int, body []byte, err error)
}

// HTTPClient wraps HTTP* interfaces.
type HTTPClient interface {
	HTTPRequester
	HTTPJsonPoster
	HTTPFormPoster
}

// PrintlnJson helps debugging JSON APIs.
func PrintlnJson(v interface{}) {
	bs, _ := json.Marshal(v)
	fmt.Println(string(bs))
}
