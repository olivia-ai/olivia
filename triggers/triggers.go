package triggers

type Trigger interface {
	ReplaceContent() string
}

// Returns all the registered triggers
func RegisteredTriggers(entry, response string) []Trigger {
	return []Trigger{
		Weather{
			Entry:    entry,
			Response: response,
		},
		Random{
			Response: response,
		},
		Capital{
			Entry: entry,
			Response: response,
		},
	}
}

// Apply the triggers
func ReplaceContent(entry, response string) string {
	for _, trigger := range RegisteredTriggers(entry, response) {
		response = trigger.ReplaceContent()
	}

	return response
}
