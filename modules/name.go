package modules

import (
	"fmt"
	"strings"

	"github.com/olivia-ai/olivia/language"
	"github.com/olivia-ai/olivia/user"
	"github.com/olivia-ai/olivia/util"
)

var (
	nameGetterTag = "name getter"
	nameSetterTag = "name setter"
)

func init() {
	RegisterModule(Module{
		Tag: nameGetterTag,
		Patterns: []string{
			"Do you know my name?",
		},
		Responses: []string{
			"Your name is %s!",
		},
		Replacer: NameGetterReplacer,
	})

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

func NameGetterReplacer(_, response, token string) (string, string) {
	name := user.GetUserInformation(token).Name

	if strings.TrimSpace(name) == "" {
		responseTag := "don't know name"
		return responseTag, util.GetMessage(responseTag)
	}

	return nameGetterTag, fmt.Sprintf(response, name)
}

func NameSetterReplacer(entry, response, token string) (string, string) {
	name := language.FindName(entry)

	// If there is no name in the entry string
	if name == "" {
		responseTag := "no name"
		return responseTag, util.GetMessage(responseTag)
	}

	// Capitalize the name
	name = strings.Title(name)

	// Change the name inside the user information
	user.ChangeUserInformation(token, func(information user.Information) user.Information {
		information.Name = name
		return information
	})

	return nameSetterTag, fmt.Sprintf(response, name)
}
