package chapter3


type Truck struct {
	numberOfDoors int
	bedSize string
	weightByTon weight
}

type weight string

const (
	oneTon weight ="One Ton"
	twoTon weight = "Two Ton"
)

func (t *Truck) getDoors()(int) {
	return t.numberOfDoors
}