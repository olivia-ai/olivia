package util

import (
	"encoding/csv"
	"io/ioutil"
	"os"
)

// ReadFile returns the bytes of a file searched in the path and beyond it
func ReadFile(path string) (bytes []byte) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		bytes, err = ioutil.ReadFile("../" + path)
	}

	if err != nil {
		panic(err)
	}

	return bytes
}

// ReadCSV returns a two-dimensional slice of strings from the given CSV file path
func ReadCSV(path string) (lines [][]string) {
	csvFile, err := os.Open(path)
	if err != nil {
		csvFile, err = os.Open("../" + path)
	}

	if err != nil {
		panic(err)
	}
	
	defer csvFile.Close()
    
    csvLines, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
		panic(err)
    }    

    return csvLines[1:]
}