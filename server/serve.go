package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gookit/color"
	"github.com/gorilla/mux"
	"github.com/olivia-ai/olivia/network"
	gocache "github.com/patrickmn/go-cache"
)

var (
	// Create the neural network variable to use it everywhere
	neuralNetwork network.Network
	// Initiatizes the cache with a 5 minute lifetime
	cache = gocache.New(5*time.Minute, 5*time.Minute)
)

// Serve serves the server in the given port
func Serve(_neuralNetwork network.Network, port string) {
	// Set the current global network as a global variable
	neuralNetwork = _neuralNetwork

	// Initializes the router
	router := mux.NewRouter()
	router.HandleFunc("/dashboard", GetDashboardData).Methods("GET")
	router.HandleFunc("/", SocketHandle)

	magenta := color.FgMagenta.Render
	fmt.Printf("\nServer listening on the port %s...\n", magenta(port))

	// Serves the chat
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
