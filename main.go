package main

import (
	"github.com/ananagame/Olivia/supports"
	"os"
)

func main() {
	os.Setenv("OLIVIA_SUPPORT", "Discord")
	os.Setenv("OLIVIA_WEATHER_KEY", "454b888cdf40e6ed808676b5a6be9783")
	os.Setenv("OLIVIA_BOT_TOKEN", "NDUxMDg3NDU3NzQwNjUyNTU0.DfHeTA.NdqdZlDcZlE8ZNg9lh15irFXpOg")
	supports.ChooseSupport()
}
