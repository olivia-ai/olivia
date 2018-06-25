package main

import (
	"github.com/gorilla/mux"
	"./analysis"
	"./training"
	"./cache"
	"net/http"
	"log"
	"fmt"
	"encoding/json"
)

type Response struct {
	Content string `json:"content"`
}

var (
	model = training.CreateNeuralNetwork()
	redis = cache.CreateClient()
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/response", PostResponse).Methods("POST")

	fmt.Println("Listening on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func PostResponse(w http.ResponseWriter, r *http.Request) {
	responseSentence := analysis.Sentence{
		Content: r.FormValue("sentence"),
	}.Calculate(redis, model, r.FormValue("authorId"))

	// Marshall the response in json
	response := Response{responseSentence}
	bytes, err := json.Marshal(response)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(bytes))
}
