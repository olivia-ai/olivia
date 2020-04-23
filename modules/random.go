package modules

import (
	"fmt"
	"math/rand"
	"strconv"
)

var RandomTag = "random number"

// RandomNumberReplacer replaces the pattern contained inside the response by a random number.
// See modules/modules.go#Module.Replacer() for more details.
func RandomNumberReplacer(_, _, response, _ string) (string, string) {
	return RandomTag, fmt.Sprintf(response, strconv.Itoa(rand.Intn(100)))
}
