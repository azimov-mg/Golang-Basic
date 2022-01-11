package main

import "fmt"

func main() {
	var A *int
	B := 35
	A = &B
	fmt.Println(*A)

	*A = 49
	fmt.Println(B)
}
