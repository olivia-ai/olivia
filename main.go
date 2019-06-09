package main

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/olivia-ai/olivia/analysis"
	"github.com/olivia-ai/olivia/training"
	gocache "github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

var (
	model = training.CreateNeuralNetwork()
	cache = gocache.New(5*time.Minute, 5*time.Minute)
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/response", PostResponse).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})

	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	magenta := color.FgMagenta.Render
	fmt.Printf("\nListening on the port %s...\n", magenta(port))
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func PostResponse(w http.ResponseWriter, r *http.Request) {
	responseSentence, responseTag := analysis.Sentence{
		Content: r.FormValue("sentence"),
	}.Calculate(*cache, model, r.FormValue("authorId"))

	// Marshall the response in json
	response := Response{responseSentence, responseTag}
	bytes, err := json.Marshal(response)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(bytes))
}
