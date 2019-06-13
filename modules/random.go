package modules

import (
	"fmt"
	"math/rand"
	"strconv"
)

func init() {
	RegisterModule(Module{
		Tag: "random_number",
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

func RandomNumberReplacer(_, response string) string {
	return fmt.Sprintf(response, strconv.Itoa(rand.Intn(100)))
}
