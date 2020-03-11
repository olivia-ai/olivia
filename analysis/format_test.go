package analysis

import (
	"reflect"
	"testing"
)

func TestSentence_WordsBag(t *testing.T) {
	sentence := Sentence{"Hi how are you"}
	words := []string{"hi", "hello", "good", "morning", "are", "is", "were", "you", "seven"}

	wordsBag := sentence.WordsBag(words)
	excepted := []float64{1, 0, 0, 0, 1, 0, 0, 1, 0}

	if !reflect.DeepEqual(excepted, wordsBag) {
		t.Errorf("sentence.WordsBag() failed, excepted %v, got %v", excepted, wordsBag)
	}
}

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
