package gonn

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"encoding/json"
	"os"
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
	Regression       bool
	Rate1            float64 //learning rate
	Rate2            float64
}

func sigmoid(X float64) float64 {
	return 1.0 / (1.0 + math.Pow(math.E, -float64(X)))
}

func dsigmoid(Y float64) float64 {
	return Y * (1.0 - Y)
}

func DumpNN(fileName string, nn *NeuralNetwork){
	out_f, err := os.OpenFile(fileName,os.O_CREATE | os.O_RDWR,0777)
	if err!=nil{
		panic("failed to dump the network to " + fileName)
	}
	defer out_f.Close()
	encoder := json.NewEncoder(out_f)
	err = encoder.Encode(nn)
	if err!=nil{
		panic(err)
	}
}

func LoadNN(fileName string) *NeuralNetwork{
	in_f, err := os.Open(fileName)
	if err!=nil{
		panic("failed to load "+fileName)
	}
	defer in_f.Close()
	decoder := json.NewDecoder(in_f)
	nn := &NeuralNetwork{}
	err = decoder.Decode(nn)
	if err!=nil{
		panic(err)
	}
	//fmt.Println(nn)
	return nn
}

func makeMatrix(rows, colums int, value float64) [][]float64 {
	mat := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		mat[i] = make([]float64, colums)
		for j := 0; j < colums; j++ {
			mat[i][j] = value
		}
	}
	return mat
}

func randomMatrix(rows, colums int, lower, upper float64) [][]float64 {
	mat := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		mat[i] = make([]float64, colums)
		for j := 0; j < colums; j++ {
			mat[i][j] = rand.Float64()*(upper-lower) + lower
		}
	}
	return mat
}

func DefaultNetwork(iInputCount, iHiddenCount, iOutputCount int, iRegression bool) *NeuralNetwork {
	return NewNetwork(iInputCount, iHiddenCount, iOutputCount, iRegression, 0.25, 0.1)
}

func NewNetwork(iInputCount, iHiddenCount, iOutputCount int, iRegression bool, iRate1, iRate2 float64) *NeuralNetwork {
	iInputCount += 1
	iHiddenCount += 1
	rand.Seed(time.Now().UnixNano())
	network := &NeuralNetwork{}
	network.Regression = iRegression
	network.Rate1 = iRate1
	network.Rate2 = iRate2
	network.InputLayer = make([]float64, iInputCount)
	network.HiddenLayer = make([]float64, iHiddenCount)
	network.OutputLayer = make([]float64, iOutputCount)
	network.ErrOutput = make([]float64, iOutputCount)
	network.ErrHidden = make([]float64, iHiddenCount)

	network.WeightHidden = randomMatrix(iHiddenCount, iInputCount, -1.0, 1.0)
	network.WeightOutput = randomMatrix(iOutputCount, iHiddenCount, -1.0, 1.0)

	network.LastChangeHidden = makeMatrix(iHiddenCount, iInputCount, 0.0)
	network.LastChangeOutput = makeMatrix(iOutputCount, iHiddenCount, 0.0)

	return network
}

func (self *NeuralNetwork) Forward(input []float64) (output []float64) {
	if len(input)+1 != len(self.InputLayer) {
		panic("amount of input variable doesn't match")
	}
	for i := 0; i < len(input); i++ {
		self.InputLayer[i] = input[i]
	}
	self.InputLayer[len(self.InputLayer)-1] = 1.0 //bias node for input layer

	for i := 0; i < len(self.HiddenLayer)-1; i++ {
		sum := 0.0
		for j := 0; j < len(self.InputLayer); j++ {
			sum += self.InputLayer[j] * self.WeightHidden[i][j]
		}
		self.HiddenLayer[i] = sigmoid(sum)
	}

	self.HiddenLayer[len(self.HiddenLayer)-1] = 1.0 //bias node for hidden layer
	for i := 0; i < len(self.OutputLayer); i++ {
		sum := 0.0
		for j := 0; j < len(self.HiddenLayer); j++ {
			sum += self.HiddenLayer[j] * self.WeightOutput[i][j]
		}
		if self.Regression {
			self.OutputLayer[i] = sum
		} else {
			self.OutputLayer[i] = sigmoid(sum)
		}
	}
	return self.OutputLayer[:]
}

