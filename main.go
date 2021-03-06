package main

import (
	"fmt"
	"log"
	"time"

	"github.com/AngelVlc/workingmonth"
)

func main() {
	var futureHolidays float32
	fmt.Println("Future holidays:")
	if _, err := fmt.Scan(&futureHolidays); err != nil {
		log.Fatal(err)
	}
	
	discountDays := futureHolidays

	now := time.Now()

	wm := workingmonth.WorkingMonth{
		Year:  now.Year(),
		Month: int(now.Month()),
	}

	hoursUntilToday := float32(wm.WorkingHoursUntilToday()) + (discountDays * 8)
	monthHours := float32(wm.WorkingHours()) // - (discountDays * 8)
	fmt.Printf("Working hours until today: %v\n", hoursUntilToday)
	fmt.Printf("Working hours until month end:  %v\n", monthHours)
}
