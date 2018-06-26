package triggers

import (
	"../language"
	"../data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Weather struct {}

type RequestResponse struct {
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
	Temperature    float64 `json:"temp"`
	Pressure       int     `json:"pressure"`
	Humidity       int     `json:"humidity"`
	TemperatureMin float64 `json:"temp_min"`
	TemperatureMax float64 `json:"temp_max"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
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

// Send a HTTP request to openweathermap
func GetWeather(cityId int) RequestResponse {
	apiUrl := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?id=%s&APPID=%s&units=metric&lang=fr",
		strconv.Itoa(cityId),
		os.Getenv("WEATHER_KEY"))

	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var weather RequestResponse
	json.Unmarshal(body, &weather)

	return weather
}

var cities = language.SerializeCities()

// Replace the content of the sentence by the actual clock
func (weather Weather) ReplaceContent() string {
	// Escape if it isn't a weather message
	if !strings.Contains(Response, "${WEATHER}") {
		return Response
	}

	possibilites := language.FindCities(Entry)

	// No cities found in this sentence
	if len(possibilites) == 0 {
		return data.GetMessage("no city")
	}

	// Respond weather with the good city
	if len(possibilites) == 1 {
		conditions := GetWeather(possibilites[0].Id)

		return strings.Replace(
			Response,
			"${WEATHER}",
			fmt.Sprintf("%s avec %d°C", conditions.Weather[0].Description, int(conditions.Main.Temperature)),
			1)
	}

	response := data.GetMessage("cities")
	for i, city := range possibilites {
		response += fmt.Sprintf("%d - %s, %s. ", i+1, city.Name, city.Country)
	}

	return response
}
