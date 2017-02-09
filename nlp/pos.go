package nlp

// Penn Treebank part-of-speech tags in alphabetical order.
// From https://cs.nyu.edu/grishman/jet/guide/PennPOS.html
// $ awk '{printf "%s","POS_"$2"=\""$2"\"" "// "; $1=$2=""; print $0}' pos.txt
const (
	POS_CC   = "CC"   //   Coordinating conjunction
	POS_CD   = "CD"   //   Cardinal number
	POS_DT   = "DT"   //   Determiner
	POS_EX   = "EX"   //   Existential there
	POS_FW   = "FW"   //   Foreign word
	POS_IN   = "IN"   //   Preposition or subordinating conjunction
	POS_JJ   = "JJ"   //   Adjective
	POS_JJR  = "JJR"  //   Adjective, comparative
	POS_JJS  = "JJS"  //   Adjective, superlative
	POS_LS   = "LS"   //   List item marker
	POS_MD   = "MD"   //   Modal
	POS_NN   = "NN"   //   Noun, singular or mass
	POS_NNS  = "NNS"  //   Noun, plural
	POS_NNP  = "NNP"  //   Proper noun, singular
	POS_NNPS = "NNPS" //   Proper noun, plural
	POS_PDT  = "PDT"  //   Predeterminer
	POS_POS  = "POS"  //   Possessive ending
	POS_PRP  = "PRP"  //   Personal pronoun
	POS_PRPo = "PRP$" //   Possessive pronoun
	POS_RB   = "RB"   //   Adverb
	POS_RBR  = "RBR"  //   Adverb, comparative
	POS_RBS  = "RBS"  //   Adverb, superlative
	POS_RP   = "RP"   //   Particle
	POS_SYM  = "SYM"  //   Symbol
	POS_TO   = "TO"   //   to
	POS_UH   = "UH"   //   Interjection
	POS_VB   = "VB"   //   Verb, base form
	POS_VBD  = "VBD"  //   Verb, past tense
	POS_VBG  = "VBG"  //   Verb, gerund or present participle
	POS_VBN  = "VBN"  //   Verb, past participle
	POS_VBP  = "VBP"  //   Verb, non-3rd person singular present
	POS_VBZ  = "VBZ"  //   Verb, 3rd person singular present
	POS_WDT  = "WDT"  //   Wh-determiner
	POS_WP   = "WP"   //   Wh-pronoun
	POS_WPo  = "WP$"  //   Possessive wh-pronoun
	POS_WRB  = "WRB"  //   Wh-adverb
)
