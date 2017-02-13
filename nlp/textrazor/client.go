package textrazor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/ShevaXu/gru/utils"
)

const (
	endpoint   = "https://api.textrazor.com" // use the secure HTTPS endpoint to protect your credentials
	authHeader = "x-textrazor-key"
	maxRetries = 2 // do not overload it
)

// Client wraps the API key and makes queries.
type Client struct {
	key string
	cl  utils.HTTPFormPoster
}

// Query makes request(s) to the TextRazor secure endpoint;
// it handles all the http stuffs including error-retry.
func (c *Client) Query(extractors, text string) (res Result, err error) {
	if text == "" {
		return res, errors.New("TextRazor query empty text")
	}

	payload := url.Values{}
	payload.Add("extractors", extractors)
	payload.Add("text", text)
	_, status, body, err := c.cl.PostFormWithRetry(endpoint, payload, maxRetries, func(r *http.Request) {
		r.Header.Add(authHeader, c.key)
	})

	if err != nil {
		goto fail
	}
	if status != http.StatusOK {
		return res, errors.New(fmt.Sprintf("TextRazor query %d: %s", status, string(body)))
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		goto fail
	}

	// success
	return

fail:
	err = errors.Wrap(err, "TextRazor query")
	return
}

// NewClient returns a functional TextRazor Client;
// the key cannot be changed once set.
func NewClient(key string, cl utils.HTTPFormPoster) *Client {
	return &Client{
		key: key,
		cl:  cl,
	}
}
