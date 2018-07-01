package chapter3

import "fmt"

type vehicle struct {
	mpg int
	numberOfDoors int
}

type car struct {
	vehicle
	color Color
}

type truck struct {
	vehicle
	color Color
}

func (t truck) getVN() string {
	return "truck VN"
}

func (c car) getVN() string {
	return "car VN"
}


func newCar() *car {
	result := car{}
	result.mpg = 20
	result.numberOfDoors = 4
	result.color = black
	return &result
}

func(v *vehicle) getMpg() {
	println("This wehicle get,", v.mpg, "mpg")
}

//custom types
type Color string
const (
	blue Color = "blue"
	red Color = "red"
	black Color = "black"
)


func init() {
	fmt.Println("=== Composition ====")
	aCar := car{}
	fmt.Println(aCar)
	aCar.mpg = 25
	aCar.numberOfDoors = 2
	fmt.Println(aCar)
	aCar.getMpg()
	bCar := car{vehicle{19, 4}, red}
	bCar.color = red
	fmt.Println(bCar)
	bCar.getMpg()

	cCar := car{}
	cCar.color = black
	cCar.mpg = 34
	fmt.Println(cCar)
	cCar.getMpg()

	defaultCar := newCar{}
	fmt.Println("defaultCar ", defaultCar)

}


