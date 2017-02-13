package wikidata

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

func IsItem(entity string) bool {
	if entity == "" {
		return false
	} else {
		return entity[0] == "Q"[0]
	}
}

func IsProperty(entity string) bool {
	if entity == "" {
		return false
	} else {
		return entity[0] == "P"[0]
	}
}

/*
 * wbgetentities API https://www.wikidata.org/w/api.php?action=help&modules=wbgetentities
 * @params props:
 * info, sitelinks, sitelinks/urls, aliases, labels, descriptions, claims, datatype
 * (default: info|sitelinks|aliases|labels|descriptions|claims|datatype)
 */
const (
	LangEnglish       = "en"
	QueryIDs          = "ids"
	QueryProps        = "props"
	ActionGetEntities = "wbgetentities"
	BasicsProps       = "labels|descriptions|info|aliases"
)

type StringValue struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

type AliasList []StringValue

type Entity struct {
	Pageid       int                    `json:"pageid"`
	Ns           int                    `json:"ns"`
	Title        string                 `json:"title"`
	Lastrevid    int                    `json:"lastrevid"`
	Modified     time.Time              `json:"modified"`
	Type         string                 `json:"type"`
	ID           string                 `json:"id"`
	Labels       map[string]StringValue `json:"labels"`
	Descriptions map[string]StringValue `json:"descriptions"`
	Aliases      map[string]AliasList   `json:"aliases"`
}

type Basics struct {
	Entities map[string]Entity `json:"entities"`
	Success  int               `json:"success"`
}

func GetBasics(c *Client, entities string) (b Basics, err error) {
	v := url.Values{}
	v.Add(QueryIDs, entities)
	v.Add(QueryProps, BasicsProps)

	body, err := c.Query(v)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &b)
	if err != nil {
		return b, errors.Wrap(err, "Wikidata basics unmarshal")
	}

	return
}
