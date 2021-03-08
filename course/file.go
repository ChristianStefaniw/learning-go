package main

import (
	"fmt"
)

func bytes() {
	var b byte
	b = 2

	fmt.Printf("%08b = %d\n", b, b)

	b = 255

	fmt.Printf("%08b = %d\n", b, b)
}

func types() {
	type gram float64
	type ounce float64

	var g gram = 1000
	var o ounce

	o = ounce(g) * 0.035274

	fmt.Println(o)


	type (
		Gram     int
		Kilogram Gram
		Ton      Kilogram
	)
	var (
		salt   Gram     = 100
		apples Kilogram = 5
		truck  Ton      = 10
	)
	fmt.Printf("salt: %d, apples: %d, truck: %d\n", salt, apples, truck)
	salt = Gram(apples)
	apples = Kilogram(truck)
	truck = Ton(Gram(int(apples)))


	var(
		byteVal byte
		uint8Val uint8
		intVal int
	)
	uint8Val = byteVal
	var (
		runeVal rune
		int32Val int32
	)
	runeVal = int32Val
	fmt.Printf("%T, %T, %T, %T", uint8Val, int32Val, intVal, runeVal)

}

func main() {
	bytes()
	types()
}
