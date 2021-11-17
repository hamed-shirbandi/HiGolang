package mathOperations

// To sum 2 int value
func Sum(x int, y int) int {
	return x + y
}

// To sum 2 string number
func ParseAndSum(x string, y string) int {
	var num1 = ParseStringToNumber(x)
	var num2 = ParseStringToNumber(y)
	return Sum(num1, num2)
}
