package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/olivia-ai/olivia/modules"

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
	// Set the intents file path
	intentsPath string
)

// Serve serves the server in the given port
func Serve(_neuralNetwork network.Network, port, _intentsPath string) {
	// Set the current global network as a global variable
	neuralNetwork = _neuralNetwork
	intentsPath = _intentsPath

	// Initializes the router
	router := mux.NewRouter()
	router.HandleFunc("/dashboard", GetDashboardData).Methods("GET")
	router.HandleFunc("/", SocketHandle)
	router.HandleFunc("/callback", modules.CompleteAuth)

	magenta := color.FgMagenta.Render
	fmt.Printf("\nServer listening on the port %s...\n", magenta(port))

	// Serves the chat
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
