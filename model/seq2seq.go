package model

import "reflect"

// Seq2Seq contains the stucture of a sequence to sequence model including
// the encoder and decoder recurrent neural networks.
type Seq2Seq struct {
	Encoder NN
	Decoder NN
}

// CreateSeq2Seq creates and returns a new sequence to sequence (seq2seq) model.
func CreateSeq2Seq(
	embeddingSize int, 
	learningRate float64,
	encoderHiddenLayersNodes []int,
	decoderHiddenLayersNodes ...int,
) Seq2Seq {
	// Use the same hidden layers nodes count as for the encoder if not specified
	// for the decoder
	if (len(decoderHiddenLayersNodes) == 0) {
		decoderHiddenLayersNodes = encoderHiddenLayersNodes
	}

	return Seq2Seq{
		// Use twice the size of the embedding size for the encoder because we need to take the input
		// embedding and the previous one as input.
		Encoder: CreateNeuralNetwork(learningRate, 2 * embeddingSize, embeddingSize, encoderHiddenLayersNodes...),
		Decoder: CreateNeuralNetwork(learningRate, 2 * embeddingSize, embeddingSize, decoderHiddenLayersNodes...),
	}
}

func (s2s *Seq2Seq) FeedForward(embeddings matrix) {
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
}