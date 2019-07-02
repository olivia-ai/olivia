package modules

import (
	"fmt"
	"github.com/olivia-ai/olivia/user"
	"strings"
)

var nameGetterTag = "name getter"

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
}

func NameGetterReplacer(_, response, token string) (string, string) {
	name := user.GetUserInformations(token).Name
	return nameGetterTag, fmt.Sprintf(response, strings.Title(name))
}
