package triggers

type Trigger interface {
	ReplaceContent() string
}

var (
	Entry    string
	Response string
)

// Returns all the registered triggers
func RegisteredTriggers(entry, response string) []Trigger {
	Entry = entry
	Response = response

	return []Trigger{
		Weather{},
		Random{},
		Capital{},
		Area{},
	}
}

// Apply the triggers
func ReplaceContent(entry, response string) string {
	for _, trigger := range RegisteredTriggers(entry, response) {
		Response = trigger.ReplaceContent()
	}

	return Response
}
