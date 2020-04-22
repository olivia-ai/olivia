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

// RemoveStopWords takes an arary of words, removes the stopwords and returns it
func RemoveStopWords(words []string) []string {
	// Don't remove stopwords for small sentences like “How are you” because it will remove all the words
	if len(words) <= 4 {
		return words
	}

	// Read the content of the stopwords file
	stopWords := string(util.ReadFile("res/datasets/stopwords.txt"))

	// Iterate through all the stopwords
	for _, stopWord := range strings.Split(stopWords, "\n") {
		// Iterate through all the words of the given array
		for i, word := range words {
			// Continue if the word isn't a stopword
			if !strings.Contains(stopWord, word) {
				continue
			}

			// Remove the word from the array
			if len(words) == 1 {
				words = []string{}
			} else {
				words[i] = words[len(words)-1]
				words = words[:len(words)-1]
			}
		}
	}

	return words
}

// Tokenize returns a list of words that have been lower-cased
func (sentence Sentence) Tokenize() (tokens []string) {
	// Split the sentence in words
	tokens = strings.Fields(sentence.Content)

	// Lower case each word
	for i, token := range tokens {
		tokens[i] = strings.ToLower(token)
	}

	tokens = RemoveStopWords(tokens)

	return
}

// Stem returns the sentence split in stemmed words
func (sentence Sentence) Stem() (tokenizeWords []string) {
	tokens := sentence.Tokenize()

	// Get the string token and push it to tokenizeWord
	for _, tokenizeWord := range tokens {
		word := stemmer.Stem(tokenizeWord)
		tokenizeWords = append(tokenizeWords, word)
	}

	return
}

// WordsBag retrieves the intents words and returns the sentence converted in a bag of words
func (sentence Sentence) WordsBag(words []string) (bag []float64) {
	for _, word := range words {
		// Append 1 if the patternWords contains the actual word, else 0
		var valueToAppend float64
		if util.Contains(sentence.Stem(), word) {
			valueToAppend = 1
		}

		bag = append(bag, valueToAppend)
	}

	return bag
}
