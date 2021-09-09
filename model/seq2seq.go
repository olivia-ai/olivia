package model

import (
	"reflect"

	nlpe "github.com/olivia-ai/olivia/nlp/embeddings"
)

// Seq2Seq contains the stucture of a sequence to sequence model including
// the encoder and decoder recurrent neural networks.
type Seq2Seq struct {
	Encoder NN
	Decoder NN
	VocabularySize int
	EOS []float64
	BOS []float64
}

// CreateSeq2Seq creates and returns a new sequence to sequence (seq2seq) model.
func CreateSeq2Seq(vocabularySize int, learningRate float64, hiddenLayersNodes ...int) Seq2Seq {
	return Seq2Seq{
		// Use twice the size of the embedding size for the encoder because we need to take the input
		// embedding and the previous one as input.
		Encoder: CreateNeuralNetwork(learningRate, 2 * vocabularySize, vocabularySize, hiddenLayersNodes...),
		Decoder: CreateNeuralNetwork(learningRate, 2 * vocabularySize, 2 * vocabularySize, hiddenLayersNodes...),

		// Initialize helpers to access data on the vocabulary
		VocabularySize: vocabularySize,
		EOS: nlpe.GetEOS(vocabularySize),
		BOS: nlpe.GetBOS(vocabularySize),
	}
}

// FeedForward processes the forward propagation over the encoder and the decoder 
// of the sequence to sequence model
func (s2s *Seq2Seq) FeedForward(embeddings matrix) matrix {
	hiddenStates := matrix{
		// Initialize the hidden states with an empty embedding
		make([]float64, len(embeddings[0])),
	}
	for i, embedding := range embeddings {
		// Concatenate the input with the given embedding and the previous hidden state
		input := append(embedding, hiddenStates[i]...)
		hidden := s2s.Encoder.FeedForward(input)
		hiddenStates = append(hiddenStates, hidden[0])
	}

	decoderHiddenStates := matrix{
		// Begin with the last hidden state of the encoder
		hiddenStates[len(hiddenStates)-1],
	}
	output := matrix{s2s.BOS}

	for i := 0; !reflect.DeepEqual(s2s.EOS, output[len(output)-1]); i++ {
		// Concatenate the previous output with the current hidden state
		input := append(output[len(output)-1], decoderHiddenStates[i]...)

		decoderOutput := s2s.Decoder.FeedForward(input)[0]

		// Split the decoder output in two equal parts for the word output and the hidden state
		output = append(output, decoderOutput[0:s2s.VocabularySize])
		decoderHiddenStates = append(decoderHiddenStates, decoderOutput[s2s.VocabularySize:])
	}

	return output
}

func (s2s *Seq2Seq) PropagateBackward(output matrix) {
	var derivatives [][]Derivative

	for i := 0; i < len(output); i++ {
		idx := len(output) - i
		expectedEmbedding := matrix{output[idx]}

		derivatives = append(derivatives, []Derivative{
			s2s.Decoder.ComputeLastLayerDerivatives(expectedEmbedding),
		})
		
		
	}
}