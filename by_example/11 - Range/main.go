package main

import "fmt"

func main() {
	/*
	range iterates over elements in a variety of data structures
	 */

	nums := []int{2,3,4}
	sum := 0
	/*
	range on arrays and slices provides both the index and value for each entry
	 */
	for _, num := range nums{
		sum += num
	}
	fmt.Println("sum:",sum)

	for i, num := range nums{
		if num == 3{
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range kvs{
		fmt.Printf("%s -> %s\n", key, value)
	}

	for key := range kvs{
		fmt.Println("key:", key)
	}

	/*
	range on strings iterates over Unicode code points
	 */
	for i, c := range "go"{
		fmt.Println(i, c)
	}


	/*
	output:
		sum: 9
		index: 1
		a -> apple
		b -> banana
		key: b
		key: a
		0 103
		1 111
	 */
}
