package analysis

import (
	"regexp"
	"strings"

	"github.com/caneroj1/stemmer"
	"github.com/neurosnap/sentences"
	"github.com/olivia-ai/olivia/util"
)

// Arrange check the format of a string to normalize it, put the string to
// lower case, remove ignored characters
func Arrange(text string) string {
	// Remove punctuation after letters
	punctuationRegex := regexp.MustCompile(`[a-zA-Z]( )?(\.|\?|!)`)
	text = punctuationRegex.ReplaceAllStringFunc(text, func(s string) string {
		punctuation := regexp.MustCompile(`(\.|\?|!)`)
		return punctuation.ReplaceAllString(s, "")
	})

	text = strings.ToLower(text)
	return strings.TrimSpace(text)
}

// Tokenize returns the sentence split in stemmed words
func (sentence Sentence) Tokenize() (tokenizedWords []string) {
	tokenizer := sentences.NewWordTokenizer(sentences.NewPunctStrings())
	tokens := tokenizer.Tokenize(strings.TrimSpace(sentence.Content), false)

	// Get the string token and push it to tokenizedWords
	for _, tokenizedWord := range tokens {
		word := stemmer.Stem(tokenizedWord.Tok)
		tokenizedWords = append(tokenizedWords, word)
	}

	return tokenizedWords
}

// WordsBag retrieves the intents words and returns the sentence converted in a bag of words
func (sentence Sentence) WordsBag(words []string) (bag []float64) {
	for _, word := range words {
		// Append 1 if the patternWords contains the actual word, else 0
		var valueToAppend float64 = 0
		if util.Contains(sentence.Tokenize(), word) {
			valueToAppend = 1
		}

		bag = append(bag, valueToAppend)
	}

	return bag
}
