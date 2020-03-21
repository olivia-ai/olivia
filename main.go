package main

import (
	"github.com/olivia-ai/olivia/chat"
	"github.com/olivia-ai/olivia/dashboard"
	"github.com/olivia-ai/olivia/training"
)

var (
	// Initialize the neural network by training it
	neuralNetwork = training.CreateNeuralNetwork()
)

func main() {
	// Serve the REST API inside a go routine
	go func() {
		dashboard.Serve(neuralNetwork, "8081")
	}()

	// Serves the chat
	chat.Serve(neuralNetwork, "8080")
}
