package data

import "testing"

func TestReadCSVConversationalDataset(t *testing.T) {
	for i, line := range ReadCSVConversationalDataset("data/mock.csv") {
		if i == 0 && (line.Question != "How are you?" || line.Answer != "Good and you?") {
			t.Errorf("data.ReadCSVConversationalDataset() failed.")
		}
	}
}