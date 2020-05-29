package modules

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/olivia-ai/olivia/language"
	"github.com/olivia-ai/olivia/util"
)

// RandomTag is the intent tag for its module
var RandomTag = "random number"

// RandomNumberReplacer replaces the pattern contained inside the response by a random number.
// See modules/modules.go#Module.Replacer() for more details.
func RandomNumberReplacer(locale, entry, response, _ string) (string, string) {
	limitArr, err := language.FindRangeLimits(locale, entry)
	if err != nil {
		if limitArr != nil {
			return RandomTag, fmt.Sprintf(response, strconv.Itoa(rand.Intn(100)))
		}

		responseTag := "no random range"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	min := limitArr[0]
	max := limitArr[1]
	randNum := rand.Intn((max - min)) + min
	return RandomTag, fmt.Sprintf(response, strconv.Itoa(randNum))
}
