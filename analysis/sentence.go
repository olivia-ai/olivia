package analysis

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/gookit/color"
	"github.com/olivia-ai/olivia/modules"
	"github.com/olivia-ai/olivia/network"
	"github.com/olivia-ai/olivia/util"
	gocache "github.com/patrickmn/go-cache"
)

type Sentence struct {
	Content string
}

// Result contains a predicted value with its tag and its value
type Result struct {
	Tag   string
	Value float64
}

var userCache = gocache.New(5*time.Minute, 5*time.Minute)

// DontUnderstand contains the tag for the don't understand messages
const DontUnderstand = "don't understand"

// NewSentence returns a Sentence object where the content has been arranged
func NewSentence(content string) (sentence Sentence) {
	sentence = Sentence{content}
	sentence.Arrange()

	return
}

// PredictTag classifies the sentence with the model
func (sentence Sentence) PredictTag(neuralNetwork network.Network, intentsPath string) string {
	words, classes, _ := Organize(intentsPath)

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
func RandomizeResponse(intentsPath, entry, tag, token string) (string, string) {
	if tag == DontUnderstand {
		return DontUnderstand, util.GetMessage(tag)
	}

	// Append the modules intents to the intents from res/intents.json
	intents := append(SerializeIntents(intentsPath), SerializeModulesIntents()...)

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
func (sentence Sentence) Calculate(cache gocache.Cache, neuralNetwork network.Network, intentsPath, token string) (string, string) {
	tag, found := cache.Get(sentence.Content)

	// Predict tag with the neural network if the sentence isn't in the cache
	if !found {
		tag = sentence.PredictTag(neuralNetwork, intentsPath)
		cache.Set(sentence.Content, tag, gocache.DefaultExpiration)
	}

	return RandomizeResponse(intentsPath, sentence.Content, tag.(string), token)
}

// LogResults print in the console the sentence and its tags sorted by prediction
func LogResults(entry string, results []Result) {
	green := color.FgGreen.Render
	yellow := color.FgYellow.Render

	color.FgCyan.Printf("\n\"%s\"\n", entry)
	for _, result := range results {
		// Arbitrary choice of 0.005 to have less tags to show
		if result.Value < 0.001 {
			continue
		}

		fmt.Printf("  %s %s - %s\n", green("▫︎"), result.Tag, yellow(result.Value))
	}
}
