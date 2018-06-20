package triggers

type Trigger interface {
	ReplaceContent() string
}

// Returns all the registered triggers
func RegisteredTriggers(entry, response string) []Trigger {
	return []Trigger{
		Clock{response},
		Weather{entry, response, cities},
	}
}
