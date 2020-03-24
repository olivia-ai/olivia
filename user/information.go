package user

// Information is the user's information retrieved from the client
type Information struct {
	Name           string   `json:"name"`
	MovieGenres    []string `json:"movie_genres"`
	MovieBlacklist []string `json:"movie_blacklist"`
}

// userInformation is a map which is the cache for user information
var userInformation = map[string]Information{}

// ChangeUserInformation requires the token of the user and a function which gives the actual
// information and returns the new information.
func ChangeUserInformation(token string, changer func(Information) Information) {
	userInformation[token] = changer(userInformation[token])
}

func SetUserInformation(token string, information Information) {
	userInformation[token] = information
}

// GetUserInformation returns the information of a user with his token
func GetUserInformation(token string) Information {
	return userInformation[token]
}
