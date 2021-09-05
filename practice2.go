package main

import (
	"log"

	"github.com/xuri/excelize"
)

func practice2() {
	f := excelize.NewFile()

	if err := f.SaveAs("Book1.xlsx"); err != nil {
		log.Fatal(err.Error())
	}
}
