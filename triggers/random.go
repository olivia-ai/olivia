package triggers

import (
	"math/rand"
	"strconv"
	"strings"
)

type Random struct {}

// Replace the key by a random number
func (random Random) ReplaceContent() string {
	return strings.Replace(
		Response,
		"${RANDOM_NUMBER}",
		strconv.Itoa(rand.Intn(100)),
		1)
}
