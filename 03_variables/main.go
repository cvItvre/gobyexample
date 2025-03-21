package main

import "fmt"

func main() {

	// var declares 1 or more variables
	var a = "initial"
	fmt.Println(a)

	// you can declare multiples variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables
	var d = true
	fmt.Println(d)

	// zero-valued decalration
	var e int
	fmt.Println(e)

	// := syntax is a shorthand for declaration and initialization
	f := "apple"
	fmt.Println(f)

}
