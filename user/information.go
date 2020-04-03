package user

import (
	"golang.org/x/oauth2"
)

// Information is the user's information retrieved from the client
type Information struct {
	Name           string        `json:"name"`
	MovieGenres    []string      `json:"movie_genres"`
	MovieBlacklist []string      `json:"movie_blacklist"`
	Reminders      []Reminder    `json:"reminders"`
	SpotifyToken   *oauth2.Token `json:"spotify_token"`
	SpotifyID      string        `json:"spotify_id"`
	SpotifySecret  string        `json:"spotify_secret"`
}

// A Reminder is something the user asked for Olivia to remember
type Reminder struct {
	Reason string `json:"reason"`
	Date   string `json:"date"`
}

// userInformation is a map which is the cache for user information
var userInformation = map[string]Information{}

// ChangeUserInformation requires the token of the user and a function which gives the actual
// information and returns the new information.
func ChangeUserInformation(token string, changer func(Information) Information) {
	userInformation[token] = changer(userInformation[token])
}

// SetUserInformation sets the user's information by its token.
func SetUserInformation(token string, information Information) {
	userInformation[token] = information
}

// GetUserInformation returns the information of a user with his token
func GetUserInformation(token string) Information {
	return userInformation[token]
}
