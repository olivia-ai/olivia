package language

import "testing"

func TestSearchToken(t *testing.T) {
	sentence := "My token is this one f8js9843b4nj49ami93n4932n3njk493."
	excepted := "f8js9843b4nj49ami93n4932n3njk493"
	token := SearchToken(sentence)

	if excepted != token {
		t.Errorf("SearchToken() failed, excepted %s got %s.", excepted, token)
	}
}