func (self *NeuralNetwork) Feedback(target []float64) {
	for i := 0; i < len(self.OutputLayer); i++ {
		self.ErrOutput[i] = self.OutputLayer[i] - target[i]
	}

	for i := 0; i < len(self.HiddenLayer)-1; i++ {
		err := 0.0
		for j := 0; j < len(self.OutputLayer); j++ {
			if self.Regression {
				err += self.ErrOutput[j] * self.WeightOutput[j][i]
			} else {
				err += self.ErrOutput[j] * self.WeightOutput[j][i] * dsigmoid(self.OutputLayer[j])
			}

		}
		self.ErrHidden[i] = err
	}

	for i := 0; i < len(self.OutputLayer); i++ {
		for j := 0; j < len(self.HiddenLayer); j++ {
			change := 0.0
			delta := 0.0
			if self.Regression {
				delta = self.ErrOutput[i]
			} else {
				delta = self.ErrOutput[i] * dsigmoid(self.OutputLayer[i])
			}
			change = self.Rate1*delta*self.HiddenLayer[j] + self.Rate2*self.LastChangeOutput[i][j]
			self.WeightOutput[i][j] -= change
			self.LastChangeOutput[i][j] = change

		}
	}

	for i := 0; i < len(self.HiddenLayer)-1; i++ {
		for j := 0; j < len(self.InputLayer); j++ {
			delta := self.ErrHidden[i] * dsigmoid(self.HiddenLayer[i])
			change := self.Rate1*delta*self.InputLayer[j] + self.Rate2*self.LastChangeHidden[i][j]
			self.WeightHidden[i][j] -= change
			self.LastChangeHidden[i][j] = change

		}
	}
}

func (self *NeuralNetwork) CalcError(target []float64) float64 {
	errSum := 0.0
	for i := 0; i < len(self.OutputLayer); i++ {
		err := self.OutputLayer[i] - target[i]
		errSum += 0.5 * err * err
	}
	return errSum
}

func genRandomIdx(N int) []int {
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = i
	}
	//randomize
	for i := 0; i < N; i++ {
		j := i + int(rand.Float64()*float64(N-i))
		A[i], A[j] = A[j], A[i]
	}
	return A
}

func (self *NeuralNetwork) Train(inputs [][]float64, targets [][]float64, iteration int) {
	if len(inputs[0])+1 != len(self.InputLayer) {
		panic("amount of input variable doesn't match")
	}
	if len(targets[0]) != len(self.OutputLayer) {
		panic("amount of output variable doesn't match")
	}

	iter_flag := -1
	for i := 0; i < iteration; i++ {
		idx_ary := genRandomIdx(len(inputs))
		cur_err := 0.0
		for j := 0; j < len(inputs); j++ {
			self.Forward(inputs[idx_ary[j]])
			self.Feedback(targets[idx_ary[j]])
			cur_err += self.CalcError(targets[idx_ary[j]])
			if (j+1)%1000 == 0 {
				if iter_flag != i {
					fmt.Println("")
					iter_flag = i
				}
				fmt.Printf("iteration %vth / progress %.2f %% \r", i+1, float64(j)*100/float64(len(inputs)))
			}
		}
		if (iteration >= 10 && (i+1)%(iteration/10) == 0) || iteration < 10 {
			fmt.Printf("\niteration %vth MSE: %.5f", i+1, cur_err / float64(len(inputs)))
		}
	}
	fmt.Println("\ndone.")
}

