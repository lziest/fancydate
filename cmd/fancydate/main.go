package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/lziest/fancydate"
)

var (
	dateStr string
	csv     bool
)

const (
	shortForm = "2006-Jan-02"
)

func main() {
	flag.StringVar(&dateStr, "date", "", "the special date, formed like 2006-Jan-02")
	flag.BoolVar(&csv, "csv", false, "output in a csv file that can be imported to Google Calendar")
	flag.Parse()
	t, err := time.Parse(shortForm, dateStr)
	if err != nil {
		fmt.Println(err)
	}

	var dates []time.Time
	dates = append(dates, fancydate.AfterManyZeros(t)...)
	dates = append(dates, fancydate.AfterLuckyNumbers(t)...)

	if csv {
		fmt.Println("Subject,Start Date,Description,All Day Event")
	}
	for _, d := range dates {
		if d.After(time.Now()) {
			fmt.Printf("Special Day, ")
			fmt.Printf("%d/%d/%d, ", d.Month(), d.Day(), d.Year())
			fmt.Printf("It is %d days from %v,",
				int(d.Sub(t).Hours()/24), t)
			if csv {
				fmt.Printf("%v", true)
			}
			fmt.Println()
		}
	}
}
