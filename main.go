package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gookit/color"

	"github.com/olivia-ai/olivia/network"

	"github.com/olivia-ai/olivia/server"
	"github.com/olivia-ai/olivia/training"
)

var neuralNetwork network.Network

func main() {
	intentsPath := flag.String("intents", "res/intents.json", "The path for intents file.")
	flag.Parse()

	magenta := color.FgMagenta.Render
	fmt.Printf("Using %s as intents file.\n", magenta(*intentsPath))

	neuralNetwork = training.CreateNeuralNetwork(*intentsPath)

	port := "8080"
	// Get port from environment variables if there is
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	// Serves the server
	server.Serve(neuralNetwork, port, *intentsPath)
}
