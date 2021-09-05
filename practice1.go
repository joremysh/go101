package main

import (
	"fmt"
	"log"
	"time"

	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"github.com/piquette/finance-go/quote"
)

func practice1() {
	//stockList := []string{"TSLA", "AAPL", "MSFT", "NFLX", "AMZN"}
	q, err := quote.Get("TSLA")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("------- %v -------\n", q.ShortName)
	fmt.Printf("Current Price: $%v\n", q.Ask)
	fmt.Printf("52wk High: $%v\n", q.FiftyTwoWeekHigh)
	fmt.Printf("52wk Low: $%v\n", q.FiftyTwoWeekLow)

	tNow := time.Now()
	tStart := tNow.AddDate(0, -5, 0)
	dtStart := datetime.New(&tStart)
	dtEnd := datetime.New(&tNow)
	params := &chart.Params{
		Symbol:   "TSLA",
		Interval: datetime.OneMonth,
		Start:    dtStart,
		End:      dtEnd,
	}
	iter := chart.Get(params)

	dataset := make([]map[string]interface{}, 0)
	for iter.Next() {
		//fmt.Printf("%s %+v\n", time.Unix(int64(iter.Bar().Timestamp), 0), iter.Bar())
		m := make(map[string]interface{})
		m["time"] = iter.Bar().Timestamp
		m["high"], _ = iter.Bar().High.Float64()
		m["low"], _ = iter.Bar().Low.Float64()
		m["close"], _ = iter.Bar().Close.Float64()
		dataset = append(dataset, m)
	}

	fmt.Println("time\thigh\tlow\tclose")
	for _, m := range dataset {
		t := time.Unix(int64(m["time"].(int)), 0).Format("2006-01-02")
		fmt.Printf("%s\t%.2f\t%.2f\t%.2f\n", t, m["high"], m["low"], m["close"])
	}
}
