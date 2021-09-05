package main

import (
	"fmt"
)

func example2() {
	// call function "foo"
	b, a := foo(2020, "Iron Man")
	fmt.Printf("%d %s\n", a, b) // result: 2022 Iron Man2

	// call "foo" and ignore the 2nd output
	b, _ = foo(2021, "go101")
	fmt.Printf("%d %s\n", a, b) // result: 2022 go1012
}

// define a function with 2 arguments for input and output
func foo(a int, b string) (string, int) {
	return b + "2", a + 2
}
