package main

import (
	"fmt"
	"log"
	"time"

	"github.com/AngelVlc/workingmonth"
)

func main() {
	var discountDays float32

	fmt.Print("Holidays: ")
	if _, err := fmt.Scan(&discountDays); err != nil {
		log.Fatal(err)
	}

	now := time.Now()

	wm := workingmonth.WorkingMonth{
		Year:  now.Year(),
		Month: int(now.Month()),
	}

	hoursUntilToday := float32(wm.WorkingHoursUntilToday()) - (discountDays * 8)
	monthHours := float32(wm.WorkingHours()) - (discountDays * 8)
	fmt.Printf("Working hours until today: %v\n", hoursUntilToday)
	fmt.Printf("Working hours until month end:  %v\n", monthHours)
	fmt.Printf("IMPORTANT: Bamboo adds future PTO days\n")
}
