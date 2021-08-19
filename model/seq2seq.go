package model

// Seq2Seq contains the stucture of a sequence to sequence model including
// the encoder and decoder recurrent neural networks.
type Seq2Seq struct {
	Encoder NN
	Decoder NN
}

// CreateSeq2Seq creates and returns a new sequence to sequence (seq2seq) model.
func CreateSeq2Seq(
	embedingSize int, 
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
		Encoder: CreateRNN(learningRate, embedingSize, encoderHiddenLayersNodes...),
		Decoder: CreateRNN(learningRate, embedingSize, decoderHiddenLayersNodes...),
	}
}