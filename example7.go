package main

import (
	"fmt"
)

// func main() {
// 	example7()
// }

func (p Point) AddX(val float64) {
	p.X = p.X + val
}

func (p *Point) AddXP(val float64) {
	p.X = p.X + val
}

type Point struct {
	X, Y float64
}

func example7() {
	p := &Point{1.1, 2.5}
	p.AddX(1.2)
	fmt.Println(p) // &{1.1 2.5}

	p.AddXP(1.2)
	fmt.Println(p) // &{2.3 2.5}

	pt := &Point{1.1, 2.5}
	describe(1)       // (1, int)
	describe("go101") // (go101, string)
	describe(pt)      // (&{1.1 2.5}, *main.Point)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
