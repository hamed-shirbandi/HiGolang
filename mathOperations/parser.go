package mathOperations

import (
	"strconv"
)

func ParseStringToNumber(x string) int {

	intVar, _ := strconv.Atoi(x)

	return intVar
}

func ParseNumberToString(x int) string {

	return strconv.Itoa(x)

}
