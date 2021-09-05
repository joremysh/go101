package main

import (
	"fmt"
)

func example4(myAge int) {
	myAge = 30
	if myAge < 13 { // <-- "{" should be here
		fmt.Println("Child")
		// "}" should be on the left of "else"
	} else if myAge >= 13 && myAge < 20 {
		fmt.Println("Teen")
	} else {
		fmt.Println("Adult")
	} // result: Adult

	switch {
	case myAge < 13:
		fmt.Println("Child") // will jump out of switch block after a case
	default:
		fmt.Println("Adult")
	} // result: Adult
}
