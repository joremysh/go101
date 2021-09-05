package main

import (
	"fmt"
)

func example1() {
	anInteger := 1                                   // define an int type
	aString := "hello"                               // string
	aBool := true                                    // bool
	aBool = false                                    // set it to false
	aFloat64 := 1.2                                  // float64
	fmt.Println(anInteger, aString, aBool, aFloat64) // result: 1 hello false 1.2

	anInteger = int(aFloat64)        // convert float64 to int
	aFloat64 = float64(anInteger)    // convert int to float64
	fmt.Println(anInteger, aFloat64) // result: 1 1
}
