package main

import (
	"fmt"
	"math"
)

const S string = "constant"

func main() {
	fmt.Println(S)

	const n = 500000000

	/*
		Constant expressions perform arithmetic with arbitrary precision.
	 */
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))


	fmt.Println(math.Sin(n))


	/*
		output:
			constant
			6e+11
			600000000000
			-0.28470407323754404
	 */
}
