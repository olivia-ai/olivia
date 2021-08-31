package model

import (
	"fmt"
	"strings"
	"testing"

	"github.com/olivia-ai/olivia/data"
	"github.com/olivia-ai/olivia/nlp/embeddings"
)

func TestS2SFeedForward(t *testing.T) {

	s := strings.Fields("how are you?")
	c := data.ReadCSVConversationalDataset("data/mock.csv")
	voc := embeddings.EstablishVocabulary(c)
	var input matrix
	for _, word := range s {
		input = append(
			input, 
			embeddings.GetLevenshteinEmbedding(voc, word),
		)
	}

	model := CreateSeq2Seq(len(voc) + 2, 0.25)
	fmt.Println(model.Decoder.Layers)
	fmt.Println(model.FeedForward(input))
}