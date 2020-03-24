package language

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
