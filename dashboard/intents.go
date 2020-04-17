package dashboard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gookit/color"
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
	analysis.CacheIntents(intents)

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

	fmt.Printf("Added %s intent.\n", color.FgMagenta.Render(intent.Tag))
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
		fmt.Printf("The intent %s was deleted.\n", color.FgMagenta.Render(intent.Tag))
	}

	WriteIntents(intents)
}

// GetIntents is the route to get the intents
func GetIntents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(analysis.GetIntents())
}

// CreateIntent is the route to create a new intent
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

	if intent.Responses == nil || intent.Patterns == nil {
		json.NewEncoder(w).Encode(Error{
			Message: "Patterns and responses can't be null",
		})
		return
	}

	// Adds the intent
	AddIntent(intent)

	json.NewEncoder(w).Encode(intent)
}

// DeleteIntent is the route used to delete an intent
func DeleteIntent(w http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,Olivia-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

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

	json.NewEncoder(w).Encode(analysis.GetIntents())
}
