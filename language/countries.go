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
	Region       string   `json:"continent"`
	SubRegion    string   `json:"subcontinent"`
	Code         string   `json:"code"`
	Borders      []string `json:"borders"`
	Area         float64  `json:"area"`
	Currency     string   `json:"currency"`
	Flag         string   `json:"flag"`
}

var countries = SerializeCountries()

// Serialize the countries.json file
func SerializeCountries() (countries []Country) {
	bytes, err := ioutil.ReadFile("res/countries.json")
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
