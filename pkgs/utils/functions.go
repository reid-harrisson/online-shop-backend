package utils

import (
	"fmt"
	"strings"
)

func GenerateSKU(name string, id uint64) string {
	return strings.ToUpper(name)[0:3] + fmt.Sprintf("-%03d", id)[0:4]
}
