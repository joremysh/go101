package main

import "fmt"

func example5() {
	arr := []string{"Iron Man", "Dr.Strange", "Thor", "Captain America", "Hulk"}
	for i := 0; i < 5; i++ {
		fmt.Printf("%s ", arr[i])
	} // result: Iron Man Dr.Strange Thor Captain America Hulk

	for i, s := range arr {
		fmt.Printf("%d %s ", i, s)
	} // 0 Iron Man 1 Dr.Strange 2 Thor 3 Captain America 4 Hulk

	i := 0
	for i < 10 {
		fmt.Printf("%d", i)
		i++
	} // result: 0123456789

	for { // will run continuously until reach break
		break
	}
}
