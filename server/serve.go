package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/olivia-ai/olivia/training"

	"github.com/olivia-ai/olivia/dashboard"

	"github.com/olivia-ai/olivia/modules/spotify"

	"github.com/gookit/color"
	"github.com/gorilla/mux"
	"github.com/olivia-ai/olivia/network"
	gocache "github.com/patrickmn/go-cache"
)

var (
	// Create the neural network variable to use it everywhere
	neuralNetwork network.Network
	// Initializes the cache with a 5 minute lifetime
	cache = gocache.New(5*time.Minute, 5*time.Minute)
)

// Serve serves the server in the given port
func Serve(_neuralNetwork network.Network, port string) {
	// Set the current global network as a global variable
	neuralNetwork = _neuralNetwork

	// Initializes the router
	router := mux.NewRouter()
	router.HandleFunc("/callback", spotify.CompleteAuth)
	// Serve the websocket
	router.HandleFunc("/websocket", SocketHandle)
	// Serve the API
	router.HandleFunc("/api/dashboard", GetDashboardData).Methods("GET")
	router.HandleFunc("/api/intent", dashboard.CreateIntent).Methods("POST")
	router.HandleFunc("/api/intent", dashboard.DeleteIntent).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/train", Train).Methods("POST")
	router.HandleFunc("/api/intents", dashboard.GetIntents).Methods("GET")

	magenta := color.FgMagenta.Render
	fmt.Printf("\nServer listening on the port %s...\n", magenta(port))

	// Serves the chat
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}

// Train is the route to re-train the neural network
func Train(w http.ResponseWriter, r *http.Request) {
	// Checks if the token present in the headers is the right one
	token := r.Header.Get("Olivia-Token")
	if !dashboard.ChecksToken(token) {
		json.NewEncoder(w).Encode(dashboard.Error{
			Message: "You don't have the permission to do this.",
		})
		return
	}

	magenta := color.FgMagenta.Render
	fmt.Printf("\nRe-training the %s..\n", magenta("neural network"))
	neuralNetwork = training.CreateNeuralNetwork(true)
}
