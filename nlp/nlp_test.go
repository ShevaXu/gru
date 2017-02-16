package nlp_test

import (
	"testing"

	. "github.com/ShevaXu/gru/nlp"
	"github.com/ShevaXu/gru/utils"
)

var testSentence = "I will go to Beijing to play basketball next monday"

func TestParse(t *testing.T) {
	assert := utils.NewAssert(t)

	res := Parse(testSentence)

	// textrazor outputs: $ I/PRP will/MD go/VB to/TO beijing/NNP to/TO play/VB basketball/NN next/IN monday/NNP
	oneLine := "$ I/NN will/MD go/VB to/TO Beijing/VBG to/TO play/VB basketball/NN next/JJ monday/NN"
	assert.Equal(oneLine, res.Pretty, "One liner works")

	assert.Equal("will", res.Words[1].Token, "tokenized")
	assert.Equal(POS_MD, res.Words[1].PartOfSpeech, "will is MD")
}
