package language

import "regexp"

type Word struct {
	Content string
}

// Stem the content of the word with a list of suffixes
func (word Word) Stem() string {
	suffixes := regexp.MustCompile("(" +
		"lement|sement|nerait|eaient|amment|aient|ement|eront|erait|ation|" +
		"erait|euse|eant|ante|ance|ions|erai|eait|les|ait|eux|nez|ive|ies|" +
		"ité|ant|ant|era|ie|ir|le|és|ée|er|ez|e|é|i|s)$")

	return suffixes.ReplaceAllString(word.Content, "")
}
