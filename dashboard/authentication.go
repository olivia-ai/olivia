package dashboard

import (
	"encoding/json"
	"os"

	"github.com/olivia-ai/olivia/util"
	"golang.org/x/crypto/bcrypt"
)

// An AuthorizedUser is a user allowed to do updates on the API
type AuthorizedUser struct {
	Token string `json:"token"`
}

// HashToken returns the hash of the token using bcrypt
func HashToken(token string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	return string(bytes), err
}

// CompareHash compares a hash with a token and returns the boolean
func CompareHashAndToken(hash, token string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(token))
	return err == nil
}

// InitializeFile creates the default authentication file with a first given token
func InitializeFile(file, token string) {
	// Create the data
	hash, _ := HashToken(token)
	data := []AuthorizedUser{
		{Token: hash},
	}

	// Create the file
	outF, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic("Failed to create " + file + ".")
	}
	defer outF.Close()

	// Encode the json data
	encoder := json.NewEncoder(outF)
	err = encoder.Encode(data)
	if err != nil {
		panic(err)
	}
}

// IsAuthenticated checks if the given token is contained in the authentication file,
// if not it creates a file with the given hashed token inside.
func IsAuthenticated(token string) bool {
	file := "../res/api/authentication.json"

	_, err := os.Open(file)
	// Checks if the authentication file exists
	if err == nil {
		// Parse the data contained in the authentication file
		var data []AuthorizedUser
		json.Unmarshal(util.ReadFile(file), &data)

		// Iterate the registered tokens to see if the given token is authorized
		for _, o := range data {
			// If the token is valid then returns true
			if CompareHashAndToken(o.Token, token) {
				return true
			}
		}

		return false
	}

	// Creates the authentication file
	InitializeFile(file, token)
	return true
}
