package analysis

import (
	"encoding/json"
	"net/http"

	"github.com/olivia-ai/olivia/locales"

	"github.com/olivia-ai/olivia/modules"
	"github.com/olivia-ai/olivia/util"
)

var (
	defaultModules  []modules.Module
	defaultIntents  []Intent
	defaultMessages []util.Message
)

// Coverage is the coverage for a single language which contains the coverage details of each section
type Coverage struct {
	Modules  CoverageDetails `json:"modules"`
	Intents  CoverageDetails `json:"intents"`
	Messages CoverageDetails `json:"message"`
}

// CoverageDetails are the details of items not covered and the coverage percentage
type CoverageDetails struct {
	NotCovered []string `json:"not_covered"`
	Coverage   int      `json:"coverage"`
}

// GetCoverage encodes the coverage of each language in json
func GetCoverage(writer http.ResponseWriter, request *http.Request) {
	defaultMessages, defaultIntents, defaultModules =
		util.GetMessages("en"), GetIntents("en"), modules.GetModules("en")

	coverage := map[string]Coverage{}

	// Calculate coverage for each language
	for _, locale := range locales.Locales {
		coverage[locale.Tag] = Coverage{
			Modules:  GetModuleCoverage(locale.Tag),
			Intents:  GetIntentCoverage(locale.Tag),
			Messages: GetMessageCoverage(locale.Tag),
		}
	}

	json.NewEncoder(writer).Encode(coverage)
}

// GetMessageCoverage returns an array of not covered messages and the percentage of message that
// aren't translated in the given locale.
func GetMessageCoverage(locale string) CoverageDetails {
	var notCoveredMessages []string

	// Iterate through the default messages which are the english ones to verify if a message isn't
	// translated in the given locale.
	for _, defaultMessage := range defaultMessages {
		message := util.GetMessageByTag(defaultMessage.Tag, locale)

		// Add the current module tag to the list of not-covered-modules
		if message.Tag != defaultMessage.Tag {
			notCoveredMessages = append(notCoveredMessages, defaultMessage.Tag)
		}
	}

	// Calculate the percentage of modules that aren't translated in the given locale
	coverage := CalculateCoverage(len(notCoveredMessages), len(defaultMessages))

	return CoverageDetails{
		NotCovered: notCoveredMessages,
		Coverage:   coverage,
	}
}

// GetIntentCoverage returns an array of not covered intents and the percentage of intents that aren't
// translated in the given locale.
func GetIntentCoverage(locale string) CoverageDetails {
	var notCoveredIntents []string

	// Iterate through the default intents which are the english ones to verify if an intent isn't
	// translated in the given locale.
	for _, defaultIntent := range defaultIntents {
		intent := GetIntentByTag(defaultIntent.Tag, locale)

		// Add the current module tag to the list of not-covered-modules
		if intent.Tag != defaultIntent.Tag {
			notCoveredIntents = append(notCoveredIntents, defaultIntent.Tag)
		}
	}

	// Calculate the percentage of modules that aren't translated in the given locale
	coverage := CalculateCoverage(len(notCoveredIntents), len(defaultModules))

	return CoverageDetails{
		NotCovered: notCoveredIntents,
		Coverage:   coverage,
	}
}

// GetModuleCoverage returns an array of not covered modules and the percentage of modules that aren't
// translated in the given locale.
func GetModuleCoverage(locale string) CoverageDetails {
	var notCoveredModules []string

	// Iterate through the default modules which are the english ones to verify if a module isn't
	// translated in the given locale.
	for _, defaultModule := range defaultModules {
		module := modules.GetModuleByTag(defaultModule.Tag, locale)

		// Add the current module tag to the list of not-covered-modules
		if module.Tag != defaultModule.Tag {
			notCoveredModules = append(notCoveredModules, defaultModule.Tag)
		}
	}

	// Calculate the percentage of modules that aren't translated in the given locale
	coverage := CalculateCoverage(len(notCoveredModules), len(defaultModules))

	return CoverageDetails{
		NotCovered: notCoveredModules,
		Coverage:   coverage,
	}
}

// CalculateCoverage returns the coverage calculated with the given length of not covered
// items and the default items length
func CalculateCoverage(notCoveredLength, defaultLength int) int {
	return 100 * (defaultLength - notCoveredLength) / defaultLength
}
