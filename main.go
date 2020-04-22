package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/olivia-ai/olivia/locales"

	"github.com/olivia-ai/olivia/training"

	"github.com/olivia-ai/olivia/dashboard"

	"github.com/olivia-ai/olivia/util"

	"github.com/gookit/color"

	"github.com/olivia-ai/olivia/network"

	"github.com/olivia-ai/olivia/server"
)

var neuralNetworks = map[string]network.Network{}

func main() {
	port := flag.String("port", "8080", "The port for the API and WebSocket.")
	flag.Parse()

	// Print the Olivia ascii text
	oliviaAscii := string(util.ReadFile("res/olivia-ascii.txt"))
	fmt.Println(color.FgLightGreen.Render(oliviaAscii))

	// Create the authentication token
	dashboard.Authenticate()

	for _, locale := range locales.Locales {
		util.SerializeMessages(locale.Tag)

		neuralNetworks[locale.Tag] = training.CreateNeuralNetwork(
			locale.Tag,
			false,
		)
	}

	// Get port from environment variables if there is
	if os.Getenv("PORT") != "" {
		*port = os.Getenv("PORT")
	}

	// Serves the server
	server.Serve(neuralNetworks, *port)
}
