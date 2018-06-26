package language

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"strconv"
)

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

// Returns a list of possible cities contains in a sentence
func FindCities(sentence string) (possibilites []City) {
	sentence = strings.ToLower(sentence) + " "

	for _, city := range SerializeCities() {
		if !strings.Contains(sentence, " "+strings.ToLower(city.Name)+" ") {
			continue
		}

		possibilites = append(possibilites, city)
	}

	for _, possibility := range possibilites {
		cityName := " " + strings.ToLower(possibility.Name) + " "
		cityNumber := regexp.MustCompile(cityName + "\\d+")

		if cityNumber.MatchString(sentence) {
			match := strings.Replace(cityNumber.FindString(sentence), cityName, "", 1)
			number, _ := strconv.Atoi(match)

			return []City{possibilites[number - 1]}
		}
	}


	return possibilites
}
