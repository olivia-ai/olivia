package network

type Network struct {
	layers  []Matrix
	weights []Matrix
	biases  []Matrix
}

// CreateNetwork creates the network by generating the layers, weights and biases
func CreateNetwork(input, output [][]float64, hiddensNodes ...int) Network {
	layers := []Matrix{
		{input},
	}

	// Generate the hidden layers
	for _, hiddenNodes := range hiddensNodes {
		layers = append(layers, CreateMatrix(len(input), hiddenNodes))
	}

	layers = append(layers, Matrix{output})

	// Generate the weights
	weightsNumber := 1 + len(hiddensNodes)
	var weights []Matrix
	var biases []Matrix

	for i := 0; i < weightsNumber; i++ {
		rows, columns := layers[i+1].Rows(), layers[i].Rows()
		weights = append(weights, RandomMatrix(rows, columns))

		biases = append(biases, RandomMatrix(layers[i].Rows(), layers[i+1].Rows()))
	}

	return Network{
		layers:  layers,
		weights: weights,
		biases:  biases,
	}
}

func (network *Network) FeedForward(input [][]float64) {
	if len(input) != network.layers[0].Rows() {
		panic("Amount of input variable does not match.")
	}
}
