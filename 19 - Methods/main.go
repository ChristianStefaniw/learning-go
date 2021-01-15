package main

import "fmt"

/*
methods are defined on struct types
 */

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim", r.perim())


	/*
	Go automatically handles conversion between values and pointers for method calls.
	 */
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())

	/*
	output:
		area: 50
		perim 30
		area: 50
		perim: 30
	 */
}
