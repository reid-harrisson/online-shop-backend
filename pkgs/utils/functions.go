package utils

import (
	"math"
)

func Round(num float64) float64 {
	return math.Round(num*100) / 100
}
