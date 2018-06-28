package main

import (
	"github.com/olivia-ai/Api/analysis"
	"github.com/olivia-ai/Api/training"
	"github.com/gorilla/mux"
	gocache "github.com/patrickmn/go-cache"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Content string `json:"content"`
}

var (
	model = training.CreateNeuralNetwork()
	cache = gocache.New(5*time.Minute, 5*time.Minute)
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/response", PostResponse).Methods("POST")

	fmt.Println("Listening on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func PostResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	responseSentence := analysis.Sentence{
		Content: r.FormValue("sentence"),
	}.Calculate(*cache, model, r.FormValue("authorId"))

	// Marshall the response in json
	response := Response{responseSentence}
	bytes, err := json.Marshal(response)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(bytes))
}
