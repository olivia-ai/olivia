package modules

import (
	"fmt"
	"github.com/olivia-ai/olivia/language"
	"github.com/olivia-ai/olivia/util"
)

var areaTag = "area"

func init() {
	RegisterModule(Module{
		Tag: areaTag,
		Patterns: []string{
			"What is the area of ",
			"Give me the area of ",
		},
		Responses: []string{
			"The area of %s is %gkm²",
		},
		Replacer: AreaReplacer,
	})
}

func AreaReplacer(entry, response string) (string, string) {
	country := language.FindCountry(entry)

	// If there isn't a country respond with a message from res/messages.json
	if country.Code == "" {
		responseTag := "no country"
		return responseTag, util.GetMessage(responseTag)
	}

	return areaTag, fmt.Sprintf(response, country.CommonName, country.Area)
}
