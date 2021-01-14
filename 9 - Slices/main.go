package main

import "fmt"

func main() {
	/*
	make: used to initialize slices with a specific length (also other stuff like make channels)
	 */
	s := make([]string, 3)

	sli := []int{1,2,3}
	fmt.Println("slice no make:", sli)

	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s), cap(s)*2)
	copy(c, s)
	fmt.Println("cpy:", c)


	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3", l)


	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i< 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)


	/*
		if capacity isn't specified, defaults to len
		4 is the capacity
		The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	 */
	sliceWithCap := make([]int, 2, 4)
	fmt.Println("slice with cap:", sliceWithCap)
	fmt.Println("specified cap:", cap(sliceWithCap))
	fmt.Println("length:", len(sliceWithCap))
	temp := sliceWithCap[:4] //OK because capacity is 4
	temp = sliceWithCap[:5] //ERROR because 5 > capacity
	fmt.Println(temp)


	/*
		output:
			slice no make: [1 2 3]
			emp: [  ]
			set: [a b c]
			get: c
			len: 3
			apd: [a b c d e f]
			cpy: [a b c d e f]
			sl1: [c d e]
			sl2: [a b c d e]
			sl3 [c d e f]
			dcl: [g h i]
			2d: [[0] [1 2] [2 3 4]]
			slice with cap: [0 0]
			specified cap: 4
			length: 2
			panic: runtime error: slice bounds out of range [:5] with capacity 4
	*/
}
