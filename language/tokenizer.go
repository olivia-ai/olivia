package language

import (
	"regexp"
	"strings"
)

type Sentence struct {
	Content string
}

const frenchChars = "[a-zàâçéèêëîïôùûœæA-ZÀÂÇÉÈÊËÎÏÔÙÛŒÆ0-9-]"

// Returns an array of tokenized words
func (sentence Sentence) Tokenize() []string {
	regex, err := regexp.Compile(
		strings.Replace(
			"(\\d+(\\.\\d+)?\\s*%)|(${LETTERS}')|(${LETTERS}+)|([^${LETTERS}\\s])",
			"${LETTERS}",
			frenchChars,
			3))

	if err != nil {
		panic(err)
	}

	return regex.FindAllString(sentence.Content, -1)
}

