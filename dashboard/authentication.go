package dashboard

import (
	"crypto/rand"
	"fmt"
	"os"

	"github.com/olivia-ai/olivia/util"

	"github.com/gookit/color"

	"golang.org/x/crypto/bcrypt"
)

var fileName = "res/authentication.txt"

var authenticationHash []byte

// GenerateToken generates a random token of 30 characters and returns it
func GenerateToken() string {
	b := make([]byte, 30)
	rand.Read(b)

	fmt.Println("hey")
	return fmt.Sprintf("%x", b)
}

// HashToken gets the given tokens and returns its hash using bcrypt
func HashToken(token string) []byte {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(token), 14)
	return bytes
}

// ChecksToken checks if the given token is the good one from the authentication file
func ChecksToken(token string) bool {
	err := bcrypt.CompareHashAndPassword(authenticationHash, []byte(token))
	return err == nil
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
		authenticationHash = util.ReadFile(fileName)
		return
	}

	// Generates the token and gives it to the user
	token := GenerateToken()
	fmt.Printf("Your authentication token is: %s\n", color.FgLightGreen.Render(token))
	fmt.Println("Save it, you won't be able to get it again unless you generate a new one.")
	fmt.Println()

	// Hash the token and save it
	hash := HashToken(token)
	SaveHash(string(hash))

	authenticationHash = hash
}
