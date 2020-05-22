package util

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// Message contains the message's tag and its contained matched sentences
type Message struct {
	Tag      string   `json:"tag"`
	Messages []string `json:"messages"`
}

var messages = map[string][]Message{}

// SerializeMessages serializes the content of `res/datasets/messages.json` in JSON
func SerializeMessages(locale string) []Message {
	var currentMessages []Message
	err := json.Unmarshal(ReadFile("res/locales/"+locale+"/messages.json"), &currentMessages)
	if err != nil {
		fmt.Println(err)
	}

	messages[locale] = currentMessages

	return currentMessages
}

// GetMessages returns the cached messages for the given locale
func GetMessages(locale string) []Message {
	return messages[locale]
}

// GetMessageByTag returns a message found by the given tag and locale
func GetMessageByTag(tag, locale string) Message {
	for _, message := range messages[locale] {
		if tag != message.Tag {
			continue
		}

		return message
	}

	return Message{}
}

// GetMessage retrieves a message tag and returns a random message chose from res/datasets/messages.json
func GetMessage(locale, tag string) string {
	for _, message := range messages[locale] {
		// Find the message with the right tag
		if message.Tag != tag {
			continue
		}

		// Returns the only element if there aren't more
		if len(message.Messages) == 1 {
			return message.Messages[0]
		}

		// Returns a random sentence
		rand.Seed(time.Now().UnixNano())
		return message.Messages[rand.Intn(len(message.Messages))]
	}

	return ""
}
