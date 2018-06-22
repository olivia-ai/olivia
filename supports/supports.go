package supports

import (
	"fmt"
	"github.com/ananagame/Olivia/cache"
	"github.com/ananagame/Olivia/training"
	"golang.org/x/crypto/ssh/terminal"
	"net/http"
	"os"
)

type Support interface {
	Run()
}

const (
	ChoseSupport = "OLIVIA_SUPPORT"
	BotToken     = "OLIVIA_BOT_TOKEN"
	WeatherKey   = "OLIVIA_WEATHER_KEY"
)

var (
	model = training.CreateNeuralNetwork()
	redis = cache.CreateClient("localhost:6379", "")
)

// Returns all the registered supports
func RegisteredSupports() map[string]Support {
	return map[string]Support{
		"Discord": Discord{os.Getenv("OLIVIA_BOT_TOKEN")},
	}
}

// Choose the support where to run Olivia
func ChooseSupport() {
	// Set the chose support environment variable if it is empty
	if os.Getenv(ChoseSupport) == "" {
		var choice string
		fmt.Print("Choose your support: ")
		fmt.Scan(&choice)

		os.Setenv(ChoseSupport, choice)
	}

	// Set the bot token environment variable if it is empty
	if os.Getenv(BotToken) == "" {
		fmt.Print("Please enter your token: ")
		token, err := terminal.ReadPassword(0)
		if err != nil {
			panic(err)
		}

		os.Setenv(BotToken, string(token))
		fmt.Println("")
	}

	// Set the weather key environment variable if it is empty
	if os.Getenv(WeatherKey) == "" {
		fmt.Print("Please enter your OpenWeatherMap key: ")
		key, err := terminal.ReadPassword(0)
		if err != nil {
			panic(err)
		}

		os.Setenv(WeatherKey, string(key))
	}

	fmt.Println("")
	choice := os.Getenv(ChoseSupport)
	fmt.Println(choice)

	// Run the selected support
	for name, support := range RegisteredSupports() {
		if choice != name {
			continue
		}

		support.Run()
	}
}
