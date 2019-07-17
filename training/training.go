package training

import (
	"github.com/gookit/color"
	"github.com/olivia-ai/olivia/analysis"
	"github.com/olivia-ai/olivia/neuralnet"
	"github.com/olivia-ai/olivia/util"
	"os"
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
func CreateNeuralNetwork() (network neuralnet.NeuralNetwork) {
	// Decide if the network is created by the save or is a new one
	saveFile := "res/training.json"

	_, err := os.Open(saveFile)
	if err != nil {
		// Train the model if there is no training file
		trainx, trainy := TrainData()
		inputLayers, outputLayers := len(trainx[0]), len(trainy[0])
		hiddenLayers := 100

		network = *neuralnet.CreateNetwork(inputLayers, hiddenLayers, outputLayers, 0.25, 0.1)
		network.Train(trainx, trainy, 2000)

		// Save the neural network in res/training.json
		neuralnet.DumpNN(saveFile, &network)
	} else {
		color.FgBlue.Println("Loading the neural network from " + saveFile)
		network = *neuralnet.LoadNN(saveFile)
	}

	return network
}
