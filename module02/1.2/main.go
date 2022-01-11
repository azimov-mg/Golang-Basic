package main

func main() {
	type AmericanVelocity int
	type EuropeanVelocity int

	msToKmh, msToMph := 120.4, 130
	var kmh = EuropeanVelocity(msToKmh * 3.6)
	var Mph = AmericanVelocity(float64(msToMph) / 1.609)

	println(kmh, Mph)
}
