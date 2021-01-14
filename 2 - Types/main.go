package main

/*
Go's basic types

-bool

-string

-int  int8  int16  int32  int64  ________________________________
-uint uint8 uint16 uint32 uint64 uintptr						|
																|
-byte alias for uint8											|
																|
-rune alias for int32 - represents a Unicode code point			|
																|
-float32 float64												|
																|
-complex64 complex128											|
																|
																|
																|
Differences between integer types  <____________________________|

 Type      Capacity

   Int16 -- (-32,768 to +32,767)

   Int32 -- (-2,147,483,648 to +2,147,483,647)

   Int64 -- (-9,223,372,036,854,775,808 to +9,223,372,036,854,775,807)
 */


import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	/*
	output:
		Type: bool Value: false
		Type: uint64 Value: 18446744073709551615
		Type: complex128 Value: (2+3i)
	 */
}