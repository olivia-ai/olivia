package modules

import (
	"fmt"

	"github.com/olivia-ai/olivia/user"

	"github.com/olivia-ai/olivia/language"
	"github.com/zmb3/spotify"

	spotifyModule "github.com/olivia-ai/olivia/modules/spotify"
)

var (
	spotifySetterTag = "spotify setter"
	spotifyPlayerTag = "spotify player"

	loginMessage = `Login in progress <meta http-equiv="refresh" content="0; url = %s" />`
)

func init() {
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

	return spotifySetterTag, spotifyModule.LoginSpotify(token, loginMessage)
}

// SpotifyPlayerReplacer plays a specified music on the user's spotify
// See modules/modules.go#Module.Replacer() for more details.
func SpotifyPlayerReplacer(entry, response, token string) (string, string) {
	// Return if the tokens are not set
	if spotifyModule.CheckTokensPresence(token) {
		return spotifySetterTag, "You need to enter your Spotify credentials."
	}

	// Renew the spotify token and get the client
	client := spotifyModule.RenewSpotifyToken(token)

	// Search for the track
	music, artist := language.SearchMusic(entry)
	track, err := SearchTrack(client, music+" "+artist)
	if err != nil {
		return spotifySetterTag, spotifyModule.LoginSpotify(token, loginMessage)
	}

	// Play the found track
	client.PlayOpt(&spotify.PlayOptions{
		URIs: []spotify.URI{track.URI},
	})
	client.Play()

	return spotifyPlayerTag, fmt.Sprintf(response, track.Name, track.Artists[0].Name)
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
