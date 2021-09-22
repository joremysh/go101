package main

import (
	"fmt"
	"log"

	"github.com/piquette/finance-go/quote"
)

func practice1() {
	q, err := quote.Get("TSLA")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("------- %v -------\n", q.ShortName)
	fmt.Printf("Current Price: $%v\n", q.Ask)
	fmt.Printf("52wk High: $%v\n", q.FiftyTwoWeekHigh)
	fmt.Printf("52wk Low: $%v\n", q.FiftyTwoWeekLow)
}

//func main() {
//	practice1()
//}
