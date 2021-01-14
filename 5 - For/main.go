package main

import "fmt"

func main() {
	i := 0

	// Single condition
	for i <= 3{
		fmt.Println(i)
		i++
	}

	// Classic for loop
	for j := 7; j <= 9; j++{
		fmt.Println(j)
	}

	/*
	Without a condition will loop repeatedly until you break out of the loop or return from the enclosing function
	 */
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++{
		if n%2 == 0{
			//continue to the next iteration of the loop
			continue
		}
		fmt.Println(n)
	}


	/*
		output:
			0
			1
			2
			3
			7
			8
			9
			loop
			1
			3
			5
	*/
}
