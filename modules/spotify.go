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
	tokenChannel     = make(chan *oauth2.Token)
	state            = "abc123"
	auth             = spotify.NewAuthenticator(
		"http://localhost:8080/callback",
		spotify.ScopeUserReadCurrentlyPlaying,
		spotify.ScopeUserReadPlaybackState,
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
}

// SpotifySetterReplacer gets the tokens in the user entry and save them into the client's information.
// See modules/modules.go#Module.Replacer() for more details.
func SpotifySetterReplacer(entry, response, token string) (string, string) {
	spotifyTokens := language.SearchTokens(entry)

	// Returns if the token is empty
	if len(spotifyTokens) != 2 {
		return spotifySetterTag, "You need to send the two tokens."
	}

	// If the user if already logged in
	if user.GetUserInformation(token).SpotifyToken != (oauth2.Token{}) {
		return spotifySetterTag, "You're already logged in to Spotify."
	}

	auth.SetAuthInfo(spotifyTokens[0], spotifyTokens[1])
	url := auth.AuthURL(state)

	// Waits for the authentication to be completed, and save the client in user's information
	go func() {
		authenticationToken := <-tokenChannel

		// Save the token
		user.ChangeUserInformation(token, func(information user.Information) user.Information {
			information.SpotifyToken = *authenticationToken
			return information
		})
	}()

	return spotifySetterTag, fmt.Sprintf(response, url)
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
