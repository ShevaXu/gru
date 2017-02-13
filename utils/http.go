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

// HttpClient is the "reversed" interface for web-utils.SafeClient
// for better unit tests.
type HTTPClient interface {
	RequestWithClose(req *http.Request) (status int, body []byte, err error)
	RequestWithRetry(req *http.Request, maxTries int) (tries, status int, body []byte, err error)
	PostJsonWithRetry(url string, v interface{}, maxTries int, f utils.RequestHook) (tries, status int, body []byte, err error)
	PostFormWithRetry(url string, v url.Values, maxTries int, f utils.RequestHook) (tries, status int, body []byte, err error)
}

// PrintlnJson helps debugging JSON APIs.
func PrintlnJson(v interface{}) {
	bs, _ := json.Marshal(v)
	fmt.Println(string(bs))
}
