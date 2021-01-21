package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Multiple Short Declare
//
//  Declare two variables using multiple short declaration
//
// EXPECTED OUTPUT
//  14 true
// -----

func main() {
	one, two := 14,  true
	fmt.Println(one, two)
}
