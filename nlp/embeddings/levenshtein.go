package embeddings

// LevenshteinDistance calculates the Levenshtein Distance between two given words and returns it.
// Please see https://en.wikipedia.org/wiki/Levenshtein_distance.
func LevenshteinDistance(first, second string) int {
	// Returns the length if it's empty
	if first == "" {
		return len(second)
	}
	if second == "" {
		return len(first)
	}

	if first[0] == second[0] {
		return LevenshteinDistance(first[1:], second[1:])
	}

	a := LevenshteinDistance(first[1:], second[1:])
	if b := LevenshteinDistance(first, second[1:]); a > b {
		a = b
	}

	if c := LevenshteinDistance(first[1:], second); a > c {
		a = c
	}

	return a + 1
}

// LevenshteinContains checks if the given word is present in the given token base using levenshtein distance
// using the given rate.
// It returns the boolean verifiying the condition and the original word found if any.
func LevenshteinContains(tokenBase []string, word string, rate int) (bool, string) {
	for _, token := range tokenBase {
		if LevenshteinDistance(token, word) <= rate {
			return true, token
		}
	}
	return false, ""
}

// GetLevenshteinEmbedding finds the closest word in the given token base using levenshtein distance
// and returns it as an embedding (bag of words).
func GetLevenshteinEmbedding(vocabulary []string, word string) []float64 {
	// Calculate levenshtein distances between every word
	min, minWord := 100, ""
	for _, token := range vocabulary {
		if distance := LevenshteinDistance(token, word); distance < min {
			min = distance
			minWord = word
		}

		if min == 0 {
			break
		}
	}

	return getEmbedding(vocabulary, minWord)
}