package start

import (
	modules2 "github.com/olivia-ai/olivia/modules"
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
	if modules2.CheckTokensPresence(token) {
		return
	}

	client := modules2.RenewSpotifyToken(token)

	// Test to search a track to see if the user is well logged in
	_, err := modules2.SearchTrack(client, "test")
	// If an error is present, login the user to spotify
	if err != nil {
		SetMessage(modules2.LoginSpotify(token))
	}
}
