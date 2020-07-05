package modules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/olivia-ai/olivia/util"
)

const adviceURL = "https://api.adviceslip.com/advice"

// AdvicesTag is the intent tag for its module
var AdvicesTag = "advices"

// AdvicesReplacer replaces the pattern contained inside the response by a random advice from the api
// specified by the adviceURL.
// See modules/modules.go#Module.Replacer() for more details.
func AdvicesReplacer(locale, entry, response, _ string) (string, string) {

	resp, err := http.Get(adviceURL)
	if err != nil {
		responseTag := "no advices"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		responseTag := "no advices"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)
	advice := result["slip"].(map[string]interface{})["advice"]

	return AdvicesTag, fmt.Sprintf(response, advice)
}
