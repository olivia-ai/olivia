package main

import (
	"regexp"
	"testing"
	"time"

	"github.com/olivia-ai/olivia/analysis"
	"github.com/olivia-ai/olivia/training"
	gocache "github.com/patrickmn/go-cache"
)

func TestCalculate(t *testing.T) {
	model := training.CreateNeuralNetwork("en", true)
	cache := gocache.New(5*time.Minute, 5*time.Minute)

	sentences := map[string]string{
		"Hello!":                                   "hello",
		"How are you today ?":                      "feeling",
		"Who are you ?":                            "identity",
		"What is your job?":                        "job",
		"What is your age":                         "age",
		"Where do you live":                        "city",
		"What can you do ?":                        "actions",
		"what's the capital of France":             "capital",
		"Give me the capital of Namibia":           "capital",
		"I'm bored":                                "movies search from data|no genres saved",
		"Give me a random number":                  "random number",
		"What are you doing ?":                     "action",
		"Calculate cos(0.5)":                       "math",
		"Can you help me ?":                        "actions",
		"My name is Hugo":                          "name setter",
		"Please wait 2 minutes":                    "wait",
		"Are you still there?":                     "still there",
		"I like movies of adventure and animation": "movies genres",
		"Can you find me a movie of adventure":     "movies search",
		"Remind me to call mom tomorrow at 9pm":    "reminder setter",
		"List my reminders?":                       "reminder getter",
	}

	for sentence, tag := range sentences {
		responseTag, _ := analysis.NewSentence("en", sentence).Calculate(*cache, model, "1")

		if !regexp.MustCompile(tag).MatchString(responseTag) {
			t.Errorf("Expected \"%s\" tag for \"%s\", found \"%s\"", tag, sentence, responseTag)
		}
	}
}
