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
	port := flag.String("port", "8080", "The port for the API and WebSocket.")
	flag.Parse()

	// Set default value of the callback url
	if os.Getenv("CALLBACK_URL") == "" {
		os.Setenv("CALLBACK_URL", "https://olivia-api.herokuapp.com/callback")
	}

	// Set default value of the redirect url
	if os.Getenv("REDIRECT_URL") == "" {
		os.Setenv("REDIRECT_URL", "https://olivia-ai.org/chat")
	}

	magenta := color.FgMagenta.Render
	fmt.Printf("Using %s as intents file.\n", magenta(*intentsPath))

	neuralNetwork = training.CreateNeuralNetwork(*intentsPath)

	// Get port from environment variables if there is
	if os.Getenv("PORT") != "" {
		*port = os.Getenv("PORT")
	}

	// Serves the server
	server.Serve(neuralNetwork, *port, *intentsPath)
}
