package main

import "fmt"

func main() {

	// In Go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued.
	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable. This syntax is only available inside functions.
	f := "apple"
	fmt.Println(f)
}