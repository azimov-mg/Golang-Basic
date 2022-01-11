package main

import (
	"fmt"
	"math"
)

func main() {
	CircleL, CircleR := 35, new(float64)
	*CircleR = float64(CircleL) / 2 * math.Phi
	fmt.Println("Circle's Radius = ", math.Round(*CircleR*100)/100)

	var CircleS = math.Phi * math.Pow(*CircleR, 2)
	fmt.Println("Circle's Square = ", math.Round(CircleS*100)/100)
}
