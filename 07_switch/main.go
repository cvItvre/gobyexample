package main

import (
	"fmt"
	"time"
)

func main() {

	// basic switch
	num := 1
	fmt.Println("write", num, "as:")
	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// use commas to separete multiple expressions in the same case statement
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its the weekend")
	default:
		fmt.Println("its a weekday")
	}

	// switch without an expression is an alternate if/else logic
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("its before noon")
	default:
		fmt.Println("its after noon")
	}

	// a type switch compares types instead of values
	// you can use to discover the type of an interface value
	whatAmI := func(i any) {
		switch t := i.(type) {
		case int:
			fmt.Println("im an int")
		case bool:
			fmt.Println("im a boolean")
		case float64:
			fmt.Println("im a float")
		case string:
			fmt.Println("im a string")
		default:
			fmt.Printf("i dontknow man... %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(3.1415)
	whatAmI("watsgood?")
	whatAmI(0)

}
