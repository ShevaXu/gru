package textrazor

const (
	EntitiesExtractor  = "entities"
	WordsExtractor     = "words"
	RelationsExtractor = "relations"

	// for now just these three
	EntitiesWordsRelations = "entities,words,relations"
)

// JSON response structs for the "entities,words" extractors;
// with the help from https://mholt.github.io/json-to-go/.

type Word struct {
	Position         int    `json:"position"`
	StartingPos      int    `json:"startingPos"`
	EndingPos        int    `json:"endingPos"`
	Stem             string `json:"stem"`
	Lemma            string `json:"lemma"`
	Token            string `json:"token"`
	PartOfSpeech     string `json:"partOfSpeech"`
	ParentPosition   int    `json:"parentPosition,omitempty"`
	RelationToParent string `json:"relationToParent,omitempty"`
}

type Sentense struct {
	Position int    `json:"position"`
	Words    []Word `json:"words"`
}

type Entity struct {
	ID              int      `json:"id"`
	Type            []string `json:"type,omitempty"`
	MatchingTokens  []int    `json:"matchingTokens"`
	EntityID        string   `json:"entityId"`
	FreebaseTypes   []string `json:"freebaseTypes,omitempty"`
	ConfidenceScore float64  `json:"confidenceScore"`
	WikiLink        string   `json:"wikiLink"`
	MatchedText     string   `json:"matchedText"`
	FreebaseID      string   `json:"freebaseId,omitempty"`
	RelevanceScore  int      `json:"relevanceScore"`
	EntityEnglishID string   `json:"entityEnglishId"`
	StartingPos     int      `json:"startingPos"`
	EndingPos       int      `json:"endingPos"`
	WikidataID      string   `json:"wikidataId,omitempty"`
}

type Param struct {
	Relation      string `json:"relation"`
	WordPositions []int  `json:"wordPositions"`
}

type Relation struct {
	ID            int     `json:"id"`
	WordPositions []int   `json:"wordPositions"`
	Params        []Param `json:"params"`
}

type Property struct {
	ID                int   `json:"id"`
	WordPositions     []int `json:"wordPositions"`
	PropertyPositions []int `json:"propertyPositions"`
}

type Response struct {
	Sentences          []Sentense `json:"sentences"`
	Language           string     `json:"language"`
	LanguageIsReliable bool       `json:"languageIsReliable"`
	Entities           []Entity   `json:"entities"`
	Relations          []Relation `json:"relations"`
	Properties         []Property `json:"properties"`
}

type Result struct {
	Resp Response `json:"response"`
	Time float64  `json:"time"`
	Ok   bool     `json:"ok"`
}

// Body is the JSON request payload.
type Body struct {
	Extractors string `json:"extractors"`
	Text       string `json:"text"`
}
