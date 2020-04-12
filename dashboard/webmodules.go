package dashboard

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/olivia-ai/olivia/util"
)

const webModulesFile = "res/webmodules.json"

// A WebModule is a module that allows you to retrieve information from a URL
type WebModule struct {
	Tag        string   `json:"tag"`
	Patterns   []string `json:"patterns"`
	Responses  []string `json:"responses"`
	Context    string   `json:"context"`
	URL        string   `json:"url"`
	Parameters []string `json:"parameters"`
}

// An Error is what the api replies when an error occurs
type Error struct {
	Message string `json:"message"`
}

// AddWebModules adds the given webmodule to the webmodules file
func AddWebModule(module WebModule) {
	// Retrieve content from the web modules file
	var webModules []WebModule
	fileContent := util.ReadFile(webModulesFile)
	json.Unmarshal(fileContent, &webModules)

	webModules = append(webModules, module)

	// Encode the json
	bytes, _ := json.Marshal(webModules)

	// Write it to the file
	file, err := os.Create(webModulesFile)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.Write(bytes)
}

// CreateWebModule is the route to create a new webmodule
func CreateWebModule(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Olivia-Token")
	if !ChecksToken(token) {
		json.NewEncoder(w).Encode(Error{
			Message: "You don't have the permission to do this.",
		})
		return
	}

	// Decode request json body
	var webModule WebModule
	json.NewDecoder(r.Body).Decode(&webModule)

	// Adds the web module
	AddWebModule(webModule)

	json.NewEncoder(w).Encode(webModule)
}
