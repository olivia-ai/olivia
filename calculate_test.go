package tests

import (
	"github.com/olivia-ai/olivia/analysis"
	"github.com/olivia-ai/olivia/training"
	gocache "github.com/patrickmn/go-cache"
	"testing"
	"time"
)

func TestCalculate(t *testing.T) {
	model := training.CreateNeuralNetwork()
	cache := gocache.New(5*time.Minute, 5*time.Minute)

	sentences := map[string]string{
		"Hello":                          "hello",
		"How are you ?":                  "feeling",
		"What can you do ?":              "actions",
		"Give me the capital of Namibia": "capital",
	}

	for sentence, tag := range sentences {
		_, responseTag := analysis.Sentence{
			Content: sentence,
		}.Calculate(*cache, model, "1")

		if tag != responseTag {
			t.Errorf("Expected \"%s\" tag for \"%s\", found \"%s\"", tag, sentence, responseTag)
		}
	}
}
