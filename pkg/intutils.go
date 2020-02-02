package intutils

import (
	"math"
)

// Square returns number raised to the second power
func Square(num float64) float64 {
	square := math.Pow(num, 2)
	return square
}
