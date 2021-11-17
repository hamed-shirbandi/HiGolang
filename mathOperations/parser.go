package mathOperations

import (
	"strconv"
)

func ParseStringToNumber(x string) int {

	intVar, _ := strconv.Atoi(x)

	return intVar
}
