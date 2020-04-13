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

// DeleteRequest is for the parameters required to delete an intent via the REST Api
type DeleteRequest struct {
	Tag string `json:"tag"`
}

// WriteIntents writes the given intents to the intents file
func WriteIntents(intents []analysis.Intent) {
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

// AddIntent adds the given intent to the intents file
func AddIntent(intent analysis.Intent) {
	intents := append(analysis.SerializeIntents(), intent)

	WriteIntents(intents)
}

// RemoveIntent removes the intent with the given tag from the intents file
func RemoveIntent(tag string) {
	intents := analysis.SerializeIntents()

	// Iterate through the intents to remove the right one
	for i, intent := range intents {
		if intent.Tag != tag {
			continue
		}

		intents[i] = intents[len(intents)-1]
		intents = intents[:len(intents)-1]
	}

	WriteIntents(intents)
}

// GetIntents is the route to get the intents
func GetIntents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(analysis.SerializeIntents())
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

// DeleteIntent is the route used to delete an intent
func DeleteIntent(w http.ResponseWriter, r *http.Request) {
	// Checks if the token present in the headers is the right one
	token := r.Header.Get("Olivia-Token")
	if !ChecksToken(token) {
		json.NewEncoder(w).Encode(Error{
			Message: "You don't have the permission to do this.",
		})
		return
	}

	var deleteRequest DeleteRequest
	json.NewDecoder(r.Body).Decode(&deleteRequest)

	RemoveIntent(deleteRequest.Tag)

	json.NewEncoder(w).Encode(analysis.SerializeIntents())
}
