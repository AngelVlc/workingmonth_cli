package main

import (
	"fmt"
	"github.com/AngelVlc/workingmonth"
	"log"
	"time"
)

func main() {
	var discountDays int

	fmt.Print("Holidays: ")
	if _, err := fmt.Scan(&discountDays); err != nil {
		log.Fatal(err)
	}

	now := time.Now()

	wm := workingmonth.WorkingMonth{
		Year:  now.Year(),
		Month: int(now.Month()),
	}

	hoursUntilToday := wm.WorkingHoursUntilToday() - (discountDays * 8)
	monthHours := wm.WorkingHours() - (discountDays * 8)
	fmt.Printf("Working hours until today: %v\n", hoursUntilToday)
	fmt.Printf("Working hours until month end:  %v\n", monthHours)
	fmt.Printf("(Bamboo adds future PTO days)\n")
}
