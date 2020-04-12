package dashboard

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/olivia-ai/olivia/analysis"
)

// An Error is what the api replies when an error occurs
type Error struct {
	Message string `json:"message"`
}

// AddIntent adds the given intent to the intents file
func AddIntent(intent analysis.Intent) {
	intents := append(analysis.SerializeIntents(), intent)

	// Encode the json
	bytes, _ := json.MarshalIndent(intents, "", "  ")

	// Write it to the file
	file, err := os.Create("res/datasets/intents.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.Write(bytes)
}

// CreateWebModule is the route to create a new intent
func CreateIntent(w http.ResponseWriter, r *http.Request) {
	// Checks if the token present in the headers is the right one
	token := r.Header.Get("Olivia-Token")
	if !ChecksToken(token) {
		json.NewEncoder(w).Encode(Error{
			Message: "You don't have the permission to do this.",
		})
		return
	}

	// Decode request json body
	var intent analysis.Intent
	json.NewDecoder(r.Body).Decode(&intent)

	// Adds the intent
	AddIntent(intent)

	json.NewEncoder(w).Encode(intent)
}
