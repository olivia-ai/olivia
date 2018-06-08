package main

import (
	"./supports"
	"flag"
)

var (
	Support string
	Token   string
)

// Retrieves the parameters from command line
func init() {
	flag.StringVar(&Support, "support", "", "The support")
	flag.StringVar(&Token, "token", "", "The bot token")
	flag.Parse()
}

func main() {
	if Support == "" || Token == "" {
		supports.ChooseSupport()
	} else {
		// Run the selected choice
		for name, support := range supports.RegisteredSupports(Token) {
			if Support != name {
				continue
			}

			support.Run()
		}
	}
}
