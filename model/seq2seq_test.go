package model

import (
	"testing"

	"github.com/olivia-ai/olivia/data"
	"github.com/olivia-ai/olivia/nlp/embeddings"
	"github.com/schollz/progressbar/v3"
)

func TestS2SFeedForward(t *testing.T) {
	c := data.ReadCSVConversationalDataset("data/mock.csv")
	voc := embeddings.EstablishVocabulary(c)
	
	model := CreateSeq2Seq(len(voc) + 2, 0.01)

	var input []matrix
	var output []matrix
	for _, conversation := range c {
		input = append(input, embeddings.GetLevenshteinEmbeddings(voc, conversation.Question))
		output = append(
			output, 
			append(embeddings.GetLevenshteinEmbeddings(voc, conversation.Answer), model.EOS),
		)
	}

	bar := progressbar.Default(100, "training")
	for epochs := 0; epochs < 100; epochs++ {
		for i := 0; i < len(input); i++ {
			calculatedOutput := model.FeedForwardWhileTraining(input[i], len(output[i]))
			model.PropagateBackward(calculatedOutput, output[i])
		}
		bar.Add(1)
	}
}