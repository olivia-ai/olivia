package language

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
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

	numberRegex := regexp.MustCompile("\\d+")
	numbers := numberRegex.FindAllString(sentence, -1)

	for _, city := range SerializeCities() {
		if !strings.Contains(sentence, " "+strings.ToLower(city.Name)+" ") {
			continue
		}

		possibilites = append(possibilites, city)
	}

	// Returns the sentence's number city
	if len(numbers) != 0 {
		number, _ := strconv.Atoi(numbers[len(numbers)-1])

		return []City{possibilites[number+1]}
	}

	return possibilites
}
