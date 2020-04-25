package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/olivia-ai/olivia/network"
)

// Dashboard contains the data sent for the dashboard
type Dashboard struct {
	Layers   Layers   `json:"layers"`
	Training Training `json:"training"`
}

// Layers contains the data of the network's layers
type Layers struct {
	InputNodes   int `json:"input"`
	HiddenLayers int `json:"hidden"`
	OutputNodes  int `json:"output"`
}

// Training contains the data related to the training of the network
type Training struct {
	Rate   float64   `json:"rate"`
	Errors []float64 `json:"errors"`
	Time   float64   `json:"time"`
}

// GetDashboardData encodes the json for the dashboard data
func GetDashboardData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := mux.Vars(r)

	dashboard := Dashboard{
		Layers:   GetLayers(data["locale"]),
		Training: GetTraining(data["locale"]),
	}

	err := json.NewEncoder(w).Encode(dashboard)
	if err != nil {
		log.Fatal(err)
	}
}

// GetLayers returns the number of input, hidden and output layers of the network
func GetLayers(locale string) Layers {
	return Layers{
		// Get the number of rows of the first layer to get the count of input nodes
		InputNodes: network.Rows(neuralNetworks[locale].Layers[0]),
		// Get the number of hidden layers by removing the count of the input and output layers
		HiddenLayers: len(neuralNetworks[locale].Layers) - 2,
		// Get the number of rows of the latest layer to get the count of output nodes
		OutputNodes: network.Columns(neuralNetworks[locale].Output),
	}
}

// GetTraining returns the learning rate, training date and error loss for the network
func GetTraining(locale string) Training {
	// Retrieve the information from the neural network
	return Training{
		Rate:   neuralNetworks[locale].Rate,
		Errors: neuralNetworks[locale].Errors,
		Time:   neuralNetworks[locale].Time,
	}
}
