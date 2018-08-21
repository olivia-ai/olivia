package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

type Message struct {
	Tag      string   `json:"tag"`
	Messages []string `json:"messages"`
}

var messages = SerializeMessages()

func SerializeMessages() (messages []Message) {
	bytes, err := ioutil.ReadFile("res/messages.json")
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(bytes, &messages)

	return messages
}

func GetMessage(tag string) string {
	for _, message := range messages {
		if message.Tag != tag {
			continue
		}

		if len(message.Messages) == 1 {
			return message.Messages[0]
		}

		return message.Messages[rand.Intn(len(message.Messages)-1)]
	}

	return messages[0].Messages[0]
}
