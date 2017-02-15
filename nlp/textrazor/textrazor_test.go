package textrazor_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/ShevaXu/gru/nlp"
	"github.com/ShevaXu/gru/nlp/textrazor"
	"github.com/ShevaXu/gru/utils"
)

func TestAPI(t *testing.T) {
	assert := utils.NewAssert(t)

	// 2017-02-13:
	// extractors=entities,relations,words"
	// text=I will go to beijing to play basketball next monday
	sample, _ := ioutil.ReadFile("./sample.json")

	var res textrazor.Result
	err := json.Unmarshal(sample, &res)
	assert.NoError(err, "Unmarshal result")

	assert.Equal("go", res.Resp.Sentences[0].Words[2].Token, "Parse sentence to tokens")
	assert.Equal(nlp.POS_VB, res.Resp.Sentences[0].Words[2].PartOfSpeech, "Parse sentence POS")
	assert.Equal(2, res.Resp.Sentences[0].Words[3].ParentPosition, "Parse sentence relation")

	assert.Equal("City", res.Resp.Entities[0].Type[3], "Parse entity type")
	assert.Equal("Beijing", res.Resp.Entities[0].EntityID, "Parse entity ID")
	assert.Equal("http://en.wikipedia.org/wiki/Beijing", res.Resp.Entities[0].WikiLink, "Parse entity wikipedia link")
	assert.Equal("Q956", res.Resp.Entities[0].WikidataID, "Parse entity wikidata ID")

	// 2017-02-20T00:00:00.000Z
	day, err := time.Parse("2006-01-02T00:00:00.000Z", res.Resp.Entities[2].EntityID)
	assert.NoError(err, "Time parsed for entity")
	assert.Equal(20, day.Day(), "Time parsed")

	// one liner
	assert.Equal("$ I/PRP will/MD go/VB to/TO beijing/NNP to/TO play/VB basketball/NN next/IN monday/NNP ./.",
		res.Resp.Sentences[0].OneLine(), "Pretty one line")
}

// TODO: use HTTPClient stub instead of DefaultClient.
func TestClient_Query(t *testing.T) {
	token := os.Getenv("TEXTRAZOR_TOKEN")
	if token == "" {
		t.Skip("No API token")
	}

	cl := textrazor.NewClient(token, utils.DefaultClient)
	res, err := cl.Query(textrazor.EntitiesWordsRelations, "I will go to beijing to play basketball next monday")
	if err != nil {
		t.Error(err)
	}
	utils.PrintlnJson(res)
}
