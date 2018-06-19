package supports

import (
	"../cache"
	"../training"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
)

type Support interface {
	Run()
}

var (
	model = training.CreateNeuralNetwork()
	redis = cache.CreateClient("localhost:6379", "")
)

// Returns all the registered supports
func RegisteredSupports(token string) map[string]Support {
	return map[string]Support{
		"Discord": Discord{token},
	}
}

// Choose the support where to run Olivia
func ChooseSupport() {
	var choice string
	fmt.Print("Choose your support: ")
	fmt.Scan(&choice)

	fmt.Printf("Please enter your %s token: ", choice)
	token, err := terminal.ReadPassword(0)

	if err != nil {
		panic(err)
	}

	fmt.Println("")

	// Run the selected support
	for name, support := range RegisteredSupports(string(token)) {
		if choice != name {
			continue
		}

		support.Run()
	}
}
