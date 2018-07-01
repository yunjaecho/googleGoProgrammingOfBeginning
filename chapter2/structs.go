package chapter2

import "fmt"

type dealerShip struct {
	name string
	city string
}

func init() {
	fmt.Println("=== Struct ===")
	d1 := dealerShip{name: "A1 Auto" , city: "Los Angeles"}
	d1.city = "los angeles"
	fmt.Println("city = " + d1.city)

	var d2 dealerShip
	d2 = dealerShip{name: "A2 Auto" , city: "Houston"}
	fmt.Println(d2)

	d3 := dealerShip{}
	fmt.Println("d3.name = " + d3.name)
}
