package triggers

type Trigger interface {
	ReplaceContent() string
}

// Returns all the registered triggers
func RegisteredTriggers(sentence string) []Trigger {
	return []Trigger{
		Clock{sentence},
	}
}
