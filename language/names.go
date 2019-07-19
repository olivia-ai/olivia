package language

import (
	"strings"

	"github.com/olivia-ai/olivia/util"
)

var names = SerializeNames()

func SerializeNames() (names []string) {
	namesFile := string(util.ReadFile("res/names.txt"))

	// Iterate each line of the file
	for _, name := range strings.Split(strings.TrimSuffix(namesFile, "\n"), "\n") {
		names = append(names, name)
	}

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
