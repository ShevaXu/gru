package wikidata

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/ShevaXu/gru/utils"
)

const (
	Endpoint   = "https://www.wikidata.org/w/api.php"
	FormatJson = "json"
	retries    = 2
)

type Client struct {
	action    string
	format    string
	languages string
	cl        utils.HTTPRequester
}

func (c *Client) Action(action string) *Client {
	rtn := *c
	rtn.action = action
	return &rtn
}

func (c *Client) Languages(lang string) *Client {
	rtn := *c
	rtn.languages = lang
	return &rtn
}

func (c *Client) Query(v url.Values) ([]byte, error) {
	// auto-fill
	v.Add("action", c.action)
	v.Add("format", c.format)
	v.Add("languages", c.languages)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", Endpoint, v.Encode()), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Wikidata new request")
	}

	_, status, body, err := c.cl.RequestWithRetry(req, retries)
	if err != nil {
		return nil, errors.Wrap(err, "Wikidata query")
	}
	if status != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Wikidata query %d: %s", status, string(body)))
	}

	return body, nil
}

func NewClient(cl utils.HTTPRequester) *Client {
	return &Client{
		format: FormatJson, // just JSON for now
		cl:     cl,
	}
}
