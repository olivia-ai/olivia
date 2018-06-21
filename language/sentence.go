package language

import (
	"regexp"
	"strings"
)

type Sentence struct {
	Content string
}

// Returns an array of tokenized words
func (sentence Sentence) Tokenize() []string {
	regex, err := regexp.Compile(
		strings.Replace(
			"(\\d+(\\.\\d+)?\\s*%)|(${LETTERS}')|(${LETTERS}+)|([^${LETTERS}\\s])",
			"${LETTERS}",
			"[a-zàâçéèêëîïôùûœæA-ZÀÂÇÉÈÊËÎÏÔÙÛŒÆ0-9-]",
			3))

	if err != nil {
		panic(err)
	}

	return regex.FindAllString(sentence.Content, -1)
}
