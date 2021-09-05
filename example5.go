package main

import "fmt"

func main() {
	example5()
}

func example5() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	} // result: 0123456789

	i := 0
	for i < 10 {
		fmt.Printf("%d", i)
		i++
	} // result: 0123456789

	for { // will run continuously until reach break
		break
	}
}
