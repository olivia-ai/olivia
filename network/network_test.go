package network

import (
	"testing"
)

func TestCreateNetwork(t *testing.T) {
	input := [][]float64{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	output := [][]float64{
		{1, 1},
		{1, 1},
		{1, 1},
	}

	net := CreateNetwork(input, output, 4)
	for i := 0; i < 1000; i++ {
		net.FeedForward()
		net.FeedBackward()
	}
}
