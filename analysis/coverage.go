package analysis

import (
	"github.com/olivia-ai/olivia/modules"
)

var (
	defaultModules = modules.GetModules("en")
	defaultIntents = GetIntents("en")
)

// GetIntentCoverage returns an array of not covered intents and the percentage of intents that aren't
// translated in the given locale.
func GetIntentCoverage(locale string) (notCoveredIntents []string, coverage int) {
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
	coverage = CalculateCoverage(len(notCoveredIntents), len(defaultModules))

	return
}

// GetModuleCoverage returns an array of not covered modules and the percentage of modules that aren't
// translated in the given locale.
func GetModuleCoverage(locale string) (notCoveredModules []string, coverage int) {
	// Iterate through the default modules which are the english ones to verify if a module isn't
	// translated in the given locale.
	for _, defaultModule := range defaultModules {
		module := GetModuleByTag(defaultModule.Tag, locale)

		// Add the current module tag to the list of not-covered-modules
		if module.Tag != defaultModule.Tag {
			notCoveredModules = append(notCoveredModules, defaultModule.Tag)
		}
	}

	// Calculate the percentage of modules that aren't translated in the given locale
	coverage = CalculateCoverage(len(notCoveredModules), len(defaultModules))

	return
}

// CalculateCoverage returns the coverage calculated with the given length of not covered
// items and the default items length
func CalculateCoverage(notCoveredLength, defaultLength int) int {
	return 100 * (defaultLength - notCoveredLength) / defaultLength
}

// GetModuleByTag returns a module found by the given tag and locale
func GetModuleByTag(tag, locale string) modules.Module {
	for _, module := range modules.GetModules(locale) {
		if tag != module.Tag {
			continue
		}

		return module
	}

	return modules.Module{}
}

// GetIntentByTag returns an intent found by given tag and locale
func GetIntentByTag(tag, locale string) Intent {
	for _, intent := range GetIntents(locale) {
		if tag != intent.Tag {
			continue
		}

		return intent
	}

	return Intent{}
}
