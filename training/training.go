package training

import (
	"../analysis"
	"../slice"
	"github.com/xamber/Varis"
)

// Return the data for training the model
func TrainData() (training [][2]varis.Vector) {
	words, classes, documents := analysis.Organize()

	for _, document := range documents {
		outputRow := make([]float64, len(classes))
		bag := document.Sentence.WordsBag(words)

		// Change value to 1 where there is the document Tag
		outputRow[slice.Index(classes, document.Tag)] = 1

		// Append data to trainx and trainy
		training = append(training, [2]varis.Vector{bag, outputRow})
	}

	return training
}

// Create the neural network
func CreateModel() varis.Perceptron {
	_, classes, _ := analysis.Organize()
	training := TrainData()
	inputLayers := len(training[0][0])

	// Create neural network with words bags
	network := varis.CreatePerceptron(inputLayers, inputLayers/2, len(classes))
	trainer := varis.PerceptronTrainer{
		Network: &network,
		Dataset: varis.Dataset(training),
	}
	trainer.BackPropagation(1000)

	return network
}
