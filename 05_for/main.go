package main

import "fmt"

func main() {

	// for is the only Go's looping structure
	// basic type, single condition

	i := 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// classic initial/condition/after for loop
	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}

	// another way todo "run this N times" iteration is range over an integer
	for iter := range 3 {
		fmt.Println("range:", iter)
	}

	// run until break or return from the enclosing function
	for {
		fmt.Println("whiletrue")
		break
	}

	// you can also continue to the next iteration
	for num := range 6 {
		if num%2 != 0 {
			fmt.Println(num)
		}
		continue
	}

}
