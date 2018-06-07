package training

import (
	"../analysis"
	"../slice"
	"strings"
	"fmt"
)

func TrainData() (trainx, trainy [][]int) {
	words, classes, documents := analysis.Organize()

	for _, document := range documents {
		var bag []int
		var patternWords []string
		outputRow := make([]int, len(classes))

		// Down case all document's words
		for _, word := range document.Words {
			patternWords = append(patternWords, strings.ToLower(word))
		}

		fmt.Println(patternWords)

		// Iterate all intents.json words
		for _, word := range words {
			// Append 1 if the patternWords contains the actual word, else 0
			valueToAppend := 0
			if slice.Contains(patternWords, word) {
				valueToAppend = 1
			}

			bag = append(bag, valueToAppend)
		}

		// Change value to 1 where there is the document Tag
		outputRow[slice.Index(classes, document.Tag)] = 1

		// Append data to trainx and trainy
		trainx = append(trainx, bag)
		trainy = append(trainy, outputRow)
	}

	return trainx, trainy
}