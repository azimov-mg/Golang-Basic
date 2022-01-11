package main

import "fmt"

func main() {
	daysToWeekend := []string{"Mon", "Tue", "Wen", "Thr", "Fri", "Sat", "San"}
	Weekdays := make([]string, 5, 5)
	copy(Weekdays, daysToWeekend)
	daysToWeekend = daysToWeekend[len(daysToWeekend)-2:]

	fmt.Println("Weekdays:", Weekdays)
	fmt.Println("Weekend:", daysToWeekend)
}
