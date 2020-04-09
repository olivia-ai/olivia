package dashboard

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/olivia-ai/olivia/util"

	"golang.org/x/crypto/bcrypt"
)

const file = "res/api/authentication.json"

// An AuthorizedUser is a user allowed to do updates on the API
type AuthorizedUser struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

// A Response is what the REST Api replies
type Response struct {
	Message string `json:"message"`
}

// FirstAuthenticationRequest is the content needed to initialize the authentication file at first
type FirstAuthenticationRequest struct {
	Name string `json:"name"`
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
func InitializeFile(token string) {
	_, err := os.Open(file)
	// Checks if the authentication file exists
	if err == nil {
		return
	}

	// Create the data
	hash, _ := HashToken(token)
	data := []AuthorizedUser{{
		Name:  "Admin",
		Token: hash,
	}}

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
	// Parse the data contained in the authentication file
	var data []AuthorizedUser
	json.Unmarshal(util.ReadFile(file), &data)

	// Iterate the registered tokens to see if the given token is authorized
	for _, user := range data {
		// If the token is valid then returns true
		if CompareHashAndToken(user.Token, token) {
			return true
		}
	}

	return false
}

// HeaderTokenCheck checks if the token is present in the error and returns it with a response
func HeaderTokenCheck(r *http.Request) (string, Response) {
	// Get the token from the headers
	token := r.Header.Get("Olivia-Token")

	if token == "" {
		return "", Response{
			Message: "You need to provide the token in the Headers as 'Olivia-Token: ...'.",
		}
	}

	return token, Response{}
}

// GetAuthentication is the route for getting if the user is authenticated or not
func GetAuthentication(w http.ResponseWriter, r *http.Request) {
	token, response := HeaderTokenCheck(r)
	if response != (Response{}) {
		json.NewEncoder(w).Encode(response)
		return
	}

	InitializeFile(token)

	// Check if the user is authenticated
	isAuthenticated := IsAuthenticated(token)

	// Return the status codes
	if isAuthenticated {
		w.WriteHeader(200) // OK
	} else {
		w.WriteHeader(403) // FORBIDDEN
	}
}
