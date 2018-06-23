package triggers

import (
	"../language"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"encoding/json"
)

type Weather struct {
	Entry    string
	Response string
	Cities   []language.City
}

type Response struct {
	Coord      Coord             `json:"coord"`
	Weather    []WeatherResponse `json:"weather"`
	Base       string            `json:"base"`
	Main       Main              `json:"main"`
	Visibility int               `json:"visibility"`
	Wind       Wind              `json:"wind"`
	Clouds     Cloud             `json:"clouds"`
	Dt         int               `json:"dt"`
	Sys        Sys               `json:"sys"`
	Id         int               `json:"id"`
	Name       string            `json:"name"`
	Cod        int               `json:"cod"`
}

type Coord struct {
	Longitude int `json:"lon"`
	Latitude  int `json:"lat"`
}

type WeatherResponse struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temperature    int `json:"temp"`
	Pressure       int `json:"pressure"`
	Humidity       int `json:"humidity"`
	TemperatureMin int `json:"temp_min"`
	TemperatureMax int `json:"temp_max"`
}

type Wind struct {
	Speed int `json:"speed"`
	Deg   int `json:"deg"`
}

type Cloud struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int    `json:"type"`
	Id      int    `json:"id"`
	Message int    `json:"message"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

func GetWeather(cityId int) Response {
	apiUrl := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?id=%s&APPID=%s",
		strconv.Itoa(cityId),
		os.Getenv("OLIVIA_WEATHER_KEY"))

	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var weather Response
	json.Unmarshal(body, &weather)

	return weather
}

var cities = language.SerializeCities()

// Replace the content of the sentence by the actual clock
func (weather Weather) ReplaceContent() string {
	// Escape if it isn't a weather message
	if !strings.Contains(weather.Response, "${WEATHER}") {
		return weather.Response
	}

	possibilites := language.FindCities(weather.Entry)

	// No cities found in this sentence
	if len(possibilites) == 0 {
		return "Je n'ai trouvé aucune ville correspondante 😦"
	}

	// Respond weather with the good city
	if len(possibilites) == 1 {
		GetWeather(possibilites[0].Id)
		return strings.Replace(
			weather.Response,
			"${WEATHER}",
			"météo machin à "+possibilites[0].Name,
			1)
	}

	response := "J'ai trouvé plusieurs villes :\n"
	for i, city := range possibilites {
		response += fmt.Sprintf("%d - %s, %s\n", i+1, city.Name, city.Country)
	}

	return response
}
