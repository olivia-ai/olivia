package modules

import (
	"fmt"
	"github.com/olivia-ai/olivia/language"
	"github.com/olivia-ai/olivia/user"
	"github.com/olivia-ai/olivia/util"
	"strings"
)

var (
	nameGetterTag = "name getter"
	nameSetterTag = "name setter"
)

func init() {
	RegisterModule(Module{
		Tag: nameGetterTag,
		Patterns: []string{
			"What is my name?",
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
	name := user.GetUserInformations(token).Name

	if name == "" {
		responseTag := "don't know name"
		return responseTag, util.GetMessage(responseTag)
	}

	return nameGetterTag, fmt.Sprintf(response, strings.Title(name))
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
