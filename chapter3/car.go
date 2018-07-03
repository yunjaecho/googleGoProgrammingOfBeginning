package chapter3

type Car struct {
	numberOfDoors int
	cylinders int
}


func (t *Car) getDoors()(int) {
	return t.numberOfDoors
}