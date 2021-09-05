package main

import (
	"fmt"
)

func example3() {
	// define an slice with 5 elements
	arr := []string{"Iron Man", "Dr.Strange", "Thor", "Captain America", "Hulk"}
	fmt.Println(arr[0]) // print the 1st element: Iron Man
	fmt.Println(arr[1]) // print the 2nd element: Dr.Strange

	var arr1 []string               // define an slice with no elements
	arr1 = append(arr1, arr[3])     // add an element into arr1
	fmt.Println(len(arr1), arr1[0]) // result: 1 Captain America

	arr2 := make([]string, 0)       // define an slice with length 0
	arr2 = append(arr1, arr[4])     // add an element into arr2
	fmt.Println(len(arr2), arr2[1]) // result: 2 Hulk

	ageOfHeroes := make(map[string]int)    // define a map with string key and int value
	ageOfHeroes["Dr.Strange"] = 45         // add a pair of key and value to the map
	fmt.Println(ageOfHeroes["Dr.Strange"]) // result: 45

	age, ok := ageOfHeroes["Dr.Strange"]
	if ok {
		fmt.Println("Dr.Strange:", age)
	} // result: Dr.Strange: 45

	age, ok = ageOfHeroes["Thor"]
	fmt.Println(age, ok) // result: 0 false
}

func main() {
	example3()
}
