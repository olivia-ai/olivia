package modules

import (
	"fmt"
	"github.com/olivia-ai/olivia/language"
	"github.com/olivia-ai/olivia/util"
)

func init() {
	RegisterModule(Module{
		Tag: "area",
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

func AreaReplacer(entry, response string) string {
	country := language.FindCountry(entry)

	// If there isn't a country respond with a message from res/messages.json
	if country.Code == "" {
		return util.GetMessage("no country")
	}

	return fmt.Sprintf(response, country.CommonName, country.Area)
}
