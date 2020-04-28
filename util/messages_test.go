package util

import (
	"testing"
)

func TestSerializeMessages(t *testing.T) {
	messages := SerializeMessages("en")

	if len(messages) == 0 {
		t.Errorf("SerializeMessages() failed.")
	}
}
