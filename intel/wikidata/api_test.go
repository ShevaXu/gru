package wikidata_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	. "github.com/ShevaXu/gru/intel/wikidata"
	"github.com/ShevaXu/gru/utils"
)

const testEntityID = "Q5372"

func TestEntityID(t *testing.T) {
	assert := utils.NewAssert(t)

	assert.Equal(false, IsItem("") || IsProperty(""), `"" is not item nor property`)
	assert.Equal(true, IsItem("Q5372"), "Q5372 is an item")
	assert.Equal(true, IsProperty("P18"), "P18 is a property")
}

func TestClient(t *testing.T) {
	assert := utils.NewAssert(t)

	c := NewClient(utils.DefaultClient)
	ac := c.Action(ActionGetEntities)
	ec := c.Languages(LangEnglish)

	addr := utils.GetAddress(c)
	assert.NotEqual(addr, utils.GetAddress(ac), "Chained call returns new instance")
	assert.NotEqual(addr, utils.GetAddress(ec), "Chained call returns new instance")
}

type localJsonStub struct {
	retries int
	status  int
	body    []byte
	err     error
}

func (c *localJsonStub) RequestWithRetry(req *http.Request, maxTries int) (tries, status int, body []byte, err error) {
	return c.retries, c.status, c.body, c.err
}

func TestClient_Query(t *testing.T) {
	assert := utils.NewAssert(t)
	v := url.Values{}
	v.Add(QueryIDs, testEntityID)
	v.Add(QueryProps, "datatype") // this should be fixed

	// error handling
	errC := NewClient(&localJsonStub{}) // status = 0
	_, err := errC.Query(v)
	assert.NotNil(err, "Stub returns error")

	// -short set
	if testing.Short() {
		t.Skip("Skip real network query")
	}
	c := NewClient(utils.DefaultClient).Action(ActionGetEntities).Languages(LangEnglish)
	body, err := c.Query(v)
	assert.NoError(err, "Query should have no error")
	assert.Equal([]byte(`{"entities":{"Q5372":{"type":"item","id":"Q5372"}},"success":1}`), body, "Query returns []byte")
}

func TestGetBasics(t *testing.T) {
	assert := utils.NewAssert(t)

	// 2017-02-13:
	// https://www.wikidata.org/w/api.php?action=wbgetentities&ids=Q5372&format=json&languages=en&props=labels|descriptions|info|aliases
	basics, _ := ioutil.ReadFile("./basics.json")
	c := NewClient(&localJsonStub{
		status: http.StatusOK,
		body:   basics,
	})

	b, err := GetBasics(c, testEntityID)
	assert.NoError(err, "GetBasics query no error")

	entity, ok := b.Entities[testEntityID]
	if !ok {
		t.Error("Should get entity", testEntityID)
	}

	assert.Equal("item", entity.Type, "Type got")
	assert.Equal("basketball", entity.Labels[LangEnglish].Value, "Labels got")
	assert.Equal("team sport played on an indoor court with baskets on either end", entity.Descriptions[LangEnglish].Value, "Description got")
	assert.Equal("basket ball", entity.Aliases[LangEnglish][0].Value, "Aliases got")

	errC := NewClient(&localJsonStub{
		body: []byte("Invalid JSON"),
	})
	_, err = GetBasics(errC, testEntityID)
	assert.NotNil(err, "Stub returns error")
}
