package triggers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
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

// Serialize the json file which contains cities as an array
func SerializeCities() (cities []City) {
	bytes, err := ioutil.ReadFile("cities.json")
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(bytes, &cities)

	return cities
}

// Returns the list of cities found in the sentence
func SentenceCities(sentence string) (possibilites []City) {
	for _, city := range cities {
		if !strings.Contains(strings.ToLower(sentence)+" ",
			" "+strings.ToLower(city.Name)+" ") {
			continue
		}

		possibilites = append(possibilites, city)
	}

	return possibilites
}

// Returns an array of numbers found in the sentence
func ScanNumbers(sentence string) []string {
	regexp, _ := regexp.Compile("[0-9]+")
	return regexp.FindAllString(sentence, -1)
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

var cities = SerializeCities()

// Replace the content of the sentence by the actual clock
func (weather Weather) ReplaceContent() string {
	// Escape if it isn't a weather message
	if !strings.Contains(weather.Response, "${WEATHER}") {
		return weather.Response
	}

	possibilites := SentenceCities(weather.Entry)

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

	numbers := ScanNumbers(weather.Entry)

	if len(numbers) == 0 {
		response := "J'ai trouvé plusieurs villes :\n"

		for i, city := range possibilites {
			response += fmt.Sprintf("%d - %s, %s\n", i+1, city.Name, city.Country)
		}

		return response
	}

	lastNumber, _ := strconv.Atoi(numbers[len(numbers)-1])
	GetWeather(possibilites[lastNumber].Id)
	return strings.Replace(
		weather.Response,
		"${WEATHER}",
		"météo machin à "+possibilites[lastNumber].Name,
		1)
}
