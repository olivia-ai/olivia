package language

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type Country struct {
	OfficialName string   `json:"official_name"`
	CommonName   string   `json:"common_name"`
	Capital      string   `json:"capital"`
	Continent    string   `json:"continent"`
	SubContinent string   `json:"subcontinent"`
	Code         string   `json:"code"`
	Borders      []string `json:"borders"`
	Area         int      `json:"area"`
	Population   int      `json:"population"`
}

var countries = SerializeCountries()

// Serialize the countries.json file
func SerializeCountries() (countries []Country) {
	bytes, err := ioutil.ReadFile("countries.json")
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(bytes, &countries)

	return countries
}

// Returns the country found in the sentence
func FindCountry(sentence string) Country {
	for _, country := range countries {
		if !strings.Contains(strings.ToLower(sentence), strings.ToLower(country.CommonName)) &&
			!strings.Contains(strings.ToLower(sentence), strings.ToLower(country.OfficialName)) {
			continue
		}

		return country
	}

	return Country{}
}
