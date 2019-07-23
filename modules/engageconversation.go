// This is not a module, it is a function that generates the first message(s) to send to the user
// when he opens the connection with the websocket.

package modules

import (
	"fmt"
	"strings"

	"github.com/olivia-ai/olivia/user"
)

func EngageConversation(userToken string) (responses []string) {
	userInformations := user.GetUserInformations(userToken)

	// Add the hello message
	if strings.TrimSpace(userInformations.Name) == "" {
		responses = append(
			responses,
			"Hello!",
			// Indicate that the user can tell his name
			"You can tell me your name by simply telling 'My name is John'",
		)
	} else {
		responses = append(responses, fmt.Sprintf("Hey back, %s!", userInformations.Name))
	}

	fmt.Println(responses)

	return
}
