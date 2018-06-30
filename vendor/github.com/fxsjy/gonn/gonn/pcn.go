package gonn

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"encoding/json"
)

type PCNNetwork struct{
	InputLayer map[string]float64
	OutputLayer []float64
	WeightOutput []map[string]float64
	LastChangeOutput []map[string]float64
	Regression bool
	Rate1 float64
	Rate2 float64
}



func DumpPCN(fileName string, nn *PCNNetwork){
	out_f, err := os.OpenFile(fileName,os.O_CREATE | os.O_RDWR,0777)
	if err!=nil{
		panic("failed to dump the network to " + fileName)
	}
	defer out_f.Close()
	encoder := json.NewEncoder(out_f)
	encoder.Encode(nn)
}

func LoadPCN(fileName string) *PCNNetwork{
	in_f, err := os.Open(fileName)
	if err!=nil{
		panic("failed to load "+fileName)
	}
	defer in_f.Close()
	decoder := json.NewDecoder(in_f)
	nn := &PCNNetwork{}
	decoder.Decode(nn)
	//fmt.Println(nn)
	return nn
}



func DefaultPCNNetwork(iOutputCount int,iRegression bool) *PCNNetwork{
	return NewPCNNetwork(iOutputCount,iRegression,0.25,0.1)
}


func SparseMatrix(rows int) []map[string]float64 {
	mat := make([]map[string]float64, rows)
	for i := 0; i < rows; i++ {
		mat[i] = make(map[string]float64)
	}
	return mat
}

func NewPCNNetwork(iOutputCount int,iRegression bool,iRate1,iRate2 float64) * PCNNetwork{
	self := &PCNNetwork{}
	rand.Seed(time.Now().UnixNano())
	self.InputLayer = make(map[string]float64)
	self.InputLayer["~"]=1.0 //bias node
	self.OutputLayer = make([]float64,iOutputCount)
	self.WeightOutput = SparseMatrix(iOutputCount)
	self.LastChangeOutput = SparseMatrix(iOutputCount,)
	self.Regression = iRegression
	self.Rate1 = iRate1
	self.Rate2 = iRate2
	return self
}


func (self *PCNNetwork) CalcError(target []float64) float64 {
	errSum := 0.0
	for i := 0; i < len(self.OutputLayer); i++ {
		err := self.OutputLayer[i] - target[i]
		errSum += 0.5 * err * err
	}
	return errSum
}



func (self *PCNNetwork) TrainMap(inputs []map[string]float64, targets [][]float64, iteration int) {

	if len(targets[0]) != len(self.OutputLayer) {
		panic("amount of output variable doesn't match")
	}
	

	iter_flag := -1
	for i := 0; i < iteration; i++ {
		idx_ary := genRandomIdx(len(inputs))
		cur_err := 0.0
		for j := 0; j < len(inputs); j++ {
			self.ForwardMap(inputs[idx_ary[j]])
			self.FeedbackMap(targets[idx_ary[j]], inputs[idx_ary[j]] )
			cur_err += self.CalcError(targets[idx_ary[j]])
			if (j+1)%1000 == 0 {
				if iter_flag != i {
					fmt.Println("")
					iter_flag = i
				}
				fmt.Printf("iteration %v / progress %.2f %% \r", i+1, float64(j)*100/float64(len(inputs)))
			}
		}
		if (iteration >= 10 && (i+1)%(iteration/10) == 0) || iteration < 10 {
			fmt.Printf("\niteration %v err: %.5f", i+1, cur_err / float64(len(inputs)))
		}
	}
	fmt.Println("\ndone.")
}

func (self *PCNNetwork) ForwardMap(input map[string]float64) []float64{
	
	for i,_ := range input {
		self.InputLayer[i] = input[i]
	}

	for i := 0; i < len(self.OutputLayer); i++ {
		sum := 0.0
		for j,_:= range input {
			if _,ok := self.WeightOutput[i][j] ; !ok{
				self.WeightOutput[i][j]  = rand.Float64()*2.0 - 1.0
			}
			sum += self.InputLayer[j] * self.WeightOutput[i][j]
		}
		if(self.Regression){
			self.OutputLayer[i] = sum
		}else{
			self.OutputLayer[i] = sigmoid(sum)
		}
	}

	return self.OutputLayer[:]
}


func (self *PCNNetwork) FeedbackMap(target []float64, input map[string]float64) {
	for i := 0; i < len(self.OutputLayer); i++ {
		err_i := self.OutputLayer[i] - target[i]
		for j,_:= range input {
			change := 0.0
			delta := 0.0
			if self.Regression {
				delta = err_i
			} else {
				delta = err_i * dsigmoid(self.OutputLayer[i])
			}
			change = self.Rate1*delta*self.InputLayer[j] + self.Rate2*self.LastChangeOutput[i][j]
			self.WeightOutput[i][j] -= change
			self.LastChangeOutput[i][j] = change
		}
	}
}

