package network

import (
	"fmt"
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

	fmt.Println(CreateNetwork(input, output, 4))
}
