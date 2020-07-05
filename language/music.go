package language

import (
	"strings"
)

// SpotifyKeyword is the map for having the music keywords in different languages
var SpotifyKeyword = map[string]SpotifyKeywords{
	"en": {
		Play: "play",
		From: "from",
		On:   "on",
	},
	"de": {
		Play: "spiele",
		From: "von",
		On:   "auf",
	},
	"fr": {
		Play: "joue",
		From: "de",
		On:   "sur",
	},
	"es": {
		Play: "Juega",
		From: "de",
		On:   "en",
	},
	"ca": {
		Play: "Juga",
		From: "de",
		On:   "a",
	},
	"it": {
		Play: "suona",
		From: "da",
		On:   "a",
	},
	"tr": {
		Play: "Başlat",
		From: "dan",
		On:   "kadar",
	},
	"nl": {
		Play: "speel",
		From: "van",
		On:   "op",
	},
	"el": {
		Play: "αναπαραγωγή",
		From: "από",
		On:   "στο",
	},
}

// SpotifyKeywords are the keywords used to get music name
type SpotifyKeywords struct {
	Play string
	From string
	On   string
}

// SearchMusic returns a music title and artist found from the given sentence
func SearchMusic(locale, sentence string) (music, artist string) {
	words := strings.Split(sentence, " ")

	// Iterate through the words of the sentence
	playAppeared, fromAppeared, onAppeared := false, false, false
	for _, word := range words {
		// If "on" appeared
		if word == SpotifyKeyword[locale].On {
			onAppeared = true
		}

		// Add the current word if its between from and on
		if fromAppeared && !onAppeared {
			artist += word + " "
		}

		// If "from" appeared
		if LevenshteinDistance(word, SpotifyKeyword[locale].From) < 2 {
			fromAppeared = true
		}

		// Add the current word if its between play and from
		if playAppeared && !fromAppeared && !onAppeared {
			music += word + " "
		}

		// If "play" appeared
		if LevenshteinDistance(word, SpotifyKeyword[locale].Play) < 2 {
			playAppeared = true
		}
	}

	// Trim the spaces and return
	return strings.TrimSpace(music), strings.TrimSpace(artist)
}
