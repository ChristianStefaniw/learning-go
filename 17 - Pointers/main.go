package main

import "fmt"

func zeroVal(ival int)  {
	ival = 0
}

func zeroPtr(iptr *int)  {
	/*
	dereferences the pointer from its memory address to the current value at that address
	Assigning a value to a dereferenced pointer changes the value at the referenced address.
	 */
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroVal(i)
	fmt.Println("zeroVal:", i)

	zeroPtr(&i)
	fmt.Println("zeroPtr:", i)

	fmt.Println("pointer:", &i)


	/*
	output:
		initial: 1
		zeroVal: 1
		zeroPtr: 0
		pointer: 0xc0000a2058
	 */
}
