package language

import (
	"testing"
)

func TestSerializeNames(t *testing.T) {
	names := SerializeNames()
	excepted := "a'isha"

	if names[1] != excepted {
		t.Errorf("SerializeNames() failed, excepted %s got %s.", excepted, names[0])
	}
}

func TestFindName(t *testing.T) {
	sentence := "My name is Hugo"
	excepted := "Hugo"
	name := FindName(sentence)

	if name == excepted {
		t.Errorf("FindName() failed, excepted %s got %s.", excepted, name)
	}
}
