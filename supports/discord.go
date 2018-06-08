package supports

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

type Discord struct {
	Token string
}

func (discord Discord) Run() {
	// Create a new Discord session using the support token.
	dg, err := discordgo.New("Bot " + discord.Token)
	if err != nil {
		fmt.Println("Error creating the discord bot session: ", err)
		return
	}

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
