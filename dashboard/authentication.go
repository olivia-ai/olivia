package dashboard

import (
	"crypto/rand"
	"fmt"
	"os"

	"github.com/gookit/color"

	"golang.org/x/crypto/bcrypt"
)

const fileName = "res/authentication.txt"

// GenerateToken generates a random token of 30 characters and returns it
func GenerateToken() string {
	b := make([]byte, 30)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// HashToken gets the given tokens and returns its hash using bcrypt
func HashToken(token string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(token), 14)
	return string(bytes)
}

// AuthenticationFileExists checks if the authentication file exists and return the condition
func AuthenticationFileExists() bool {
	_, err := os.Open(fileName)
	return err == nil
}

// SaveHash saves the given hash to the authentication file
func SaveHash(hash string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.WriteString(hash)
}

// Authenticate checks if the authentication file exists and if not it generates the file with a new token
func Authenticate() {
	// Do nothing if the authentication file exists
	if AuthenticationFileExists() {
		return
	}

	// Generates the token and gives it to the user
	token := GenerateToken()
	fmt.Printf("Your authentication token is: %s\n", color.FgLightGreen.Render(token))
	fmt.Println("Save it, you won't be able to get it again unless you generate a new one.")
	fmt.Println()

	// Hash the token and save it
	hash := HashToken(token)
	SaveHash(hash)
}
