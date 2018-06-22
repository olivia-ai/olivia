package triggers

import (
	"../language"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Weather struct {
	Entry    string
	Response string
	Cities   []language.City
}

func GetWeather(cityId int) string {
	apiUrl := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?id=%s&APPID=%s",
		strconv.Itoa(cityId),
		os.Getenv("OLIVIA_WEATHER_KEY"))

	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err)
		return ":("
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return string(body)
}

var cities = language.SerializeCities()

// Replace the content of the sentence by the actual clock
func (weather Weather) ReplaceContent() string {
	// Escape if it isn't a weather message
	if !strings.Contains(weather.Response, "${WEATHER}") {
		return weather.Response
	}

	possibilites := language.FindCities(weather.Entry)

	// No cities found in this sentence
	if len(possibilites) == 0 {
		return "Je n'ai trouvé aucune ville correspondante 😦"
	}

	// Respond weather with the good city
	if len(possibilites) == 1 {
		return strings.Replace(
			weather.Response,
			"${WEATHER}",
			"météo machin à "+possibilites[0].Name,
			1)
	}

	response := "J'ai trouvé plusieurs villes :\n"
	for i, city := range possibilites {
		response += fmt.Sprintf("%d - %s, %s\n", i+1, city.Name, city.Country)
	}

	return response
}
