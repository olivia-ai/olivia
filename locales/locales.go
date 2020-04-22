package locales

import (
	// Import these packages to trigger the init() function
	_ "github.com/olivia-ai/olivia/res/locales/en"
	_ "github.com/olivia-ai/olivia/res/locales/fr"
)

var Locales = []Locale{
	{
		Tag:  "en",
		Name: "english",
	},
	{
		Tag:  "fr",
		Name: "french",
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
