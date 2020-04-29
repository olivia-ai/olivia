package modules

import (
	"fmt"
	"strings"

	"github.com/olivia-ai/olivia/util"

	"github.com/olivia-ai/olivia/user"

	"github.com/olivia-ai/olivia/language"
	"github.com/zmb3/spotify"

	spotifyModule "github.com/olivia-ai/olivia/modules/spotify"
)

var (
	// SpotifySetterTag is the intent tag for its module
	SpotifySetterTag = "spotify setter"
	// SpotifyPlayerTag is the intent tag for its module
	SpotifyPlayerTag = "spotify player"
)

// SpotifySetterReplacer gets the tokens in the user entry and save them into the client's information.
// See modules/modules.go#Module.Replacer() for more details.
func SpotifySetterReplacer(locale, entry, _, token string) (string, string) {
	spotifyTokens := language.SearchTokens(entry)

	// Returns if the token is empty
	if len(spotifyTokens) != 2 {
		return SpotifySetterTag, util.GetMessage(locale, "spotify tokens")
	}

	// Save the tokens in the user's information
	user.ChangeUserInformation(token, func(information user.Information) user.Information {
		information.SpotifyID = spotifyTokens[0]
		information.SpotifySecret = spotifyTokens[1]

		return information
	})

	return SpotifySetterTag, spotifyModule.LoginSpotify(locale, token)
}

// SpotifyPlayerReplacer plays a specified music on the user's spotify
// See modules/modules.go#Module.Replacer() for more details.
func SpotifyPlayerReplacer(locale, entry, response, token string) (string, string) {
	// Return if the tokens are not set
	if spotifyModule.CheckTokensPresence(token) {
		return SpotifySetterTag, util.GetMessage(locale, "spotify credentials")
	}

	// Renew the spotify token and get the client
	client := spotifyModule.RenewSpotifyToken(token)

	// Search for the track
	music, artist := language.SearchMusic(locale, entry)
	track, err := SearchTrack(client, music+" "+artist)
	if err != nil {
		return SpotifySetterTag, spotifyModule.LoginSpotify(locale, token)
	}

	// Search if there is a device name in the entry
	device := SearchDevice(client, entry)
	options := &spotify.PlayOptions{
		URIs: []spotify.URI{track.URI},
	}

	// Add the device ID if a device is contained
	if device != (spotify.PlayerDevice{}) {
		options.DeviceID = &device.ID
	}

	// Play the found track
	client.PlayOpt(options)
	client.Play()

	return SpotifyPlayerTag, fmt.Sprintf(response, track.Name, track.Artists[0].Name)
}

// SearchTrack searches for a given track name and returns the found track and the error
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

// SearchDevice searches for a device name inside the given sentence and returns it
func SearchDevice(client spotify.Client, content string) spotify.PlayerDevice {
	devices, _ := client.PlayerDevices()

	// Iterate through the devices to check if the content contains a device name
	for _, device := range devices {
		if strings.Contains(content, strings.ToLower(device.Name)) ||
			strings.Contains(content, strings.ToLower(device.Type)) {
			return device
		}
	}

	return spotify.PlayerDevice{}
}
