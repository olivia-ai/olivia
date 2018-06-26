package triggers

import (
	"math/rand"
	"strconv"
	"strings"
)

type Random struct {
	Response string
}

// Replace the key by a random number
func (random Random) ReplaceContent() string {
	return strings.Replace(
		random.Response,
		"${RANDOM_NUMBER}",
		strconv.Itoa(rand.Intn(100)),
		1)
}
