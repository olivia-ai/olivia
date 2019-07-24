package user

type Information struct {
	Name           string   `json:"name"`
	MovieGenres    []string `json:"movie_genres"`
	MovieBlacklist []string `json:"movie_blacklist"`
}

var userInformation = map[string]Information{}

// ChangeUserInformations requires the token of the user and a function which gives the actual
// informations and returns the new informations.
func ChangeUserInformations(token string, changer func(Information) Information) {
	userInformation[token] = changer(userInformation[token])
}

func SetUserInformations(token string, information Information) {
	userInformation[token] = information
}

// GetUserInformations returns the informations of a user with his token
func GetUserInformations(token string) Information {
	return userInformation[token]
}
