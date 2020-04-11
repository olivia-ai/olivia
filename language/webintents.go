package language

import (
	"regexp"
	"strings"
)

// SearchPatterns looks up for patterns in messages for the WebIntents like:
// “{posts[0]}” or “{posts.name}”
func SearchPatterns(content string) {
	patternRegex := regexp.MustCompile(`{([a-zA-Z_-]+(\[\d+\])?(\.)?)+}`)
	patterns := patternRegex.FindAllString(content, -1)

	// Iterate through the patterns to parse them
	for _, pattern := range patterns {
		// Remove braces from the pattern
		pattern = strings.Replace(
			strings.Replace(pattern, "{", "", 1),
			"}", "", -1,
		)
	}
}
