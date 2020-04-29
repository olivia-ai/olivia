package locales

import (
	// Import these packages to trigger the init() function
	_ "github.com/olivia-ai/olivia/res/locales/ca"
	_ "github.com/olivia-ai/olivia/res/locales/en"
	_ "github.com/olivia-ai/olivia/res/locales/es"
	_ "github.com/olivia-ai/olivia/res/locales/fr"
)

// Locales is the list of locales's tags and names
var Locales = []Locale{
	{
		Tag:  "en",
		Name: "english",
	},
	{
		Tag:  "fr",
		Name: "french",
	},
	{
		Tag:  "es",
		Name: "spanish",
	},
	{
		Tag:  "ca",
		Name: "catalan",
	},
}

// A Locale is a registered locale in the file
type Locale struct {
	Tag  string
	Name string
}

// GetNameByTag returns the name of the given locale's tag
func GetNameByTag(tag string) string {
	for _, locale := range Locales {
		if locale.Tag != tag {
			continue
		}

		return locale.Name
	}

	return ""
}

// GetTagByName returns the tag of the given locale's name
func GetTagByName(name string) string {
	for _, locale := range Locales {
		if locale.Name != name {
			continue
		}

		return locale.Tag
	}

	return ""
}

// Exists checks if the given tag exists in the list of locales
func Exists(tag string) bool {
	for _, locale := range Locales {
		if locale.Tag == tag {
			return true
		}
	}

	return false
}
