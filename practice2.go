package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"github.com/xuri/excelize"
)

func practice2() {
	tNow := time.Now()
	tStart := tNow.AddDate(0, 0, -30)
	dtStart := datetime.New(&tStart)
	dtEnd := datetime.New(&tNow)
	params := &chart.Params{
		Symbol:   "TSLA",
		Interval: datetime.OneDay,
		Start:    dtStart,
		End:      dtEnd,
	}
	iter := chart.Get(params)

	dataset := make([]map[string]interface{}, 0)
	for iter.Next() {
		fmt.Printf("%s %+v\n", time.Unix(int64(iter.Bar().Timestamp), 0).Format("2006-01-02"), iter.Bar())
		m := make(map[string]interface{})
		// save the time and price values from Bar here

		dataset = append(dataset, m)
	}

	filePath := "Book1.xlsx"
	os.Remove(filePath)
	f := excelize.NewFile()

	st := "Sheet1"
	f.SetCellValue(st, "A1", "time")
	f.SetCellValue(st, "B1", "high")
	f.SetCellValue(st, "C1", "low")
	f.SetCellValue(st, "D1", "close")
	for i, data := range dataset {
		f.SetCellValue(st, fmt.Sprintf("A%v", i+2), data["time"])
		// set the price values here

	}

	if err := f.AddChart("Sheet1", "F1", fmt.Sprintf(`{
	"type":"line",
	"series":[
	   {"name":"Sheet1!$B$1","categories":"Sheet1!$A$2:$A$%v",
	       "values":"Sheet1!$B$2:$B$%v"}
	   ],
	   "title":{"name":"TSLA stock price"}}`, len(dataset)+1, len(dataset)+1)); err != nil {
		log.Fatal(err.Error())
	}

	if err := f.SaveAs(filePath); err != nil {
		log.Fatal(err.Error())
	}

	if err := iter.Err(); err != nil {
		log.Fatal(err.Error())
	}
}

// func main() {
// 	practice2()
// }
