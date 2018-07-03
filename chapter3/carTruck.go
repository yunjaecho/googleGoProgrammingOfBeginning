package chapter3

import "fmt"

func init() {
	fmt.Println("=== Car and Truck")
	c := Car{4, 6}
	fmt.Println(c)
	fmt.Println(c.getDoors())

	t := Truck{2, "full", oneTon}
	fmt.Println(t)
	fmt.Println(t.getDoors())
}