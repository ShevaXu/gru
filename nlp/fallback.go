package nlp

import (
	"github.com/mvryan/fasttag"
)

type Word struct {
	Position     int    `json:"position"`
	Token        string `json:"token"`
	PartOfSpeech string `json:"partOfSpeech"`
}

type Sentence struct {
	Pretty string `json:"pretty"`
	Words  []Word `json:"words"`
}

func Parse(s string) Sentence {
	tokens := fasttag.WordsToSlice(s)
	tags := fasttag.BrillTagger(tokens)
	n := len(tokens)
	sentence := Sentence{
		Pretty: "$",
		Words:  make([]Word, n),
	}
	for i, t := range tokens {
		sentence.Words[i] = Word{
			Position:     i,
			Token:        t,
			PartOfSpeech: tags[i],
		}
		sentence.Pretty += (" " + t + "/" + tags[i])
	}
	return sentence
}
