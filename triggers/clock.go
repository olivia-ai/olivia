package triggers

import (
	"time"
	"fmt"
	"strings"
)

type Clock struct {
	Sentence string
}

// Replace the content of the sentence by the actual clock
func (clock Clock) ReplaceContent() string {
	hours, minutes, _ := time.Now().Clock()
	actualClock := fmt.Sprintf("%dh%d", hours, minutes)

	return strings.Replace(clock.Sentence, "${CLOCK}", actualClock, 1)
}