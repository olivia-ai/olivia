package main

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
		"what's the capital of France":   "capital",
		"Give me the capital of Namibia": "capital",
		"What is your name ?":            "name",
		"What's your name ?":             "name",
		"Where do you live?":             "city",
		"What are you doing ?":           "action",
		"Calculate cos(0.5)":             "math",
		"Can you help me ?":              "actions",
		"Why is your name Olivia?":       "why name",
	}

	for sentence, tag := range sentences {
		_, responseTag := analysis.NewSentence(sentence).Calculate(*cache, model, "1")

		if tag != responseTag {
			t.Errorf("Expected \"%s\" tag for \"%s\", found \"%s\"", tag, sentence, responseTag)
		}
	}
}
