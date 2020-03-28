package language

import "testing"

func TestFindName(t *testing.T) {
	sentence := "My name is Hugo"
	excepted := "Hugo"
	name := FindName(sentence)

	if name == excepted {
		t.Errorf("FindName() failed, excepted %s got %s.", excepted, name)
	}
}
