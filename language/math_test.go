package language

import (
	"testing"
)

func TestFindMathOperation(t *testing.T) {
	sentence := "Calculate cos(0.5) * 5.3 please"
	excepted := "cos(0.5) * 5.3"
	operation := FindMathOperation(sentence)

	if operation != excepted {
		t.Errorf("FindMathOperation() failed, excepted %s got %s.", excepted, operation)
	}
}

func TestFindNumberOfDecimals(t *testing.T) {
	sentence := "Calculate cos(0.5) * 5.3 with 8 decimals please"
	excepted := 8
	decimals := FindNumberOfDecimals("en", sentence)

	if decimals != excepted {
		t.Errorf("FindNumberOfDecimals() failed, excepted %d got %d.", excepted, decimals)
	}
}
