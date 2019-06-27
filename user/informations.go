package user

import "github.com/olivia-ai/olivia/util"

type Informations struct {
	Name string `json:"name"`
}

var (
	userInformations         = map[string]Informations{}
	userInformationsUpdates []string
)

// ChangeUserInformations requires the token of the user and a function which gives the actual
// informations and returns the new informations.
func ChangeUserInformations(token string, changer func (Informations) Informations) {
	userInformations[token] = changer(userInformations[token])
	// Add the user's token to userInformationsUpdates to
	userInformationsUpdates = append(userInformationsUpdates, token)
}

// GetUserInformations returns the informations of a user with his token, returns an empty information
// if there were no modifications since.
func GetUserInformations(token string) Informations {
	if !util.Contains(userInformationsUpdates, token) {
		return Informations{}
	}

	userInformationsUpdates = util.Remove(userInformationsUpdates, token)

	return userInformations[token]
}