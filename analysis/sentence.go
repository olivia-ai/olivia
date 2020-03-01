package analysis

import (
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/caneroj1/stemmer"
	"github.com/gookit/color"
	"github.com/neurosnap/sentences"
	"github.com/olivia-ai/olivia/modules"
	"github.com/olivia-ai/olivia/network"
	"github.com/olivia-ai/olivia/util"
	gocache "github.com/patrickmn/go-cache"
)

type Sentence struct {
	Content string
}

type Result struct {
	Tag   string
	Value float64
}

var userCache = gocache.New(5*time.Minute, 5*time.Minute)

const DontUnderstand = "don't understand"

// NewSentence returns a Sentence object where the content has been arranged
func NewSentence(content string) Sentence {
	return Sentence{Arrange(content)}
}

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

// Classify the sentence with the model
func (sentence Sentence) PredictTag(neuralNetwork network.Network) string {
	words, classes, _ := Organize()

	// Predict with the model
	predict := neuralNetwork.Predict(sentence.WordsBag(words))

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

	return resultsTag[0].Tag
}

// RandomizeResponse takes the entry message, the response tag and the token and returns a random
// message from res/intents.json where the triggers are applied
func RandomizeResponse(entry string, tag string, token string) (string, string) {
	if tag == DontUnderstand {
		return DontUnderstand, util.GetMessage(tag)
	}

	// Append the modules intents to the intents from res/intents.json
	intents := append(SerializeIntents(), SerializeModulesIntents()...)

	for _, intent := range intents {
		if intent.Tag != tag {
			continue
		}

		// Reply a "don't understand" message if the context isn't correct
		cacheTag, _ := userCache.Get(token)
		if intent.Context != "" && cacheTag != intent.Context {
			return DontUnderstand, util.GetMessage(DontUnderstand)
		}

		// Set the actual context
		userCache.Set(token, tag, gocache.DefaultExpiration)

		// Choose a random response in intents
		response := intent.Responses[0]
		if len(intent.Responses) > 1 {
			response = intent.Responses[rand.Intn(len(intent.Responses))]
		}

		// And then apply the triggers on the message
		return modules.ReplaceContent(tag, entry, response, token)
	}

	return DontUnderstand, util.GetMessage(DontUnderstand)
}

// Calculate send the sentence content to the neural network and returns a response with the matching tag
func (sentence Sentence) Calculate(cache gocache.Cache, neuralNetwork network.Network, token string) (string, string) {
	tag, found := cache.Get(sentence.Content)

	// Predict tag with the neural network if the sentence isn't in the cache
	if !found {
		tag = sentence.PredictTag(neuralNetwork)
		cache.Set(sentence.Content, tag, gocache.DefaultExpiration)
	}

	return RandomizeResponse(sentence.Content, tag.(string), token)
}

// LogResults print in the console the sentence and its tags sorted by prediction
func LogResults(entry string, results []Result) {
	green := color.FgGreen.Render
	yellow := color.FgYellow.Render

	color.FgCyan.Printf("\n\"%s\"\n", entry)
	for _, result := range results {
		// Arbitrary choice of 0.005 to have less tags to show
		if result.Value < 0.005 {
			continue
		}

		fmt.Printf("  %s %s - %s\n", green("▫︎"), result.Tag, yellow(result.Value))
	}
}
