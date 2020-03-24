package network

import (
	"testing"
)

func TestSigmoid(t *testing.T) {
	result := Sigmoid(10)
	if result != 0.9999546021312976 {
		t.Errorf("Sigmoid() failed.")
	}
}

func TestMultipliesByTwo(t *testing.T) {
	result := MultipliesByTwo(2)
	if result != 4 {
		t.Errorf("MultipliesByTwo() failed.")
	}
}

func TestSubtractsOne(t *testing.T) {
	result := SubtractsOne(21)
	if result != 20 {
		t.Errorf("SubtractsOne() failed.")
	}
}
