package analysis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"../slice"
)

type Intent struct {
	Tag       string   `json:"tag"`
	Patterns  []string `json:"patterns"`
	Responses []string `json:"responses"`
	Context   string   `json:"context"`
}

type Document struct {
	Sentence Sentence
	Tag   string
}

// Return the intents json file's content
func Read() []byte {
	bytes, err := ioutil.ReadFile("intents.json")
	if err != nil {
		fmt.Println(err)
	}

	return bytes
}

// Unmarshal the json and return the array of Intents
func Serialize() []Intent {
	var intents []Intent
	json.Unmarshal(Read(), &intents)

	return intents
}

// Organize intents with an array of all words, an array with a representative word of each tag
// and an array of Documents which contains a word list associated with a tag
func Organize() (words, classes []string, documents []Document) {
	for _, intent := range Serialize() {
		for _, pattern := range intent.Patterns {
			// Tokenize the pattern's sentence
			patternSentence := Sentence{pattern}

			// Add each word to response
			for _, word := range patternSentence.Tokenize() {
				words = append(words, word)
			}

			// Add a new document
			documents = append(documents, Document{
				patternSentence,
				intent.Tag,
			})

			// Add the intent tag to class if it doesn't exists
			if !slice.Contains(classes, intent.Tag) {
				classes = append(classes, intent.Tag)
			}
		}
	}

	return words, classes, documents
}
