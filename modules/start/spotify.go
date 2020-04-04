package start

import (
	_modules "github.com/olivia-ai/olivia/modules"
	"github.com/olivia-ai/olivia/modules/spotify"
)

func init() {
	RegisterModule(Module{
		Action: CheckSpotifyLogin,
	})
}

// CheckSpotifyLogin checks if the user is well logged in into spotify and if not,
// it logs the user in
func CheckSpotifyLogin(token string) {
	// Return if the user was not logged in before
	if spotify.CheckTokensPresence(token) {
		return
	}

	client := spotify.RenewSpotifyToken(token)

	// Test to search a track to see if the user is well logged in
	_, err := _modules.SearchTrack(client, "test")
	// If an error is present, login the user to spotify
	if err != nil {
		SetMessage(spotify.LoginSpotify(token))
	}
}
