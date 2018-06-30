package training

import (
	"github.com/fxsjy/gonn/gonn"
	"github.com/olivia-ai/Api/analysis"
	"github.com/olivia-ai/Api/slice"
)

// Return the inputs and targets generated from the intents for the neural network
func TrainData() (inputs, targets [][]float64) {
	words, classes, documents := analysis.Organize()

	for _, document := range documents {
		outputRow := make([]float64, len(classes))
		bag := document.Sentence.WordsBag(words)

		// Change value to 1 where there is the document Tag
		outputRow[slice.Index(classes, document.Tag)] = 1

		// Append data to trainx and trainy
		inputs = append(inputs, bag)
		targets = append(targets, outputRow)
	}

	return inputs, targets
}

// Returns a new neural network and learn from the TrainData()'s inputs and targets
func CreateNeuralNetwork() (network gonn.NeuralNetwork) {
	trainx, trainy := TrainData()
	inputLayers, outputLayers := len(trainx[0]), len(trainy[0])
	hiddenLayers := int(float64(outputLayers)*30/13 + 0.5)

	network = *gonn.DefaultNetwork(inputLayers, hiddenLayers, outputLayers, true)

	network.Train(trainx, trainy, 1000)

	return network
}
