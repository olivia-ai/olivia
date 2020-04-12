package language

import "regexp"

// SearchTokens searches 2 tokens in the given sentence and returns it.
func SearchTokens(sentence string) []string {
	// Search the token with a regex
	tokenRegex := regexp.MustCompile(`[a-z0-9]{32}`)
	// Returns the found token
	return tokenRegex.FindAllString(sentence, 2)
}
