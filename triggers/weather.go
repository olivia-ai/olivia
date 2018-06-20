package triggers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type Weather struct {
	Entry    string
	Response string
	Cities   []City
}

type City struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func SerializeCities() (cities []City) {
	bytes, err := ioutil.ReadFile("cities.json")
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(bytes, &cities)

	return cities
}

func SentenceContainsCity(sentence string) (possibilites []City) {
	for _, city := range cities {
		if !strings.Contains(strings.ToLower(sentence)+" ",
			" "+strings.ToLower(city.Name)+" ") {
			continue
		}

		possibilites = append(possibilites, city)
	}

	return possibilites
}

var cities = SerializeCities()

// Replace the content of the sentence by the actual clock
func (weather Weather) ReplaceContent() string {
	// Escape if it isn't a weather message
	if !strings.Contains(weather.Response, "${WEATHER}") {
		return weather.Response
	}

	possibilites := SentenceContainsCity(weather.Entry)

	if len(possibilites) == 0 {
		return "Je n'ai trouvé aucune ville correspondante 😦"
	}

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
