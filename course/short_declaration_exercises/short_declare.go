package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Short Declare
//
//  Declare and then print four variables using
//  the short declaration statement.
//
// EXPECTED OUTPUT
//  i: 314 f: 3.14 s: Hello b: true
// -------

func main() {
	one := 314
	two := 3.14
	three := "hello"
	four := true

	fmt.Println(one, two, three, four)
}
