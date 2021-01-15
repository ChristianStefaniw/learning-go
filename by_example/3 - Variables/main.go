package main

import "fmt"

func main() {
	var a = "string variable"
	fmt.Println(a)

	//variables are int's
	var b,c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	//Go integer variables are initialized to 0
	var e int
	fmt.Println(e)


	/*
	  ":=" is shorthand for declaring and initializing a variable
	 	   same as
	       var f string = "walrus operator" in this case.
	 */
	f := "walrus operator"
	fmt.Println(f)


	/*
	output:
		string variable
		1 2
		true
		0
		walrus operator
	 */
}
