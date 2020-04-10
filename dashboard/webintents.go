package dashboard

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/olivia-ai/olivia/util"
)

const webIntentsFile = "res/api/webintents.json"

// A WebIntent is a dynamic intent which retrieves the data from an URL
type WebIntent struct {
	Tag       string   `json:"tag"`
	Patterns  []string `json:"patterns"`
	Responses []string `json:"responses"`
	Context   string   `json:"context"`
	GetURL    string   `json:"get_url"`
}

// SerializeWebIntents parses the content of the web intents file and returns the array
// of WebIntents
func SerializeWebIntents() []WebIntent {
	var webIntents []WebIntent

	err := json.Unmarshal(util.ReadFile(webIntentsFile), &webIntents)
	if err != nil {
		panic(err)
	}

	return webIntents
}

// GetWebIntents is the route to get the web intents
func GetWebIntents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(SerializeWebIntents())
}

// PostWebIntent is the route to create a WebIntent
func PostWebIntent(w http.ResponseWriter, r *http.Request) {
	// Get the body content
	var webIntent WebIntent
	json.NewDecoder(r.Body).Decode(&webIntent)

	// Add the web intents to the existing ones
	webIntents := SerializeWebIntents()
	webIntents = append(webIntents, webIntent)

	// Write the web intents file with the new content
	outF, _ := os.OpenFile(webIntentsFile, os.O_CREATE|os.O_RDWR, 0777)
	defer outF.Close()

	encoder := json.NewEncoder(outF)
	err := encoder.Encode(webIntents)
	if err != nil {
		panic(err)
	}

	// Reply with the existing web intents
	json.NewEncoder(w).Encode(webIntents)
}
