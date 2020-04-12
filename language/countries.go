package language

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/olivia-ai/olivia/util"
)

// Country is the serializer of the countries.json file in the res folder
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

// SerializeCountries returns a list of countries, serialized from `res/datasets/countries.json`
func SerializeCountries() (countries []Country) {
	err := json.Unmarshal(util.ReadFile("res/datasets/countries.json"), &countries)
	if err != nil {
		fmt.Println(err)
	}

	return countries
}

// FindCountry returns the country found in the sentence and if no country is found, returns an empty Country struct
func FindCountry(sentence string) Country {
	for _, country := range countries {
		// If the actual country isn't contained in the sentence, continue
		if !strings.Contains(strings.ToLower(sentence), strings.ToLower(country.CommonName)) &&
			!strings.Contains(strings.ToLower(sentence), strings.ToLower(country.OfficialName)) {
			continue
		}

		// Returns the right country
		return country
	}

	// Returns an empty country if none has been found
	return Country{}
}
