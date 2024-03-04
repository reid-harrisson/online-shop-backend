package utils

import (
	"math"
	"regexp"
	"strings"
)

func Round(num float64) float64 {
	return math.Round(num*100) / 100
}

func CleanSpecialLetters(sku string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	return strings.ToUpper(reg.ReplaceAllString(sku, ""))
}
