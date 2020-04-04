package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/olivia-ai/olivia/modules/start"

	"github.com/gookit/color"
	"github.com/gorilla/websocket"
	"github.com/olivia-ai/olivia/analysis"
	"github.com/olivia-ai/olivia/user"
	"github.com/olivia-ai/olivia/util"
)

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// RequestMessage is the structure that uses entry connections to chat with the websocket
type RequestMessage struct {
	Type        int              `json:"type"` // 0 for handshakes and 1 for messages
	Content     string           `json:"content"`
	Token       string           `json:"user_token"`
	Information user.Information `json:"information"`
}

// ResponseMessage is the structure used to reply to the user through the websocket
type ResponseMessage struct {
	Content     string           `json:"content"`
	Tag         string           `json:"tag"`
	Information user.Information `json:"information"`
}

// SocketHandle manages the entry connections and reply with the neural network
func SocketHandle(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	fmt.Println(color.FgGreen.Render("A new connection has been opened"))

	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// Unmarshal the json content of the message
		var request RequestMessage
		if err = json.Unmarshal(msg, &request); err != nil {
			continue
		}

		// Set the information from the client into the cache
		if reflect.DeepEqual(user.GetUserInformation(request.Token), user.Information{}) {
			user.SetUserInformation(request.Token, request.Information)
		}

		// If the type of requests is a handshake then execute the start modules
		if request.Type == 0 {
			start.ExecuteModules(request.Token)

			message := start.GetMessage()
			if message == "" {
				message = "What can I do for you?"
			}

			// Generate the response to send to the user
			response := ResponseMessage{
				Content:     message,
				Tag:         "start module",
				Information: user.GetUserInformation(request.Token),
			}

			bytes, err := json.Marshal(response)
			if err != nil {
				panic(err)
			}

			if err = conn.WriteMessage(msgType, bytes); err != nil {
				continue
			}

			continue
		}

		// Write message back to browser
		response := Reply(request)
		if err = conn.WriteMessage(msgType, response); err != nil {
			continue
		}
	}
}

// Reply takes the entry message and returns an array of bytes for the answer
func Reply(request RequestMessage) []byte {
	var responseSentence, responseTag string

	// Send a message from res/messages.json if it is too long
	if len(request.Content) > 500 {
		responseTag = "too long"
		responseSentence = util.GetMessage(responseTag)
	} else {
		responseTag, responseSentence = analysis.NewSentence(
			request.Content,
		).Calculate(*cache, neuralNetwork, intentsPath, request.Token)
	}

	// Marshall the response in json
	response := ResponseMessage{
		Content:     responseSentence,
		Tag:         responseTag,
		Information: user.GetUserInformation(request.Token),
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	return bytes
}
