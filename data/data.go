package data

import (
	"github.com/olivia-ai/olivia/util"
)

// Conversation describes an item of a conversation dataset containing a question
// and an answer.
type Conversation struct {
	Question string
	Answer string
}

// ReadCSVConversationalDataset reads the CSV file from the given path and returns
// a slice of Conversation structs.
func ReadCSVConversationalDataset(path string) (conversations []Conversation) {
	for _, lines := range util.ReadCSV(path) {
		conversations = append(conversations, Conversation{
			Question: lines[0],
			Answer: lines[1],
		})
	}
	
	return
}