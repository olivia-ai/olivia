package modules

import (
	"fmt"
	"log"
	"net/http"

	"github.com/olivia-ai/olivia/user"

	"golang.org/x/oauth2"

	"github.com/olivia-ai/olivia/language"
	"github.com/zmb3/spotify"
)

var (
	spotifySetterTag = "spotify setter"
	spotifyPlayerTag = "spotify player"

	tokenChannel = make(chan *oauth2.Token)
	state        = "abc123"
	auth         = spotify.NewAuthenticator(
		"http://localhost:8080/callback",
		spotify.ScopeStreaming,
		spotify.ScopeUserModifyPlaybackState,
	)
)

func init() {
	RegisterModule(Module{
		Tag: spotifySetterTag,
		Patterns: []string{
			"Here are my spotify tokens",
			"My spotify secrets",
		},
		Responses: []string{
			`Login in progress <meta http-equiv="refresh" content="0; url = %s" />`,
		},
		Replacer: SpotifySetterReplacer,
	})

	RegisterModule(Module{
		Tag: spotifyPlayerTag,
		Patterns: []string{
			"Play from on Spotify",
		},
		Responses: []string{
			"Playing %s from %s on Spotify.",
		},
		Replacer: SpotifyPlayerReplacer,
	})
}

// SpotifySetterReplacer gets the tokens in the user entry and save them into the client's information.
// See modules/modules.go#Module.Replacer() for more details.
func SpotifySetterReplacer(entry, response, token string) (string, string) {
	spotifyTokens := language.SearchTokens(entry)

	// Returns if the token is empty
	if len(spotifyTokens) != 2 {
		return spotifySetterTag, "You need to send the two tokens."
	}

	// Generate the authentication url
	auth.SetAuthInfo(spotifyTokens[0], spotifyTokens[1])
	url := auth.AuthURL(state)

	// Waits for the authentication to be completed, and save the client in user's information
	go func() {
		authenticationToken := <-tokenChannel

		// Save the authentication token
		user.ChangeUserInformation(token, func(information user.Information) user.Information {
			information.SpotifyToken = authenticationToken

			return information
		})
	}()

	return spotifySetterTag, fmt.Sprintf(response, url)
}

// SpotifyPlayerReplacer plays a specified music on the user's spotify
// See modules/modules.go#Module.Replacer() for more details.
func SpotifyPlayerReplacer(entry, response, token string) (string, string) {
	authenticationToken := user.GetUserInformation(token).SpotifyToken
	music, artist := language.SearchMusic(entry)

	client := auth.NewClient(authenticationToken)

	results, _ := client.Search(music+" "+artist, spotify.SearchTypeTrack)
	track := results.Tracks.Tracks[0]

	client.PlayOpt(&spotify.PlayOptions{
		URIs: []spotify.URI{track.URI},
	})
	client.Play()

	return spotifyPlayerTag, fmt.Sprintf(response, track.Name, track.Artists[0].Name)
}

// CompleteAuth completes the Spotify authentication.
func CompleteAuth(w http.ResponseWriter, r *http.Request) {
	token, err := auth.Token(state, r)

	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}

	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		fmt.Printf("State mismatch: %s != %s\n", st, state)
	}

	// Use the token to get an authenticated client
	w.Header().Set("Content-Type", "text/html")
	// Redirect the user
	fmt.Fprintf(w, `<meta http-equiv="refresh" content="0; url = http://localhost:8081/chat" />`)

	tokenChannel <- token
}
