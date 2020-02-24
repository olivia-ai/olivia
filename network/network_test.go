package network

import (
	"fmt"
	"testing"
)

func TestCreateNetwork(t *testing.T) {
	input := [][]float64{
		{1, 1, 1},
		{0, 1, 1},
		{0, 0, 1},
		{1, 0, 1},
		{1, 1, 0},
		{0, 0, 0},
	}

	output := [][]float64{
		{1},
		{1},
		{0},
		{1},
		{1},
		{0},
	}

	network := CreateNetwork(0.25, input, output, 4, 4)
	network.Train(1000)

	network.FeedForward()
	fmt.Println(network.Layers[3])
}
