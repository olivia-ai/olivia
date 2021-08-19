package model

import (
	"testing"
)

func TestSigmoid(t *testing.T) {
	result := sigmoid(10)
	if result != 0.9999546021312976 {
		t.Errorf("Sigmoid() failed.")
	}
}

func TestMultipliesByTwo(t *testing.T) {
	result := multipliesByTwo(2)
	if result != 4 {
		t.Errorf("MultipliesByTwo() failed.")
	}
}

func TestSubtractsOne(t *testing.T) {
	result := subtractsOne(21)
	if result != 20 {
		t.Errorf("SubtractsOne() failed.")
	}
}
