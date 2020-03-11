package analysis

import (
	"regexp"
	"strings"

	"github.com/caneroj1/stemmer"
	"github.com/olivia-ai/olivia/util"
)

// Arrange checks the format of a string to normalize it, remove ignored characters
func (sentence *Sentence) Arrange() {
	// Remove punctuation after letters
	punctuationRegex := regexp.MustCompile(`[a-zA-Z]( )?(\.|\?|!)`)
	sentence.Content = punctuationRegex.ReplaceAllStringFunc(sentence.Content, func(s string) string {
		punctuation := regexp.MustCompile(`(\.|\?|!)`)
		return punctuation.ReplaceAllString(s, "")
	})

	sentence.Content = strings.TrimSpace(sentence.Content)
}

// Tokenize returns a list of words that have been lower-cased
func (sentence Sentence) Tokenize() []string {
	// Split the sentence in words
	tokens := strings.Fields(sentence.Content)

	// Lower case each word
	for i, token := range tokens {
		tokens[i] = strings.ToLower(token)
	}

	return tokens
}

// Stem returns the sentence split in stemmed words
func (sentence Sentence) Stem() (tokenizedWords []string) {
	tokens := sentence.Tokenize()

	// Get the string token and push it to tokenizedWords
	for _, tokenizedWord := range tokens {
		word := stemmer.Stem(tokenizedWord)
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
