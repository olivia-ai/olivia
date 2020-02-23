package neuralnet

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
	network := CreateNetwork(3, 4, 1, 0.25, 0.1)
	network.Train(input, output, 1000)

	fmt.Println(network.FeedForward([]float64{1, 1, 1}))
}
