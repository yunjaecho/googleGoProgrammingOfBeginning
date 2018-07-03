package chapter4

import (
	"fmt"
	"os"
	"time"
)

func init() {
	fmt.Println("=== Go Routine ===")
	cars := fillCars()

	go showCars(cars, "first goroutine")
	go showCars(cars, "second goroutine")
	go showCars(cars, "third goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(2 * time.Second)

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

}

type Cars map[string]int

func fillCars() Cars {
	cars := make(map[string]int)
	cars["Jeep"] = 22
	cars["Buick"] = 11
	cars["Ford"] = 15
	cars["Chevy"] = 9
	cars["Nissan"] = 29
	return cars
}

func showCars(c Cars, m string) {
	for key, value := range c {
		fmt.Fprintf(os.Stdout, "[%[1]v] %[2]v = %[3]v %[4]v\n", m, value, key, c[key])
	}
}
