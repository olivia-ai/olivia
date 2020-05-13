package analysis

import (
	"reflect"
	"testing"
)

func TestSentence_WordsBag(t *testing.T) {
	sentence := Sentence{"en", "Hi how are you"}
	words := Sentence{"en", "hi hello good morning are is were you seven"}.stem()

	wordsBag := sentence.WordsBag(words)
	excepted := []float64{0, 0, 0, 0, 1, 0}

	if !reflect.DeepEqual(excepted, wordsBag) {
		t.Errorf("sentence.WordsBag() failed, excepted %v, got %v", excepted, wordsBag)
	}
}

func TestSentence_Arrange(t *testing.T) {
	sentence := Sentence{"en", "Hello. how are you!   "}
	sentence.arrange()

	excepted := "Hello how are you"

	if sentence.Content != excepted {
		t.Errorf("sentence.Arrange() failed, excepted %v, got %v", excepted, sentence.Content)
	}
}

func TestSentence_Tokenize(t *testing.T) {
	sentence := Sentence{"en", "Hello How are you"}
	tokens := sentence.tokenize()

	excepted := []string{"hello", "how", "are", "you"}

	if !reflect.DeepEqual(tokens, excepted) {
		t.Errorf("sentence.Tokenize() failed, excepted %v, got %v", excepted, tokens)
	}
}
