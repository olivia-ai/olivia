package training

import (
	"../analysis"
	"../slice"
	"strings"
)

func TrainData() {
	words, classes, documents := analysis.Organize()

	for _, document := range documents {
		var bag []int
		var patternWords []string
		outputRow := make([]int, len(classes))

		// Down case all document's words
		for _, word := range document.Words {
			patternWords = append(patternWords, strings.ToLower(word))
		}

		// Iterate all intents.json words
		for _, word := range words {
			// Append 1 if the patternWords contains the actual word, else 0
			valueToAppend := 0
			if slice.SliceContains(patternWords, word) {
				valueToAppend = 1
			}

			bag = append(bag, valueToAppend)
		}
	}
}