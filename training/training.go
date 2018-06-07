package training

import (
	"../analysis"
	"../slice"
	"strings"
	"github.com/stevenmiller888/go-mind"
)

// Return the data for training the model
func TrainData() (training [][][]float64) {
	words, classes, documents := analysis.Organize()

	for _, document := range documents {
		var bag []float64
		var patternWords []string
		outputRow := make([]float64, len(classes))

		// Down case all document's words
		for _, word := range document.Words {
			patternWords = append(patternWords, strings.ToLower(word))
		}

		// Iterate all intents.json words
		for _, word := range words {
			// Append 1 if the patternWords contains the actual word, else 0
			var valueToAppend float64 = 0
			if slice.Contains(patternWords, word) {
				valueToAppend = 1
			}

			bag = append(bag, valueToAppend)
		}

		// Change value to 1 where there is the document Tag
		outputRow[slice.Index(classes, document.Tag)] = 1

		// Append data to trainx and trainy
		training = append(training, [][]float64{bag, outputRow})
	}

	return training
}

// Create the go-mind model
func CreateModel() *mind.Mind {
	training := TrainData()

	// Initialize the model
	model := mind.New(0.7, 1000, len(training[0][0]), "sigmoid")
	model.Learn(training)

	return model
}