package main

import (
	"fmt"
	"time"
)

func main() {
	month := "09"
	yymm := "2305"

	yymmAsDate, err := time.Parse("0601", yymm)

	if err != nil {
		panic(err)
	}

	monthAsDate, err := time.Parse("01", month)

	monthAsDate = monthAsDate.AddDate(0, 3, 0)

	if err != nil {
		print("ERROR: ", err)
	}

	yymmAsDate = yymmAsDate.AddDate(-1, 3, 5)

	fmt.Printf("Parsed month: %d\n", monthAsDate.Month())
	fmt.Printf("Parsed yymm: %02d\n", yymmAsDate.Month())

	formatString := `Day: %02d
Month: %02d
Year: %02d`

	fmt.Printf(formatString, yymmAsDate.Day(), yymmAsDate.Month(), yymmAsDate.Year())
}
