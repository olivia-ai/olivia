package modules

import "testing"

func TestFindMathOperation(t *testing.T) {
	sentencesOperations := map[string]string{
		"Calculate cos(0.5)*2.2+(4^24) please":                  "cos(0.5)*2.2+(4^24)",
		"What's sqrt(tan(1.2)) plz":                             "sqrt(tan(1.2))",
		"Give me the result of log(34)*17/(842+1.23232) please": "log(34)*17/(842+1.23232)",
		"What's 4 x 3 ?":                                        "4 * 3",
	}

	for sentence, operation := range sentencesOperations {
		operationFound := FindMathOperation(sentence)

		if operationFound != operation {
			t.Errorf("Expected \"%s\" operation for \"%s\", found \"%s\"", operation, sentence, operationFound)
		}
	}
}

func TestFindNumberOfDecimals(t *testing.T) {
	sentencesDecimals := map[string]int{
		"calculate cos(0.5) with 15 decimals please":              15,
		"what's sin(3.2)*0.3^2 with a total of 7 decimals":        7,
		"could you calculate 3*4 and the number of decimals is 3": 3,
		"calculate tan(2.3) as a 14-decimal number":               14,
		"give me the result of 123^3 with 4 decimal digits":       4,
	}

	for sentence, decimals := range sentencesDecimals {
		decimalsFound := FindNumberOfDecimals(sentence)

		if decimalsFound != decimals {
			t.Errorf("Expected \"%d\" number of decimals for \"%s\", found \"%d\"", decimals, sentence, decimalsFound)
		}
	}
}
