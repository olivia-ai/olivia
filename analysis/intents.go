package analysis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Intent struct {
	Tag       string   `json:"tag"`
	Patterns  []string `json:"patterns"`
	Responses []string `json:"responses"`
	Context   string   `json:"context"`
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
