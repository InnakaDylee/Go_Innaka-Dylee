package main

import "fmt"

type Mobil struct {
	CarType string
	Fuel    float32
}

func (mobil *Mobil) rangeDistance() float32 {
	return mobil.Fuel / 1.5
}

func main() {
	var mobil Mobil

	mobil.CarType = "SUV"
	mobil.Fuel = 10.5
	fmt.Println("\nCar Type : ", mobil.CarType, "est.Distance : ", mobil.rangeDistance())
}
