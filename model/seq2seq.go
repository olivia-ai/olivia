package model

import "reflect"

// Seq2Seq contains the stucture of a sequence to sequence model including
// the encoder and decoder recurrent neural networks.
type Seq2Seq struct {
	Encoder NN
	Decoder NN
	VocabularySize int
}

// CreateSeq2Seq creates and returns a new sequence to sequence (seq2seq) model.
func CreateSeq2Seq(vocabularySize int, learningRate float64, hiddenLayersNodes ...int) Seq2Seq {
	// Use twice the size of the embedding size for the encoder because we need to take the input
	// embedding and the previous one as input.
	nn := CreateNeuralNetwork(learningRate, 2 * vocabularySize, vocabularySize, hiddenLayersNodes...)

	return Seq2Seq{
		Encoder: nn,
		Decoder: nn,
		VocabularySize: vocabularySize,
	}
}

// FeedForward processes the forward propagation over the encoder and the decoder 
// of the sequence to sequence model
func (s2s *Seq2Seq) FeedForward(embeddings matrix) matrix {
	hiddenStates := matrix{
		make([]float64, len(embeddings[0])),
	}
	for i, embedding := range embeddings {
		input := append(embedding, hiddenStates[i]...)
		hidden := s2s.Encoder.FeedForward(input)
		hiddenStates = append(hiddenStates, hidden[0])
	}

	decoderHiddenStates := matrix{
		hiddenStates[len(hiddenStates)-1],
	}
	output := matrix{
		// START token
		make([]float64, len(embeddings[0])),
	}

	// END token
	for i := 0 ;reflect.DeepEqual(output[len(output)-1], make([]float64, len(embeddings[0]))); i++ {
		input := append(output[len(output)-1], decoderHiddenStates[i]...)
		output = append(output, s2s.Decoder.FeedForward(input)[0])
	}

	return output[1:]
}