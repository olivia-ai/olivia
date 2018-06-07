package analysis

import (
	"../slice"
	"github.com/neurosnap/sentences"
	"github.com/stevenmiller888/go-mind"
	"sort"
	"strings"
	"math/rand"
)

type Sentence struct {
	Content string
}

type Result struct {
	Tag   string
	Value float64
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

// Classify the sentence with the model
func (sentence Sentence) Classify(model *mind.Mind) Result {
	words, classes, _ := Organize()

	// Predict with the model
	predict := model.Predict([][]float64{
		sentence.WordsBag(words),
	})

	// Put the predict results in an array
	var results []float64
	_, size := predict.Dims()
	for i := 0; i < size; i++ {
		results = append(results, predict.At(0, i))
	}

	// Sort the results in ascending order
	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})

	// Return the good value
	for i, result := range results {
		if result == predict.At(0, 1) {
			return Result{classes[i], result}
		}
	}

	return Result{classes[0], 0}
}

// Returns the human readable response
func (sentence Sentence) Response(model *mind.Mind) string {
	result := sentence.Classify(model)

	// Iterate all the json intents
	for _, intent := range Serialize() {
		if intent.Tag != result.Tag {
			continue
		}

		// Return a random response
		return intent.Responses[rand.Intn(len(intent.Responses) - 1)]
	}

	// Error
	return "Je ne comprends pas :("
}