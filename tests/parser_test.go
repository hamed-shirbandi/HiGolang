package tests

import (
	"testing"

	"codeblock.ir/helloGoLang/mathOperations"
)

func Test_Can_Parse_String_To_Number(t *testing.T) {
	str := "01"
	number := mathOperations.ParseStringToNumber(str)
	if number != 1 {
		t.Fail()
	}
}

func Test_Can_Parse_Number_To_String(t *testing.T) {
	number := 1
	str := mathOperations.ParseNumberToString(number)
	if str != "1" {
		t.Fail()
	}
}

func Test_Can_Sum(t *testing.T) {
	num1 := 5
	num2 := 6

	sum := mathOperations.Sum(num1, num2)
	if sum != 11 {
		t.Fail()
	}
}

func Test_Can_Parse_string_And_Sum(t *testing.T) {
	strNum1 := "5"
	strNum2 := "6"

	sum := mathOperations.ParseAndSum(strNum1, strNum2)
	if sum != 11 {
		t.Fail()
	}
}
