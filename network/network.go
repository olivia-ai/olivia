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
		rows, columns := layers[i].Columns(), layers[i+1].Columns()
		weights = append(weights, RandomMatrix(rows, columns))

		biases = append(biases, RandomMatrix(layers[i].Rows(), columns))
	}

	return Network{
		layers:  layers,
		weights: weights,
		biases:  biases,
	}
}

func (network *Network) FeedForward() {
	for i := 0; i < len(network.layers)-1; i++ {
		layer, weights, biases := network.layers[i], network.weights[i], network.biases[i]
		productMatrix := layer.DotProduct(weights)
		productMatrix.Add(biases)
		productMatrix.ApplyFunction(Sigmoid)

		network.layers[i+1] = productMatrix
	}
}

func (network *Network) FeedBackward() {

}
