package core_test

import (
	//"os"
	"encoding/json"
	"io/ioutil"
	"net/url"
	"testing"

	. "github.com/ShevaXu/gru/core"
	"github.com/ShevaXu/gru/utils"
)

func TestGoogleSearch(t *testing.T) {
	assert := utils.NewAssert(t)

	// TODO: cannot reach google so far (curl + proxy works tho)
	//key := os.Getenv("GOOGLE_KEY")
	//engineID := os.Getenv("ENGINE_ID")
	//if key == "" || engineID == "" {
	//	t.Skip("Skipped! Must set GOOGLE_KEY & ENGINE_ID")
	//}
	//gs := NewGoogleSearch(key, engineID, utils.DefaultClient)
	//res, err := gs.Search("apple")

	sample, _ := ioutil.ReadFile("./apple.json")

	var res SearchResult
	err := json.Unmarshal(sample, &res)
	assert.NoError(err, "Unmarshal result")
	assert.Equal(10, len(res.Items), "Should return 10 results")

	if len(res.Items) > 0 {
		item := res.Items[0]
		assert.True(item.Title != "", "Result has title")
		_, err = url.Parse(item.Link)
		assert.NoError(err, "Result has valid link")
	}
}
