package modules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/olivia-ai/olivia/util"
)

const JOKEURL = "https://official-joke-api.appspot.com/random_joke"

// JokesTag is the intent tag for its module
var JokesTag = "jokes"

type Joke struct {
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

// CurrencyReplacer replaces the pattern contained inside the response by the currency of the country
// specified in the message.
// See modules/modules.go#Module.Replacer() for more details.
func JokesReplacer(locale, entry, response, _ string) (string, string) {

	resp, err := http.Get(JOKEURL)
	if err != nil {
		responseTag := "no jokes"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		responseTag := "no jokes"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	joke := &Joke{}

	err = json.Unmarshal(body, joke)
	if err != nil {
		responseTag := "no jokes"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	jokeStr := joke.Setup + " " + joke.Punchline

	return JokesTag, fmt.Sprintf(response, jokeStr)
}
