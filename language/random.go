package language

import (
	"errors"
	"regexp"
	"sort"
	"strconv"
)

var decimal = "\\b\\d+([\\.,]\\d+)?"

// FindRangeLimits finds the range for random numbers and returns a sorted integer array
func FindRangeLimits(local, entry string) ([]int, error) {
	decimalsRegex := regexp.MustCompile(decimal)
	limitStrArr := decimalsRegex.FindAllString(entry, 2)
	limitArr := make([]int, 0)

	if limitStrArr == nil {
		return make([]int, 0), errors.New("No range")
	}

	if len(limitStrArr) != 2 {
		return nil, errors.New("Need 2 numbers, a lower and upper limit")
	}

	for _, v := range limitStrArr {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("Non integer range")
		}
		limitArr = append(limitArr, num)
	}

	sort.Ints(limitArr)
	return limitArr, nil
}
