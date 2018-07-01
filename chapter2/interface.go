package chapter2

import "fmt"

type person struct {
	firstname string
	lastname string
}

type behaviors interface {
	talk() string
	walk() int
	run() int
}

type policeOfficer struct {
	badgeNumber int
	percinct string
}

func init() {
	fmt.Println("=== Interface ===")
	p := new(person)
	fmt.Println(p.talk())
	fmt.Println(p.walk())

	o := new(policeOfficer)
	fmt.Println(o.talk())
	fmt.Println(o.walk())
	o.badgeNumber = 1000
	fmt.Println(o.badgeNumber)
	fmt.Println(o.run())

}

// person implementation
func (p* person) talk() string {
	return "hi there!"
}


func (p* person) walk() int {
	return 10
}

func (p* person) run() int {
	return 10
}

// officer implementation
func (o* policeOfficer) talk() string {
	return "Have a good day"
}

func (o* policeOfficer) walk() int {
	return 20
}

func (o* policeOfficer) run() int {
	return 50
}
