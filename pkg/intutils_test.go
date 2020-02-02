package intutils

import (
	"fmt"
	"testing"
)

func TestSquareSingle(t *testing.T) {
	square := Square(2)
	if square != 4 {
		t.Errorf("Square(2) = %f; want 4", square)
	}
}

func TestSquareMulti(t *testing.T) {
	var tests = []struct {
		num    float64
		square float64
	}{
		{-1, 1},
		{0, 0},
		{1, 1},
		{2, 4},
		{3, 9},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%f", tt.num)
		t.Run(testname, func(t *testing.T) {
			square := Square(tt.num)
			if square != tt.square {
				t.Errorf("got %f, want %f", square, tt.square)
			}
		})
	}
}
