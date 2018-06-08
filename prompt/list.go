package prompt

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type List struct {
	Question string
	Choices []string
	DefaultChoice int
}

// Display the choices
func (list List) Display() {
	// Print the question
	fmt.Printf("\033[36m?\033[0m %s\n", list.Question)

	// Iterate and print choices
	for i, choice := range list.Choices {
		choiceToPrint := "  " + choice
		if list.DefaultChoice == i {
			choiceToPrint = "\033[36m✔\033[0m " + choice
		}

		fmt.Println(choiceToPrint)
	}
}

// Reset the termbox
func (list List) Reset() {
	termbox.Sync()
}

// Run the prompt until you choose or you press echap
func (list List) Run() string {
	// Initialize the termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	// Display a first time the choices
	list.Display()

keyPress:
	for {
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break keyPress
			case termbox.KeyArrowUp:
				list.Reset()
				if list.DefaultChoice != 0 {
					list.DefaultChoice--
				}
				list.Display()
			case termbox.KeyArrowDown:
				list.Reset()
				if list.DefaultChoice != len(list.Choices) - 1 {
					list.DefaultChoice++
				}
				list.Display()
			case termbox.KeyEnter:
				return list.Choices[list.DefaultChoice]
			}
		}
	}

	return "There is a problem"
}