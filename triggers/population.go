package triggers

import (
	"github.com/olivia-ai/Api/data"
	"github.com/olivia-ai/Api/language"
	"strconv"
	"strings"
)

type Population struct{}

func (population Population) ReplaceContent() string {
	// Escape if it isn't a message which contains a Country
	if !strings.Contains(Response, "${POPULATION}") {
		return Response
	}

	country := language.FindCountry(Entry)

	// If there isn't a country respond with a data.Message
	if country.Code == "" {
		return data.GetMessage("no country")
	}

	if country.Population == 0 {
		return strings.Replace(data.GetMessage("no population"), "${COUNTRY}", country.OfficialName, 1)
	}

	response := strings.Replace(Response, "${POPULATION}", strconv.Itoa(country.Population), 1)
	response = strings.Replace(response, "${COUNTRY}", country.OfficialName, 1)

	return response
}
