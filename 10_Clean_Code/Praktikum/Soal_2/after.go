package main

type Vehicle struct {
	Tire       int
	speedPerHour int
}

type Cars struct {
	Vehicle
}

func (car *Cars) moving() {
	car.speedUp(10)
}

func (car *Cars) speedUp(newSpeed int) {
	car.speedPerHour = car.speedPerHour + newSpeed
}

func main() {
	fastCar := Cars{}
	fastCar.moving()
	fastCar.moving()
	fastCar.moving()

	slowCar := Cars{}
	slowCar.moving()
}