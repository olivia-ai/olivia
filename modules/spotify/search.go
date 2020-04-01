package spotify

import (
	"encoding/json"

	"github.com/rapito/go-spotify/spotify"
)

// Results of the Search of tracks
type Search struct {
	Tracks Tracks `json:"tracks"`
}

// A Track is found on the Search
type Tracks struct {
	Items []Item `json:"items"`
}

// An Item is an item found on the search of Tracks
type Item struct {
	URI     string   `json:"uri"`
	Name    string   `json:"name"`
	Artists []Artist `json:"artists"`
}

// Artist is one of the artists of Item
type Artist struct {
	Name string `json:"name"`
}

// SearchTracks returns the first track found with the given name and artist name
func SearchTracks(client spotify.Spotify, name, artist string) Item {
	// Execute the get request to the spotify API
	search, _ := client.Get(
		"search?q=%s&type=track&limit=1",
		nil,
		name+" "+artist,
	)

	// Unmarshal the response into the structure
	response := Search{}
	err := json.Unmarshal(search, &response)
	if err != nil {
		return Item{}
	}

	// Returns the first and only track
	return response.Tracks.Items[0]
}
