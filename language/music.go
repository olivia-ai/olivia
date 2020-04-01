package language

import (
	"strings"
)

// SearchMusic returns a music title and artist found from the given sentence
func SearchMusic(sentence string) (music, artist string) {
	words := strings.Split(sentence, " ")

	// Iterate through the words of the sentence
	playAppeared, fromAppeared, onAppeared := false, false, false
	for _, word := range words {
		// If "on" appeared
		if word == "on" {
			onAppeared = true
		}

		// Add the current word if its between from and on
		if fromAppeared && !onAppeared {
			artist += word + " "
		}

		// If "from" appeared
		if LevenshteinDistance(word, "from") < 2 {
			fromAppeared = true
		}

		// Add the current word if its between play and from
		if playAppeared && !fromAppeared {
			music += word + " "
		}

		// If "play" appeared
		if LevenshteinDistance(word, "play") < 2 {
			playAppeared = true
		}
	}

	// Trim the spaces and return
	return strings.TrimSpace(music), strings.TrimSpace(artist)
}
