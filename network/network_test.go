package network

import (
	"fmt"
	"testing"

	"github.com/olivia-ai/olivia/analysis"
	"github.com/olivia-ai/olivia/training"
)

func TestCreateNetwork(t *testing.T) {
	input, output := training.TrainData()
	words, classes, _ := analysis.Organize()

	network := CreateNetwork(0.1, input, output, 50)
	network.Train(1000)

	a := network.FeedForwardWithValue(
		analysis.NewSentence("what is the capital of france").WordsBag(words),
	)[0]

	var max string
	var maxi float64 = 0
	for i := 0; i < len(a); i++ {
		if a[i] > maxi {
			maxi = a[i]
			max = classes[i]
		}
	}

	fmt.Println(max)
	fmt.Println(maxi / 0.1)
}
