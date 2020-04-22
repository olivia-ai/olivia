package locales

import (
	"encoding/json"

	"github.com/olivia-ai/olivia/util"
)

var locales = SerializeLocales()

// A Locale is a registered locale in the file
type Locale struct {
	Locale string `json:"locale"`
	Name   string `json:"name"`
}

// SerializeLocales returns the locales retrieved from the locales JSON file
func SerializeLocales() (locales []Locale) {
	err := json.Unmarshal(util.ReadFile("res/locales/locales.json"), &locales)
	if err != nil {
		panic(err)
	}

	return
}

// GetLocales returns the current locales
func GetLocales() []Locale {
	return locales
}

// GetNameByLocale returns the name of the given locale
func GetNameByLocale(_locale string) string {
	for _, locale := range locales {
		if locale.Locale != _locale {
			continue
		}

		return locale.Name
	}

	return ""
}

func GetLocaleByName(name string) string {
	for _, locale := range locales {
		if locale.Name != name {
			continue
		}

		return locale.Locale
	}

	return ""
}
