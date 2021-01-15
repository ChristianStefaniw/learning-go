package main

import "fmt"

/*
Variadic functions can be called with any number of trailing arguments.
Example: fmt.Println
 */

func sum(nums ...int){
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums{
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...) // "..." is unpacking the slice


	/*
	output:
		[1 2] 3
		[1 2 3] 6
		[1 2 3 4] 10
	 */
}
