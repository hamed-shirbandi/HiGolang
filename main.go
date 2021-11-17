package main

import (
	"fmt"

	"codeblock.ir/helloGoLang/mathOperations"
)

func main() {

	var firstNum string
	var secondNum string

	fmt.Println("enter num 1 :")
	fmt.Scanln(&firstNum)
	fmt.Println("enter num 2 :")
	fmt.Scanln(&secondNum)

	fmt.Println("the sum is :")

	var a = mathOperations.ParseAndSum(firstNum, secondNum)
	fmt.Println(a)
}
