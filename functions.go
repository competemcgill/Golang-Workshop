package main

import "fmt"

func main() {
	// Single return value
	firstSum := sum(3, 4)
	// Comma adds a space automatically!
	fmt.Println("Sum:", firstSum)

	secondSum, product := sumAndProduct(5, 6)

	fmt.Println("Sum:", secondSum, "Product:", product)
}

// Function parameters go in the first pair of parentheses
// Return type(s) go after
func sum(x, y int) int {
	return x + y
}

// Functions can return multiple values!
func sumAndProduct(x, y int) (int, int) {
	return x + y, x * y
}

// We can also provide names to our return values
func namedSum(x, y int) (sum int) {
	// Its as if we've declared sum up here:
	// var sum int
	sum = x + y
	// And we return like this:
	// return sum
	return
}

func namedSumAndProduct(x, y int) (sum int, product int) {
	sum = x + y
	product = x * y
	return
}
