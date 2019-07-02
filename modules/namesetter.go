package modules

import (
	"fmt"
	"github.com/olivia-ai/olivia/language"
	"github.com/olivia-ai/olivia/user"
	"strings"
)

var nameSetterTag = "name setter"

func init() {
	RegisterModule(Module{
		Tag: nameSetterTag,
		Patterns: []string{
			"My name is ",
			"You can call me ",
		},
		Responses: []string{
			"Great! Hi %s",
		},
		Replacer: NameSetterReplacer,
	})
}

func NameSetterReplacer(entry, response, token string) (string, string) {
	name := language.FindName(entry)
	// Change the name inside the user information
	user.ChangeUserInformations(token, func(information user.Information) user.Information {
		information.Name = name
		return information
	})

	return nameSetterTag, fmt.Sprintf(response, strings.Title(name))
}
