package analysis

import (
	"../slice"
	"github.com/neurosnap/sentences"
	"strings"
)

type Sentence struct {
	Content string
}

// Returns an array of words which are tokenized with natural processing
func (sentence Sentence) Tokenize() (tokenizedWords []string) {
	tokenizer := sentences.NewWordTokenizer(sentences.NewPunctStrings())
	tokens := tokenizer.Tokenize(sentence.Content, false)

	// Initialize an array of ignored characters
	ignoredChars := []string{"?", "-"}

	// Get the string token and add it to tokenizedWords
	for _, tokenizedWord := range tokens {
		word := strings.ToLower(tokenizedWord.Tok)

		// Remove all ignored characters from the word
		for _, ignoredChar := range ignoredChars {
			word = strings.Replace(word, ignoredChar, "", -1)
		}

		tokenizedWords = append(tokenizedWords, word)
	}

	return tokenizedWords
}

// Retrieves all the intents words and returns the bag of words of the Sentence content
func (sentence Sentence) WordsBag(words []string) (bag []float64) {
	for _, word := range words {
		// Append 1 if the patternWords contains the actual word, else 0
		var valueToAppend float64 = 0
		if slice.Contains(sentence.Tokenize(), word) {
			valueToAppend = 1
		}

		bag = append(bag, valueToAppend)
	}

	return bag
}
