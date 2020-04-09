package dashboard

import (
	"encoding/json"
	"net/http"

	"github.com/olivia-ai/olivia/analysis"
)

func GetIntents(w http.ResponseWriter, r *http.Request) {
	intents := analysis.SerializeIntents("res/intents.json")
	json.NewEncoder(w).Encode(intents)
}
