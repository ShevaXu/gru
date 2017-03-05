package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/ShevaXu/gru/utils"
)

const (
	googleCustomSearchEndpoint = "https://www.googleapis.com/customsearch/v1"
	googleCustomSearchMethod   = "GET"

	googleKeyParam        = "key"
	googleEIDParam        = "cx"
	googleQueryParam      = "q"
	googleFieldsParams    = "fields"
	googleStartIndexParam = "start"

	// for unified search result, *UNSTABLE*
	googleLiteFields = "items(title,link,snippet)"

	googleTries = 3
)

type SearchItem struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Snippet string `json:"snippet"`
}

type SearchResult struct {
	Items []SearchItem `json:"items"`
}

type Searcher interface {
	Search(string) SearchResult
}

type GoogleSearch struct {
	key      string
	engineID string
	cl       utils.HTTPRequester
}

func (g *GoogleSearch) Search(q string) (res SearchResult, err error) {
	v := url.Values{}
	v.Add(googleKeyParam, g.key)
	v.Add(googleEIDParam, g.engineID)
	v.Add(googleQueryParam, q)
	v.Add(googleFieldsParams, googleLiteFields)
	req, _ := http.NewRequest(googleCustomSearchMethod, fmt.Sprintf("%s?%s", googleCustomSearchEndpoint, v.Encode()), nil)
	_, status, body, err := g.cl.RequestWithRetry(req, googleTries)

	if err != nil {
		goto fail
	}
	if status != http.StatusOK {
		return res, errors.New(fmt.Sprintf("Google search %d: %s", status, string(body)))
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		goto fail
	}

	// success
	return

fail:
	err = errors.Wrap(err, "Google search")
	return
}

func NewGoogleSearch(key, engineID string, cl utils.HTTPRequester) *GoogleSearch {
	return &GoogleSearch{
		key:      key,
		engineID: engineID,
		cl:       cl,
	}
}
