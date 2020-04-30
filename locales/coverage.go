package locales

import (
	"github.com/olivia-ai/olivia/modules"
)

var defaultModules = modules.GetModules("en")

// GetCoverage returns an array of not covered modules and the percentage of modules that aren't
// translated in the given locale.
func GetCoverage(locale string) (notCoveredModules []string, coverage int) {
	// Iterate through the default modules which are the english ones to verify if a module isn't
	// translated in the given locale.
	for _, englishModule := range defaultModules {
		module := GetModuleByTag(englishModule.Tag, locale)

		// Add the current module tag to the list of not-covered-modules
		if module.Tag != englishModule.Tag {
			notCoveredModules = append(notCoveredModules, englishModule.Tag)
		}
	}

	// Calculate the percentage of modules that aren't translated in the given locale
	defaultLength := len(defaultModules)
	coverage = 100 * (defaultLength - len(notCoveredModules)) / defaultLength

	return
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
