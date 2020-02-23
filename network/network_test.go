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

	net := CreateNetwork(0.1, input, output, 4)
	for i := 0; i < 3000; i++ {
		net.FeedForward()
		net.FeedBackward()
	}

	net.FeedForward()
	fmt.Println(net.Layers[2].Outputs)
}
