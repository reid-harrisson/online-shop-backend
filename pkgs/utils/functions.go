package utils

import (
	"fmt"
	"math"
	"strings"
)

func GenerateSKU(name string, id uint64) string {
	return strings.ToUpper(name)[0:3] + fmt.Sprintf("-%03d", id)[0:4]
}

func Round(num float64) float64 {
	return math.Round(num*100) / 100
}
