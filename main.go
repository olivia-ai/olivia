package main

import (
	"encoding/json"
	"fmt"
	"github.com/olivia-ai/Api/analysis"
	"github.com/olivia-ai/Api/training"
	"google.golang.org/appengine"
	gocache "github.com/patrickmn/go-cache"
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
	http.HandleFunc("/api/response", PostResponse)
	appengine.Main()
}

func PostResponse(w http.ResponseWriter, r *http.Request) {
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
