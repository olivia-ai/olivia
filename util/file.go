package util

import (
	"io/ioutil"
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
