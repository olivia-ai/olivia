package triggers

import (
	"time"
	"math/rand"
	"strings"
	"strconv"
)

type Random struct {
	Response string
}

// Replace the key by a random number
func (random Random) ReplaceContent() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return strings.Replace(
		random.Response,
		"${RANDOM_NUMBER}",
		strconv.Itoa(r.Intn(100)),
		1)
}
