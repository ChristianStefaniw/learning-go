package main

import "fmt"

/*
A closure is a function value that references variables from outside its body.
The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables
 */


func intSeq() func() int {
	i := 0
	// anonymous function
	return func() int {
		i++
		return i
	}
}

func main() {
	/*
	intSeq function returns a closure. Each closure is bound to its own "i" variable
	 */
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())


	/*
	output:
		1
		2
		3
		1
	 */
}
