package main

import "fmt"

func main() {

	// creating an array. the type  of elements and the length are
	// part of array's type
	var a [5]int
	fmt.Println(a) // 0 value

	// set and get a value array[index] syntax
	a[4] = 100
	fmt.Println(a)
	fmt.Println(a[4])

	// len returns the length of an array
	fmt.Println("len:", len(a))

	// syntax to declare and initialize
	b := [5]int{10, 20, 30, 40, 50}
	fmt.Println(b)

	// we can let the compiler count the numb of elems
	c := [...]int{2, 4, 6, 8, 10, 12}
	fmt.Println(c, len(c))

	// if we specify the index with :, the elems in between will be zeroed
	d := [...]int{100, 3: 400, 500}
	fmt.Println(d)

	// multi-dimension arrays
	var tab [3][3]int
	for i := range 3 {
		for j := range 3 {
			tab[i][j] = i + j
		}
	}
	fmt.Println(tab)

	// create and initialize
	camp := [2][2]int{
		{0, 0},
		{0, 0},
	}
	fmt.Println(camp)

}
