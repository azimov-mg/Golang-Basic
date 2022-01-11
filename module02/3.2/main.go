package main

import "fmt"

func main() {
	Weekdays, Weekend := []string{"Mon", "Tue", "Wen", "Thr", "Fri"}, []string{"Sat", "San"}
	DaysOfTheWeek := make([]string, 7, 7)
	DaysOfTheWeek = append(Weekdays, Weekend...)

	fmt.Println(DaysOfTheWeek)
	fmt.Println(Weekdays)
	fmt.Println(Weekend)

}
