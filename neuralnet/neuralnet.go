package neuralnet

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gookit/color"
	"gopkg.in/cheggaaa/pb.v1"
)

type NeuralNetwork struct {
	HiddenLayer      []float64
	InputLayer       []float64
	OutputLayer      []float64
	WeightHidden     [][]float64
	WeightOutput     [][]float64
	ErrOutput        []float64
	ErrHidden        []float64
	LastChangeHidden [][]float64
	LastChangeOutput [][]float64
	Rate1            float64
	Rate2            float64
}

func LoadNeuralNetwork(fileName string) *NeuralNetwork {
	inF, err := os.Open(fileName)
	if err != nil {
		panic("failed to load " + fileName)
	}
	defer inF.Close()

	decoder := json.NewDecoder(inF)
	neuralNetwork := &NeuralNetwork{}
	err = decoder.Decode(neuralNetwork)
	if err != nil {
		panic(err)
	}

	return neuralNetwork
}

// CreateNetwork returns a new network where layers are built with number of input, hidden and output
// layers and the learning rates.
func CreateNetwork(input, hidden, output int, rate1, rate2 float64) *NeuralNetwork {
	input++
	hidden++

	rand.Seed(time.Now().UnixNano())

	return &NeuralNetwork{
		Rate1:            rate1,
		Rate2:            rate2,
		InputLayer:       make([]float64, input),
		HiddenLayer:      make([]float64, hidden),
		OutputLayer:      make([]float64, output),
		ErrOutput:        make([]float64, output),
		ErrHidden:        make([]float64, hidden),
		WeightHidden:     RandomMatrix(hidden, input, -1.0, 1.0),
		WeightOutput:     RandomMatrix(output, hidden, -1.0, 1.0),
		LastChangeHidden: MakeMatrix(hidden, input, 0.0),
		LastChangeOutput: MakeMatrix(output, hidden, 0.0),
	}
}

func (neuralNetwork NeuralNetwork) Save(fileName string) {
	outF, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic("failed to dump the network to " + fileName)
	}
	defer outF.Close()

	encoder := json.NewEncoder(outF)
	err = encoder.Encode(neuralNetwork)
	if err != nil {
		panic(err)
	}
}

// FeedForward makes forward propagation for a single input
func (neuralNetwork *NeuralNetwork) FeedForward(input []float64) (output []float64) {
	if len(input)+1 != len(neuralNetwork.InputLayer) {
		panic("amount of input variable doesn't match")
	}

	for i, inputValue := range input {
		neuralNetwork.InputLayer[i] = inputValue
	}
	neuralNetwork.InputLayer[len(neuralNetwork.InputLayer)-1] = 1.0 // Bias node for input layer

	// Apply weights on the input layer to give the hidden layer
	hiddenLayer := ApplyWeights(
		len(neuralNetwork.HiddenLayer)-1,
		neuralNetwork.InputLayer,
		neuralNetwork.WeightHidden,
	)
	neuralNetwork.HiddenLayer = ApplyFunc(hiddenLayer, Sigmoid)
	neuralNetwork.HiddenLayer[len(neuralNetwork.HiddenLayer)-1] = 1.0 // Bias node for hidden layer

	// Apply weights on the hidden layer to give the output layer
	neuralNetwork.OutputLayer = ApplyWeights(
		len(neuralNetwork.OutputLayer),
		neuralNetwork.HiddenLayer,
		neuralNetwork.WeightOutput,
	)

	return neuralNetwork.OutputLayer
}

// FeedBack makes back propagation for a single target
func (neuralNetwork *NeuralNetwork) FeedBack(target []float64) {
	// Insert output errors in the array
	for i := 0; i < len(neuralNetwork.OutputLayer); i++ {
		neuralNetwork.ErrOutput[i] = neuralNetwork.OutputLayer[i] - target[i]
	}

	// Calculate the errors in the hidden layer
	for i := 0; i < len(neuralNetwork.HiddenLayer)-1; i++ {
		err := 0.0
		for j := 0; j < len(neuralNetwork.OutputLayer); j++ {
			err += neuralNetwork.ErrOutput[j] * neuralNetwork.WeightOutput[j][i]
		}
		neuralNetwork.ErrHidden[i] = err
	}

	// Apply the changes to the output weights
	for i := 0; i < len(neuralNetwork.OutputLayer); i++ {
		for j := 0; j < len(neuralNetwork.HiddenLayer); j++ {
			delta := neuralNetwork.ErrOutput[i]
			change := neuralNetwork.Rate1*delta*neuralNetwork.HiddenLayer[j] +
				neuralNetwork.Rate2*neuralNetwork.LastChangeOutput[i][j]

			neuralNetwork.WeightOutput[i][j] -= change
			neuralNetwork.LastChangeOutput[i][j] = change
		}
	}

	// Apply the changes to the hidden weights
	for i := 0; i < len(neuralNetwork.HiddenLayer)-1; i++ {
		for j := 0; j < len(neuralNetwork.InputLayer); j++ {
			delta := neuralNetwork.ErrHidden[i] * SigmoidDerivative(neuralNetwork.HiddenLayer[i])
			change := neuralNetwork.Rate1*delta*neuralNetwork.InputLayer[j] +
				neuralNetwork.Rate2*neuralNetwork.LastChangeHidden[i][j]

			neuralNetwork.WeightHidden[i][j] -= change
			neuralNetwork.LastChangeHidden[i][j] = change
		}
	}
}

func (neuralNetwork *NeuralNetwork) CalculateError(target []float64) float64 {
	errSum := 0.0
	for i := 0; i < len(neuralNetwork.OutputLayer); i++ {
		err := neuralNetwork.OutputLayer[i] - target[i]
		errSum += 0.5 * err * err
	}

	return errSum
}

func RandomIndexes(length int) []int {
	indexes := make([]int, length)
	for i := range indexes {
		indexes[i] = i
	}

	for i := 0; i < length; i++ {
		j := i + int(rand.Float64()*float64(length-i))
		indexes[i], indexes[j] = indexes[j], indexes[i]
	}

	return indexes
}

func (neuralNetwork *NeuralNetwork) Train(inputs [][]float64, targets [][]float64, iterations int) {
	if len(inputs[0])+1 != len(neuralNetwork.InputLayer) {
		panic("The amount of input variable doesn't match.")
	}
	if len(targets[0]) != len(neuralNetwork.OutputLayer) {
		panic("The amount of output variable doesn't match.")
	}

	blue := color.FgBlue.Render

	bar := pb.New(iterations).Postfix(fmt.Sprintf(" - %s", blue("Creating the neural network")))
	bar.Format("(██ )")
	bar.SetMaxWidth(60)
	bar.ShowCounters = false
	bar.Start()

	currentError := 0.0
	for i := 0; i < iterations; i++ {
		indexesArray := RandomIndexes(len(inputs))

		for j := 0; j < len(inputs); j++ {
			neuralNetwork.FeedForward(inputs[indexesArray[j]])
			neuralNetwork.FeedBack(targets[indexesArray[j]])
			// Sum the error to the current error
			if i == iterations-1 {
				currentError += neuralNetwork.CalculateError(targets[indexesArray[j]])
			}
		}
		// Increment the progress bar
		bar.Increment()
	}

	bar.Finish()

	arrangedError := fmt.Sprintf("%.5f", currentError/float64(len(inputs)))
	red := color.FgGreen.Render
	fmt.Printf("The error rate is %s.", red(arrangedError))
}
