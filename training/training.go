package training

import (
	"../analysis"
	"../slice"
	"github.com/stevenmiller888/go-mind"
)

// Return the data for training the model
func TrainData() (training [][][]float64) {
	words, classes, documents := analysis.Organize()

	for _, document := range documents {
		outputRow := make([]float64, len(classes))
		bag := document.Sentence.WordsBag(words)

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