func (self *NeuralNetwork) TrainMap(inputs []map[int]float64, targets [][]float64, iteration int) {
	if len(targets[0]) != len(self.OutputLayer) {
		panic("amount of output variable doesn't match")
	}

	iter_flag := -1
	for i := 0; i < iteration; i++ {
		idx_ary := genRandomIdx(len(inputs))
		cur_err := 0.0
		for j := 0; j < len(inputs); j++ {
			self.ForwardMap(inputs[idx_ary[j]])
			self.FeedbackMap(targets[idx_ary[j]],inputs[idx_ary[j]] )
			cur_err += self.CalcError(targets[idx_ary[j]])
			if (j+1)%1000 == 0 {
				if iter_flag != i {
					fmt.Println("")
					iter_flag = i
				}
				fmt.Printf("iteration %vth / progress %.2f %% \r", i+1, float64(j)*100/float64(len(inputs)))
			}
		}
		if (iteration >= 10 && (i+1)%(iteration/10) == 0) || iteration < 10 {
			fmt.Printf("\niteration %vth MSE: %.5f", i+1, cur_err / float64(len(inputs)))
		}
	}
	fmt.Println("\ndone.")
}


func (self *NeuralNetwork) ForwardMap(input map[int]float64) (output []float64) {
	for k,v := range input {
		self.InputLayer[k] = v
	}
	self.InputLayer[len(self.InputLayer)-1] = 1.0 //bias node for input layer

	for i := 0; i < len(self.HiddenLayer)-1; i++ {
		sum := 0.0
		for j,_ := range input{
			sum += self.InputLayer[j] * self.WeightHidden[i][j]
		}
		self.HiddenLayer[i] = sigmoid(sum)
	}

	self.HiddenLayer[len(self.HiddenLayer)-1] = 1.0 //bias node for hidden layer
	for i := 0; i < len(self.OutputLayer); i++ {
		sum := 0.0
		for j := 0; j < len(self.HiddenLayer); j++ {
			sum += self.HiddenLayer[j] * self.WeightOutput[i][j]
		}
		if self.Regression {
			self.OutputLayer[i] = sum
		} else {
			self.OutputLayer[i] = sigmoid(sum)
		}
	}
	return self.OutputLayer[:]
}

func (self *NeuralNetwork) FeedbackMap(target []float64,input map[int]float64) {
	for i := 0; i < len(self.OutputLayer); i++ {
		self.ErrOutput[i] = self.OutputLayer[i] - target[i]
	}

	for i := 0; i < len(self.HiddenLayer)-1; i++ {
		err := 0.0
		for j := 0; j < len(self.OutputLayer); j++ {
			if self.Regression {
				err += self.ErrOutput[j] * self.WeightOutput[j][i]
			} else {
				err += self.ErrOutput[j] * self.WeightOutput[j][i] * dsigmoid(self.OutputLayer[j])
			}

		}
		self.ErrHidden[i] = err
	}

	for i := 0; i < len(self.OutputLayer); i++ {
		for j := 0; j < len(self.HiddenLayer); j++ {
			change := 0.0
			delta := 0.0
			if self.Regression {
				delta = self.ErrOutput[i]
			} else {
				delta = self.ErrOutput[i] * dsigmoid(self.OutputLayer[i])
			}
			change = self.Rate1*delta*self.HiddenLayer[j] + self.Rate2*self.LastChangeOutput[i][j]
			self.WeightOutput[i][j] -= change
			self.LastChangeOutput[i][j] = change

		}
	}

	for i := 0; i < len(self.HiddenLayer)-1; i++ {
		for j , _ := range input {
			delta := self.ErrHidden[i] * dsigmoid(self.HiddenLayer[i])
			change := self.Rate1*delta*self.InputLayer[j] + self.Rate2*self.LastChangeHidden[i][j]
			self.WeightHidden[i][j] -= change
			self.LastChangeHidden[i][j] = change

		}
	}
}
