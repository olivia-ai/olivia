package language

import "regexp"

// SearchToken searchs a token in the given sentence and returns it.
func SearchToken(sentence string) string {
	// Search the token with a regex
	tokenRegex := regexp.MustCompile(`[a-z0-9]{32}`)
	// Returns the found token
	return tokenRegex.FindString(sentence)
}
