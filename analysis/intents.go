package analysis

import (
	"encoding/json"
	"sort"

	"github.com/olivia-ai/olivia/modules"
	"github.com/olivia-ai/olivia/util"
)

const intentsFile = "res/datasets/intents.json"

var intents []Intent

// Intent is a way to group sentences that mean the same thing and link them with a tag which
// represents what they mean, some responses that the bot can reply and a context
type Intent struct {
	Tag       string   `json:"tag"`
	Patterns  []string `json:"patterns"`
	Responses []string `json:"responses"`
	Context   string   `json:"context"`
}

// Document is any sentence from the intents' patterns linked with its tag
type Document struct {
	Sentence Sentence
	Tag      string
}

// CacheIntents set the given intents to the global variable intents
func CacheIntents(_intents []Intent) {
	intents = _intents
}

// GetIntents returns the cached intents
func GetIntents() []Intent {
	return intents
}

// SerializeIntents returns a list of intents retrieved from `res/datasets/intents.json`
func SerializeIntents() []Intent {
	err := json.Unmarshal(util.ReadFile(intentsFile), &intents)
	if err != nil {
		panic(err)
	}

	return intents
}

// SerializeModulesIntents retrieves all the registered modules and returns an array of Intents
func SerializeModulesIntents() []Intent {
	registeredModules := modules.GetModules()
	intents := make([]Intent, len(registeredModules))

	for _, module := range registeredModules {
		intents = append(intents, Intent{
			Tag:       module.Tag,
			Patterns:  module.Patterns,
			Responses: module.Responses,
			Context:   "",
		})
	}

	return intents
}

// Organize intents with an array of all words, an array with a representative word of each tag
// and an array of Documents which contains a word list associated with a tag
func Organize() (words, classes []string, documents []Document) {
	// Append the modules intents to the intents from res/datasets/intents.json
	intents := append(SerializeIntents(), SerializeModulesIntents()...)

	for _, intent := range intents {
		for _, pattern := range intent.Patterns {
			// Tokenize the pattern's sentence
			patternSentence := Sentence{"en", pattern}
			patternSentence.Arrange()

			// Add each word to response
			for _, word := range patternSentence.Stem() {

				if !util.Contains(words, word) {
					words = append(words, word)
				}
			}

			// Add a new document
			documents = append(documents, Document{
				patternSentence,
				intent.Tag,
			})
		}

		// Add the intent tag to classes
		classes = append(classes, intent.Tag)
	}

	sort.Strings(words)
	sort.Strings(classes)

	return words, classes, documents
}
