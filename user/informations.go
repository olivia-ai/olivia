package user

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