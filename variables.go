package main

import "fmt"

func main() {

	// Declare variable with type
	var myInt int
	// Initialize value
	myInt = 3

	// Declaration & Initialization with type
	//automatically inferred
	myOtherInt := 5

	fmt.Println(myInt, myOtherInt)
}
