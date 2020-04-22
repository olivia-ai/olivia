package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/olivia-ai/olivia/training"

	"github.com/olivia-ai/olivia/dashboard"

	"github.com/olivia-ai/olivia/util"

	"github.com/gookit/color"

	"github.com/olivia-ai/olivia/network"

	"github.com/olivia-ai/olivia/server"
)

var neuralNetworks = map[string]network.Network{}

// A Locale is a registered locale in the file
type Locale struct {
	Locale string `json:"locale"`
	Name   string `json:"name"`
}

// SerializeLocales returns the locales retrieved from the locales JSON file
func SerializeLocales() (locales []Locale) {
	err := json.Unmarshal(util.ReadFile("res/locales/locales.json"), &locales)
	if err != nil {
		panic(err)
	}

	return
}

func main() {
	port := flag.String("port", "8080", "The port for the API and WebSocket.")
	flag.Parse()

	// Print the Olivia ascii text
	oliviaAscii := string(util.ReadFile("res/olivia-ascii.txt"))
	fmt.Println(color.FgLightGreen.Render(oliviaAscii))

	// Create the authentication token
	dashboard.Authenticate()

	for _, locale := range SerializeLocales() {
		neuralNetworks[locale.Locale] = training.CreateNeuralNetwork(
			locale.Locale,
			true,
		)
	}

	// Get port from environment variables if there is
	if os.Getenv("PORT") != "" {
		*port = os.Getenv("PORT")
	}

	// Serves the server
	server.Serve(neuralNetworks, *port)
}
