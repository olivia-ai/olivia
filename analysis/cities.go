package analysis

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type City struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func ReadCities() []byte {
	bytes, err := ioutil.ReadFile("cities.json")
	if err != nil {
		fmt.Println(err)
	}

	return bytes
}

// Unmarshal the json and return the array of Cities
func SerializeCities() []City {
	var cities []City
	json.Unmarshal(ReadCities(), &cities)

	return cities
}