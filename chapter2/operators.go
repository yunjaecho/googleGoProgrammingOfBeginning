package chapter2

import "fmt"

func init() {
	fmt.Println("=== operators ===")

	carLotA := 10
	carLotB := 20
	//carLotC := 25

	carLotTotal := carLotA + carLotB
	fmt.Println("ccarLotTotal", carLotTotal)

	carLotTotal *= 10
	fmt.Println("ccarLotTotal", carLotTotal)

	carLotTotal = carLotB / carLotA
	fmt.Println("ccarLotTotal", carLotTotal)
}
