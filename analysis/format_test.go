package analysis

import (
	"reflect"
	"testing"
)

func TestSentence_Arrange(t *testing.T) {
	sentence := Sentence{"Hello. how are you!   "}
	sentence.Arrange()

	excepted := "Hello how are you"

	if sentence.Content != excepted {
		t.Errorf("sentence.Arrange() failed, excepted %v, got %v", excepted, sentence.Content)
	}
}

func TestSentence_Tokenize(t *testing.T) {
	sentence := Sentence{"Hello How are you"}
	tokens := sentence.Tokenize()

	excepted := []string{"hello", "how", "are", "you"}

	if !reflect.DeepEqual(tokens, excepted) {
		t.Errorf("sentence.Tokenize() failed, excepted %v, got %v", excepted, tokens)
	}
}
