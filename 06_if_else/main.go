package main

import "fmt"

func main() {

	// basic example
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// you can have an if statement without an else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// operators like && and || are useful in conditions
	if (8%2 == 0) || (7%2 == 0) {
		fmt.Println("8 or 7 are even")
	}

	// a statement can preced conditions
	// any variables declared in this statement are available in the current and all subsequent branches
	if age := 15; age < 16 {
		fmt.Println("you cannot drive")
	} else {
		fmt.Println("you're allowed to drive")
	}

	// there is no ternary if :(

}
