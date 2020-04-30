package analysis

import (
	"testing"
)

func TestGetModuleCoverage(t *testing.T) {
	notCovered, coverage := GetModuleCoverage("en")

	if len(notCovered) != 0 || coverage != 100 {
		t.Errorf("GetModuleCoverage() failed.")
	}
}

func TestGetIntentCoverage(t *testing.T) {
	notCovered, coverage := GetIntentCoverage("en")

	if len(notCovered) != 0 || coverage != 100 {
		t.Errorf("GetIntentCoverage() failed.")
	}
}

func TestGetMessageCoverage(t *testing.T) {
	notCovered, coverage := GetIntentCoverage("en")

	if len(notCovered) != 0 || coverage != 100 {
		t.Errorf("GetIntentCoverage() failed.")
	}
}
