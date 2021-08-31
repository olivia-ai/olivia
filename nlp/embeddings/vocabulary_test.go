package embeddings

import (
	"reflect"
	"testing"

	"github.com/olivia-ai/olivia/data"
)

var (
	conversations = []data.Conversation{
		{"Hello you", "Hi"},
		{"How are you ?", "Good"},
	}
	vocabulary = []string{
		"hi", "hello", "you", "good", "how", "are", "?",
	}
)

func TestEstablishVocabulary(t *testing.T) {
	calculatedVocabulary := EstablishVocabulary(conversations)

	if !reflect.DeepEqual(vocabulary, calculatedVocabulary) {
		t.Errorf("EstablishVocabulary() failed.")
	}
}

func TestGetEmbedding(t *testing.T) {
	e := getEmbedding(vocabulary, "hello")
	if e[3] != 1 || e[2] != 0 {
		t.Errorf("getEmbedding() failed.")
	}
}