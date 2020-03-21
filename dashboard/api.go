package dashboard

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/olivia-ai/olivia/network"
)

var neuralNetwork network.Network

type Layers struct {
	InputNodes   int `json:"input"`
	HiddenLayers int `json:"hidden"`
	OutputNodes  int `json:"output"`
}

// Serve serves the dashboard REST API on the port 8081 by default.
func Serve(_neuralNetwork network.Network) {
	// Set the current global network as a global variable
	neuralNetwork = _neuralNetwork

	router := mux.NewRouter()

	// Create the routes
	router.HandleFunc("/dashboard/layers", GetLayers).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))
}

// GetLayers encode the number of input, hidden and output layers in json
func GetLayers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	layers := Layers{
		// Get the number of rows of the first layer to get the count of input nodes
		InputNodes: network.Rows(neuralNetwork.Layers[0]),
		// Get the number of hidden layers by removing the count of the input and output layers
		HiddenLayers: len(neuralNetwork.Layers) - 2,
		// Get the number of rows of the latest layer to get the count of output nodes
		OutputNodes: network.Rows(neuralNetwork.Output),
	}

	err := json.NewEncoder(w).Encode(layers)
	if err != nil {
		log.Fatal(err)
	}
}
