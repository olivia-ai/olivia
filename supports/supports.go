package supports

import (
	"../prompt"
	"fmt"
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
	var token string
	fmt.Printf("Please enter your %s token: ", choice)
	fmt.Scanf("%s", &token)

	// Run the selected choice
	for name, support := range RegisteredSupports(token) {
		if choice != name {
			continue
		}

		support.Run()
	}
}
