package supports

import (
	"../analysis"
	"fmt"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

type Telegram struct {
	Token string
}

func (telegram Telegram) Run() {
	// Create a new telegram bot using the support token
	b, err := telebot.NewBot(telebot.Settings{
		Token:  telegram.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	// Handle the sended message
	b.Handle(telebot.OnText, func(m *telebot.Message) {
		// Get the response with the chatbot
		response := analysis.Sentence{
			Content: m.Text,
		}.Response(model, string(m.Sender.ID))

		// Respond it
		b.Send(m.Sender, response)
	})

	fmt.Println("Telegram support started")
	b.Start()
}
