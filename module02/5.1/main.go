package main

import "fmt"

func contains(a []string, x string) bool {
	for _, key := range a {
		if key == x {
			return true
		}
	}
	return false
}

func getMax(val ...int) int {
	var max int
	for i := 0; i < 5; i++ {
		if val[i] > val[i+1] {
			max = val[i]
		}
	}
	return max
}

func main() {
	x, a := "key", []string{"success", "way", "map", "key", "go"}

	fmt.Println(contains(a, x))
	fmt.Println(getMax(5, 2, 3, 6, 10, 1))
}
