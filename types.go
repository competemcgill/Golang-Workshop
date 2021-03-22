package main

import "fmt"

func main() {

	// ============================ Strings and Chars ============================

	// Strings are pretty straight-forward
	myStr := "The Go Gopher looks creepy!"

	// But you can also put them on multiple lines
	myStr2 := `Finally, decent support for line
	breaks in a string!`

	// A single character, just like Java
	myChar := 'c'

	// That's right, emojis in code...
	// (UTF-8 Source Code)
	coffee := '☕'

	fmt.Println(myStr)
	fmt.Println(myStr2)
	fmt.Println(myChar)
	fmt.Println(coffee)

	// ==========================================================================

	// ============================ Floats ============================

	// 64 Bit Float
	pi := 22. / 7

	// 32 Bit Float
	var lessPrecisePi float32 = 22. / 7

	fmt.Println(pi)
	fmt.Println(lessPrecisePi)

	// ===============================================================

	// ============================ Arrays ============================

	// Arrays have fixed size!
	// The below would be equivalent to:
	// {0, 0, 0, 0}
	var arr [4]int

	// You can also declare and initialize an
	// array in one GO:
	// The size is still fixed, automatically
	// calculated as 5.
	anotherArr := [...]int{1, 3, 2, 0, 100}

	fmt.Println(arr)
	fmt.Println(anotherArr)

	// ===============================================================

	// ============================ Slices ============================

	// Slices are dynamically sized
	slice := []int{1, 3, 5}
	// Which means we can append to them
	// (like ArrayLists in Java):
	slice = append(slice, 7, 9) // {1, 3, 5, 7, 9}

	// You can also initialize with default values:
	zeroSlice := make([]int, 5) // {0, 0, 0, 0, 0}

	// Arrays are copied so arrCopy and arr
	// refer to different arrays!:
	arrCopy := arr
	// arrCopy: {0, 0, 0, 0}
	// arr: {0, 0, 0, 0}
	arrCopy[0] = 1
	// arrCopy: {1, 0, 0, 0}
	// arr: {0, 0, 0, 0}

	// Slices refer to the slice so sliceCopy and
	// zeroSlice refer to the same slice!:
	sliceCopy := zeroSlice
	// zeroSlice: {0, 0, 0, 0, 0}
	// sliceCopy: {0, 0, 0, 0, 0}
	sliceCopy[0] = 1
	// zeroSlice: {1, 0, 0, 0, 0}
	// sliceCopy: {1, 0, 0, 0, 0}

	fmt.Println(slice)
	fmt.Println(zeroSlice)

	// ================================================================

	// ============================ Maps ============================

	// Maps are another useful data structure
	ages := map[string]int{
		"Imad": 22,
		"Moh":  34,
		"Val":  23,
	}
	// Happy Birthday!
	ages["Imad"] = 23

	// Maps also have reference semantics,
	// meaning they work like slices when we
	// "copy" them

	// ==============================================================

	// ============================ Pointers ============================

	// Intialize a pointer
	var p *int
	// Set it to the address of some variable
	myNum := 1
	p = &myNum

	// Dereference to read the value
	fmt.Println(*p)
	// Or to set the value
	*p = 21

	// ==================================================================

}
