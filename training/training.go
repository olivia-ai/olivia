package training

import (
	"github.com/olivia-ai/olivia/analysis"
	"github.com/olivia-ai/olivia/network"
	"github.com/olivia-ai/olivia/util"
)

// TrainData returns the inputs and targets for the neural network
func TrainData() (inputs, targets [][]float64) {
	words, classes, documents := analysis.Organize()

	for _, document := range documents {
		outputRow := make([]float64, len(classes))
		bag := document.Sentence.WordsBag(words)

		// Change value to 1 where there is the document Tag
		outputRow[util.Index(classes, document.Tag)] = 1

		// Append data to trainx and trainy
		inputs = append(inputs, bag)
		targets = append(targets, outputRow)
	}

	return inputs, targets
}

// CreateNeuralNetwork returns a new neural network which is loaded from res/training.json or
// trained from TrainData() inputs and targets.
func CreateNeuralNetwork() (neuralNetwork network.Network) {
	// Train the model if there is no training file
	trainx, trainy := TrainData()

	neuralNetwork = network.CreateNetwork(0.1, trainx, trainy, 50)
	neuralNetwork.Train(1000)

	return
}
