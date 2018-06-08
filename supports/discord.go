package supports

import (
	"../analysis"
	"../training"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Discord struct {
	Token string
}

var model = training.CreateModel()

func (discord Discord) Run() {
	// Create a new Discord session using the support token.
	dg, err := discordgo.New("Bot " + discord.Token)
	if err != nil {
		fmt.Println("Error creating the discord bot session: ", err)
		return
	}

	// Register the event
	dg.AddHandler(messageCreate)

	// Open the connection
	err = dg.Open()
	if err != nil {
		fmt.Println("Rrror opening connection: ", err)
		return
	}

	// Wait here until CTRL-C
	fmt.Println("Discord support is running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// Retrieves the user entry when the discord bot is mentionned and respond with the chatbot
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Discord format for bot mention
	botMention := fmt.Sprintf("<@%s>", s.State.User.ID)

	// If the message is "ping" reply with "Pong!"
	if strings.HasPrefix(m.Content, botMention) {
		// Get the response with the chatbot
		response := analysis.Sentence{
			Content: strings.Replace(m.Content, botMention, "", 1),
		}.Response(model, m.Author.ID)

		// Respond it with a user mention
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%s> %s", m.Author.ID, response))
	}
}
