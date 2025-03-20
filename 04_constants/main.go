package main

import (
	"fmt"
	"math"
)

// const declares a constant
const s string = "this is a constant :)"

func main() {

	// Go supports consts of char, string, booleans and numeric values
	fmt.Println(s)

	// const statement can apper anywhere a var statement can
	const pi = 3.1415

	// constant expressions performs arithmetic with arbitrary precision
	const d = 3e20 / pi
	fmt.Println(d)

	// numeric constants has no type until its given one, e.g. explicit conversion
	fmt.Println(float64(d))

	// a number can be given a type in a context that requires one
	// e.g. variable assignemnts or function call
	fmt.Println(math.Sin(d))

}
