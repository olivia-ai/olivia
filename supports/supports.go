package supports

import (
	"../prompt"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
)

type Support interface {
	Run()
}

// Returns all the registered supports
func RegisteredSupports(token string) map[string]Support {
	return map[string]Support{
		"Discord": Discord{token},
	}
}

// Choose which support the user wants to use
func ChooseSupport() {
	var choices []string
	// Iterate with null token for getting names
	for name := range RegisteredSupports("") {
		choices = append(choices, name)
	}

	// Select the support
	choice := prompt.List{
		Question:      "Which support would you like to use?",
		Choices:       choices,
		DefaultChoice: 0,
	}.Run()

	// Get the token
	fmt.Printf("Please enter your %s token: ", choice)
	token, err := terminal.ReadPassword(0)
	if err != nil {
		panic(err)
	}

	fmt.Println("")

	// Run the selected choice
	for name, support := range RegisteredSupports(string(token)) {
		if choice != name {
			continue
		}

		support.Run()
	}
}
