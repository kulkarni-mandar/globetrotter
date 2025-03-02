package utils

import "strconv"

func ToInt(valStr string) int {
	valInt, _ := strconv.Atoi(valStr)
	return valInt
}
