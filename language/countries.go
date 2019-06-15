package language

import (
	"encoding/json"
	"fmt"
	"github.com/olivia-ai/olivia/util"
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

func SerializeCountries() (countries []Country) {
	err := json.Unmarshal(util.ReadFile("res/countries.json"), &countries)
	if err != nil {
		fmt.Println(err)
	}

	return countries
}

// FindCountry returns the country found in the sentence and if no country is found, returns an empty Country struct
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
