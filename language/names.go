package language

import (
	"strings"

	"github.com/olivia-ai/olivia/util"
)

var names = SerializeNames()

// SerializeNames retrieves all the names from res/datasets/names.txt and returns an array of names
func SerializeNames() (names []string) {
	namesFile := string(util.ReadFile("res/datasets/names.txt"))

	// Iterate each line of the file
	names = append(names, strings.Split(strings.TrimSuffix(namesFile, "\n"), "\n")...)
	return
}

// FindName returns a name found in the given sentence or an empty string if no name has been found
func FindName(sentence string) string {
	for _, name := range names {
		if !strings.Contains(strings.ToLower(" "+sentence+" "), " "+name+" ") {
			continue
		}

		return name
	}

	return ""
}
