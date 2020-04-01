package language

import (
	"reflect"
	"testing"
)

func TestSearchToken(t *testing.T) {
	sentence := "My token is this one f8js9843b4nj49ami93n4932n3njk493 and f8js9843b4nj49ami93n4932n3njk492."
	excepted := []string{"f8js9843b4nj49ami93n4932n3njk493", "f8js9843b4nj49ami93n4932n3njk492"}
	tokens := SearchTokens(sentence)

	if !reflect.DeepEqual(excepted, tokens) {
		t.Errorf("SearchToken() failed, excepted %s got %s.", excepted, tokens)
	}
}
