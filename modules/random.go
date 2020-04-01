package modules

import (
	"fmt"
	"math/rand"
	"strconv"
)

var randomTag = "random number"

func init() {
	RegisterModule(Module{
		Tag: randomTag,
		Patterns: []string{
			"Give me a random number",
			"Generate a random number",
		},
		Responses: []string{
			"The number is %s",
		},
		Replacer: RandomNumberReplacer,
	})
}

// RandomNumberReplacer replaces the pattern contained inside the response by a random number.
// See modules/modules.go#Module.Replacer() for more details.
func RandomNumberReplacer(_, response, _ string) (string, string) {
	return randomTag, fmt.Sprintf(response, strconv.Itoa(rand.Intn(100)))
}
