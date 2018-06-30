package triggers

import (
	"github.com/olivia-ai/Api/data"
	"github.com/olivia-ai/Api/language"
	"strconv"
	"strings"
)

type Area struct{}

func (area Area) ReplaceContent() string {
	// Escape if it isn't a message which contains a Country
	if !strings.Contains(Response, "${AREA}") {
		return Response
	}

	country := language.FindCountry(Entry)

	// If there isn't a country respond with a data.Message
	if country.Code == "" {
		return data.GetMessage("no country")
	}

	response := strings.Replace(Response, "${AREA}", strconv.Itoa(country.Area), 1)
	response = strings.Replace(response, "${COUNTRY}", country.OfficialName, 1)

	return response
}
