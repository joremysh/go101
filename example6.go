package main

import (
	"fmt"
)

func example6() {
	str := "Iron Man" // define a string
	var ptr *string   // define a pointer which points to a string
	ptr = &str        // point "ptr" to "str"

	fmt.Printf("%p \n", ptr)  // result like: 0xc000010240
	fmt.Printf("%s \n", *ptr) // result: Iron Man

	stephen := Hero{name: "Dr.Strange"}
	fmt.Println(stephen) // {Dr.Strange 0 0}

	tony := Hero{"IronMan", 30, 666}
	fmt.Println("Before change:", tony) // Before change: {IronMan 30 666}

	changeHero(tony)
	fmt.Println("After change:", tony) // After change: {IronMan 30 666}

	changeHeroPt(&tony)
	fmt.Println("After change:", tony) // After change: {Dr.Strange 55 9999}
}

func changeHero(h Hero) {
	h.name = "Dr.Strange"
	h.age = 55
	h.power = 9999
}

func changeHeroPt(h *Hero) {
	h.name = "Dr.Strange"
	h.age = 55
	h.power = 9999
}

type Hero struct {
	name       string
	age, power int
}
