package language

import (
	"reflect"
	"testing"
)

func TestFindRangeLimits(t *testing.T) {
	limitArr, err := FindRangeLimits("", "Pick a number between 10 and 50")
	expected := []int{10, 50}
	if err != nil {
		t.Errorf("FindRangeLimit() failed with error %s", err.Error())
	}
	if !reflect.DeepEqual(limitArr, expected) {
		t.Errorf("FindRangeLimit() failed, excepted %v got %v.", expected, limitArr)
	}

	limitArr, err = FindRangeLimits("", "Pick a number between 50 and 10")
	expected = []int{10, 50}
	if err != nil {
		t.Errorf("FindRangeLimit() failed with error %s", err.Error())
	}
	if !reflect.DeepEqual(limitArr, expected) {
		t.Errorf("FindRangeLimit() failed, excepted %v got %v.", expected, limitArr)
	}

	limitArr, err = FindRangeLimits("", "Pick a number between 10.8 and 50")
	expectedErr := "Non integer range"
	if err.Error() != expectedErr {
		t.Errorf("FindRangeLimit() failed with error %s instead of %s", err.Error(), expectedErr)
	}
}
