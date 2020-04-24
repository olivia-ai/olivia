package spotify

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/olivia-ai/olivia/util"

	"github.com/olivia-ai/olivia/user"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

var (
	redirectURL = os.Getenv("REDIRECT_URL")
	callbackURL = os.Getenv("CALLBACK_URL")

	tokenChannel = make(chan *oauth2.Token)
	state        = "abc123"
	auth         spotify.Authenticator
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

	// Initialize the authenticator
	auth = spotify.NewAuthenticator(
		callbackURL,
		spotify.ScopeStreaming,
		spotify.ScopeUserModifyPlaybackState,
		spotify.ScopeUserReadPlaybackState,
	)
}

// LoginSpotify logins the user with its token to Spotify
func LoginSpotify(locale, token string) string {
	information := user.GetUserInformation(token)

	// Generate the authentication url
	auth.SetAuthInfo(information.SpotifyID, information.SpotifySecret)
	url := auth.AuthURL(state)

	// Waits for the authentication to be completed, and save the client in user's information
	go func() {
		authenticationToken := <-tokenChannel

		// If the token is empty reset the credentials of the user
		if *authenticationToken == (oauth2.Token{}) {
			user.ChangeUserInformation(token, func(information user.Information) user.Information {
				information.SpotifyID = ""
				information.SpotifySecret = ""

				return information
			})
		}

		// Save the authentication token
		user.ChangeUserInformation(token, func(information user.Information) user.Information {
			information.SpotifyToken = authenticationToken

			return information
		})
	}()

	return fmt.Sprintf(util.GetMessage(locale, "spotify login"), url)
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

// CompleteAuth completes the Spotify authentication.
func CompleteAuth(w http.ResponseWriter, r *http.Request) {
	token, err := auth.Token(state, r)

	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		tokenChannel <- &oauth2.Token{}
		return
	}

	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		tokenChannel <- &oauth2.Token{}
		return
	}

	// Use the token to get an authenticated client
	w.Header().Set("Content-Type", "text/html")
	// Redirect the user
	fmt.Fprintf(w, `<meta http-equiv="refresh" content="0; url = %s" />`, redirectURL)

	tokenChannel <- token
}
