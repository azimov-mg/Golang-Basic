package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	a, b := "104", 35
	intToStr := strconv.Itoa(b)
	strToInt, err := strconv.Atoi(a)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%q\n", intToStr)
	fmt.Println(strToInt)
}
