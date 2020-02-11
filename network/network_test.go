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

func TestMatrix_DotProduct(t *testing.T) {
	a := Matrix{[][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}}

	b := Matrix{[][]float64{
		{7, 8},
		{9, 10},
		{11, 12},
	}}

	fmt.Println(a.DotProduct(b))
}
