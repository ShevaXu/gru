package wikidata

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// wbgetclaims API
const (
	ActionGetClaims = "wbgetclaims"
	PropertyImage   = "P18"
	ImgEndpoint     = "https://upload.wikimedia.org/wikipedia/commons"
)

type DataValueObj struct {
	Value json.RawMessage `json:"value"`
	Type  string          `json:"type"`
}

type EntityID struct {
	EntityType string `json:"entity-type"`
	NumericID  int    `json:"numeric-id"`
	ID         string `json:"id"`
}

type Snak struct {
	Snaktype  string       `json:"snaktype"`
	Property  string       `json:"property"`
	DataValue DataValueObj `json:"datavalue"`
	Datatype  string       `json:"datatype"`
}

type Reference struct {
	Hash       string          `json:"hash"`
	Snaks      map[string]Snak `json:"snaks"`
	SnaksOrder []string        `json:"snaks-order"`
}

type Property struct {
	Mainsnak   Snak        `json:"mainsnak"`
	Type       string      `json:"type"`
	ID         string      `json:"id"`
	Rank       string      `json:"rank"`
	References []Reference `json:"references"`
}

type Claim []Property

type ImageClaim struct {
	Claims struct {
		P18 Claim `json:"P18"`
	} `json:"claims"`
}

// Follow the trick @ http://stackoverflow.com/questions/34393884/how-to-get-image-url-property-from-wikidata-item-by-api
func GetImage(c *Client, entity string) (imageUrl string, err error) {
	v := url.Values{}
	v.Add("entity", entity)
	v.Add("property", PropertyImage)

	body, err := c.Query(v)
	if err != nil {
		return
	}

	var ic ImageClaim
	err = json.Unmarshal(body, &ic)
	if err != nil {
		return "", errors.Wrap(err, "Wikidata image unmarshal")
	}

	// for P18 the RawMessage should a string
	var val string
	err = json.Unmarshal(ic.Claims.P18[0].Mainsnak.DataValue.Value, &val)
	if err != nil || val == "" {
		return val, nil
	}
	val = strings.Replace(val, " ", "_", len(val)) // replace " "
	sum := fmt.Sprintf("%x", md5.Sum([]byte(val)))
	imageUrl = fmt.Sprintf("%s/%c/%s/%s", ImgEndpoint, sum[0], sum[0:2], val)

	return
}
