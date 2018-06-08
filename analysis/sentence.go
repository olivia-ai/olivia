package analysis

import (
	"../slice"
	"github.com/neurosnap/sentences"
	"github.com/stevenmiller888/go-mind"
	"math/rand"
	"sort"
	"strings"
)

// Initialize the user's context cache
var cache = make(map[int]string)

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

	// Enumerate the results with the intent tags
	var resultsTag []Result
	for i, result := range results {
		resultsTag = append(resultsTag, Result{classes[i], result})
	}

	// Sort the results in ascending order
	sort.Slice(resultsTag, func(i, j int) bool {
		return resultsTag[i].Value > resultsTag[j].Value
	})

	return resultsTag[0]
}

// Returns the human readable response
func (sentence Sentence) Response(model *mind.Mind, userId int) string {
	result := sentence.Classify(model)

	// Iterate all the json intents
	for _, intent := range Serialize() {
		if intent.Tag != result.Tag {
			continue
		}

		if intent.Context != "" && cache[userId] != intent.Context {
			return "Je ne comprends pas :("
		}

		cache[userId] = intent.Tag

		// Return a random response
		if len(intent.Responses) != 1 {
			return intent.Responses[rand.Intn(len(intent.Responses)-1)]
		} else {
			return intent.Responses[0]
		}
	}

	// Error
	return "Je ne comprends pas :("
}
