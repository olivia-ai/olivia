package model

import "testing"

func TestS2SFeedForward(t *testing.T) {
	model := CreateSeq2Seq(10, 0.25)
	empty := make([]float64, 10)
	model.FeedForward([][]float64{empty, empty, empty})
}