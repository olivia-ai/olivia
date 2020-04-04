package modules

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/olivia-ai/olivia/user"

	"golang.org/x/oauth2"

	"github.com/olivia-ai/olivia/language"
	"github.com/zmb3/spotify"
)

var (
	spotifySetterTag = "spotify setter"
	spotifyPlayerTag = "spotify player"

	redirectURL = os.Getenv("REDIRECT_URL")
	callbackURL = os.Getenv("CALLBACK_URL")

	tokenChannel = make(chan *oauth2.Token)
	state        = "abc123"
	auth         = spotify.NewAuthenticator(
		callbackURL,
		spotify.ScopeStreaming,
		spotify.ScopeUserModifyPlaybackState,
	)

	loginMessage = `Login in progress <meta http-equiv="refresh" content="0; url = %s" />`
)

func init() {
	// Set default value of the callback url
	if callbackURL == "" {
		callbackURL = "https://olivia-api.herokuapp.com/callback"
	}

	// Set default value of the redirect url
	if redirectURL == "" {
		redirectURL = "https://olivia-ai.org/chat"
	}

	RegisterModule(Module{
		Tag: spotifySetterTag,
		Patterns: []string{
			"Here are my spotify tokens",
			"My spotify secrets",
		},
		Responses: []string{
			loginMessage,
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
func SpotifySetterReplacer(entry, _, token string) (string, string) {
	spotifyTokens := language.SearchTokens(entry)

	// Returns if the token is empty
	if len(spotifyTokens) != 2 {
		return spotifySetterTag, "You need to send the two tokens."
	}

	// Save the tokens in the user's information
	user.ChangeUserInformation(token, func(information user.Information) user.Information {
		information.SpotifyID = spotifyTokens[0]
		information.SpotifySecret = spotifyTokens[1]

		return information
	})

	return spotifySetterTag, LoginSpotify(token)
}

// SpotifyPlayerReplacer plays a specified music on the user's spotify
// See modules/modules.go#Module.Replacer() for more details.
func SpotifyPlayerReplacer(entry, response, token string) (string, string) {
	// Return if the tokens are not set
	if CheckTokensPresence(token) {
		return spotifySetterTag, "You need to enter your Spotify credentials."
	}

	// Renew the spotify token and get the client
	client := RenewSpotifyToken(token)

	// Search for the track
	music, artist := language.SearchMusic(entry)
	track, err := SearchTrack(client, music+" "+artist)
	if err != nil {
		return spotifySetterTag, LoginSpotify(token)
	}

	// Play the found track
	client.PlayOpt(&spotify.PlayOptions{
		URIs: []spotify.URI{track.URI},
	})
	client.Play()

	return spotifyPlayerTag, fmt.Sprintf(response, track.Name, track.Artists[0].Name)
}

// RenewSpotifyToken renews the spotify token with the user's information token and returns
// the spotify client.
func RenewSpotifyToken(token string) spotify.Client {
	authenticationToken := user.GetUserInformation(token).SpotifyToken
	client := auth.NewClient(authenticationToken)

	// Renew the authentication token
	if m, _ := time.ParseDuration("5m30s"); time.Until(authenticationToken.Expiry) < m {
		user.ChangeUserInformation(token, func(information user.Information) user.Information {
			information.SpotifyToken, _ = client.Token()
			return information
		})
	}

	return client
}

// CheckTokensPresence checks if the spotify tokens are present
func CheckTokensPresence(token string) bool {
	information := user.GetUserInformation(token)
	return information.SpotifyID == "" || information.SpotifySecret == ""
}

// SearchTrack searchs for a given track name and returns the found track and the error
func SearchTrack(client spotify.Client, content string) (spotify.FullTrack, error) {
	// Get the results from a track search with the given content
	results, err := client.Search(content, spotify.SearchTypeTrack)
	if err != nil {
		return spotify.FullTrack{}, err
	}

	// Returns an empty track and empty error if no track was found with this content
	if len(results.Tracks.Tracks) == 0 {
		return spotify.FullTrack{}, nil
	}

	// Return the found
	return results.Tracks.Tracks[0], nil
}

// LoginSpotify logins the user with its token to Spotify
func LoginSpotify(token string) string {
	information := user.GetUserInformation(token)

	// Generate the authentication url
	auth.SetAuthInfo(information.SpotifyID, information.SpotifySecret)
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

	return fmt.Sprintf(loginMessage, url)
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
	fmt.Fprintf(w, `<meta http-equiv="refresh" content="0; url = %s" />`, redirectURL)

	tokenChannel <- token
}
