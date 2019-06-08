package analysis

import (
	"fmt"
	"github.com/caneroj1/stemmer"
	"github.com/gookit/color"
	"github.com/neurosnap/sentences"
	"github.com/olivia-ai/gonn/gonn"
	"github.com/olivia-ai/olivia/triggers"
	"github.com/olivia-ai/olivia/util"
	gocache "github.com/patrickmn/go-cache"
	"math/rand"
	"sort"
	"strings"
	"time"
)

type Sentence struct {
	Content string
}

// NewSentence creates a Sentence and arrange the content by removing ignored chars
// and trim spaces.
func NewSentence(content string) (sentence Sentence) {
	sentence = Sentence{content}
	sentence.Trim()

	return sentence
}

type Result struct {
	Tag   string
	Value float64
}

var userCache = gocache.New(5*time.Minute, 5*time.Minute)

// Trim removes ignored chars, trim spaces and change the sentence to lower case.
func (sentence Sentence) Trim() {
	var text string
	ignoredChars := []string{"?", "-", ".", "!"}
	for _, ignoredChar := range ignoredChars {
		text = strings.Replace(text, ignoredChar, " ", -1)
	}

	text = strings.TrimSpace(text)
	sentence.Content = strings.ToLower(text)
}

// Tokenize returns an array of the sentence's words stemmed and in lower case.
func (sentence Sentence) Tokenize() (tokenizedWords []string) {
	tokenizer := sentences.NewWordTokenizer(sentences.NewPunctStrings())
	tokens := tokenizer.Tokenize(sentence.Content, false)

	// Get the string token and add it to tokenizedWords
	for _, tokenizedWord := range tokens {
		word := strings.ToLower(tokenizedWord.Tok)
		word = stemmer.Stem(word)

		tokenizedWords = append(tokenizedWords, word)
	}

	return tokenizedWords
}

// WordsBag retrieves all the intents words and returns the bag of words of the Sentence content
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

// PredictTag classify the sentence with the model and returns the matching tag of res/intents.json.
func (sentence Sentence) PredictTag(n gonn.NeuralNetwork) string {
	words, classes, _ := Organize()

	// Predict with the model
	predict := n.Forward(sentence.WordsBag(words))

	// Enumerate the results with the intent tags
	var resultsTag []Result
	for i, result := range predict {
		resultsTag = append(resultsTag, Result{classes[i], result})
	}

	// Sort the results in ascending order
	sort.Slice(resultsTag, func(i, j int) bool {
		return resultsTag[i].Value > resultsTag[j].Value
	})

	LogResults(sentence.Content, resultsTag)

	if resultsTag[0].Value < 0.35 {
		return "don't understand"
	}

	return resultsTag[0].Tag
}

// Returns the human readable response
func RandomizeResponse(entry string, tag string, userId string) string {
	if tag == "don't understand" {
		return util.GetMessage(tag)
	}

	// Iterate all the json intents
	for _, intent := range SerializeIntents() {
		if intent.Tag != tag {
			continue
		}

		cacheTag, _ := userCache.Get(userId)
		if intent.Context != "" && cacheTag != intent.Context {
			return util.GetMessage("don't understand")
		}

		userCache.Set(userId, tag, gocache.DefaultExpiration)

		response := intent.Responses[0]
		// Return a random response if there are more than one
		if len(intent.Responses) > 1 {
			response = intent.Responses[rand.Intn(len(intent.Responses))]
		}

		return triggers.ReplaceContent(entry, response)
	}

	// Error
	return util.GetMessage("don't understand")
}

// Respond with the cache or the model
func (sentence Sentence) Calculate(cache gocache.Cache, network gonn.NeuralNetwork, userId string) (string, string) {
	tag, found := cache.Get(sentence.Content)

	// If the sentence isn't in the redis database
	if !found {
		tag = sentence.PredictTag(network)
		cache.Set(sentence.Content, tag, gocache.DefaultExpiration)
	}

	return RandomizeResponse(sentence.Content, tag.(string), userId), tag.(string)
}

func LogResults(entry string, results []Result) {
	green := color.FgGreen.Render
	yellow := color.FgYellow.Render

	color.FgCyan.Printf("\n\"%s\"\n", entry)
	for _, result := range results {
		if result.Value < 0.05 {
			continue
		}

		fmt.Printf("  %s %s - %s\n", green("▫︎"), result.Tag, yellow(result.Value))
	}
}
