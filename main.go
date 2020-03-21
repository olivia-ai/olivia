package main

import (
	"os"

	"github.com/olivia-ai/olivia/server"
	"github.com/olivia-ai/olivia/training"
)

var (
	// Initialize the neural network by training it
	neuralNetwork = training.CreateNeuralNetwork()
)

func main() {
	port := "8080"
	// Get port from environment variables if there is
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	// Serves the server
	server.Serve(neuralNetwork, port)
}
