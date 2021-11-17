// Golang program to show how
// to take input from the user
package main

import "fmt"

// main function
func main() {

	fmt.Println("Hello, GoLang!")

	fmt.Println("Enter Your First Name: ")
	var first string
	fmt.Scanln(&first)

	fmt.Println("Enter Second Last Name: ")
	var second string
	fmt.Scanln(&second)

	fmt.Println("YourFull Name is: ")
	fmt.Println(first +" " + second)

	fmt.Println("Press any key t exit ...")
	fmt.Scanln()

}
