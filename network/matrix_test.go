package network

import (
	"reflect"
	"testing"
)

func TestMatrix_ApplyFunction(t *testing.T) {
	a := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	ApplyFunction(a, func(x float64) float64 {
		return x + 1
	})

	// Excepted value
	r := [][]float64{
		{2, 3, 4},
		{5, 6, 7},
	}

	if !reflect.DeepEqual(a, r) {
		t.Errorf("ApplyFunction(fn(x)=x+1) failed, excepted %v, got %v", r, a)
	}
}

func TestMatrix_Add(t *testing.T) {
	a := Matrix{[][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}}

	b := Matrix{[][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}}

	a.Add(b)

	// Excepted value
	r := Matrix{[][]float64{
		{2, 4, 6},
		{8, 10, 12},
	}}

	if !reflect.DeepEqual(a.value, r.value) {
		t.Errorf("Add(%v) failed, excepted %v, got %v", b, r, a)
	}
}

func TestMatrix_Substract(t *testing.T) {
	a := Matrix{[][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}}

	b := Matrix{[][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}}

	a.Substract(b)

	// Excepted value
	r := Matrix{[][]float64{
		{0, 0, 0},
		{0, 0, 0},
	}}

	if !reflect.DeepEqual(a.value, r.value) {
		t.Errorf("Remove(%v) failed, excepted %v, got %v", b, r, a)
	}
}

func TestMatrix_DotProduct(t *testing.T) {
	a := Matrix{[][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}}

	b := Matrix{[][]float64{
		{7, 8},
		{9, 10},
		{11, 12},
	}}

	// Actual value
	p := a.DotProduct(b)

	// Excepted value
	r := Matrix{[][]float64{
		{58, 64},
		{139, 154},
	}}

	if !reflect.DeepEqual(p.value, r.value) {
		t.Errorf("DotProduct(%v, %v) failed, excepted %v, got %v", a, b, r, p)
	}
}

func TestMatrix_Transpose(t *testing.T) {
	a := Matrix{[][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}}

	aT := a.Transpose()

	r := Matrix{[][]float64{
		{1, 4},
		{2, 5},
		{3, 6},
	}}

	if !reflect.DeepEqual(aT.value, r.value) {
		t.Errorf("Transpose failed, excepted %v, got %v", r, aT)
	}
}
