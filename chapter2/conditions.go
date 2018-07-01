package chapter2

import "fmt"

func init()  {
	fmt.Println("=== Conditions ===")
	carLotA := 9
	carLotB := 29

	if carLotB > carLotA {
		fmt.Println("car lot B is greater than A")
	} else if carLotB == carLotA {
		fmt.Println("car lot A is same B")
	} else {
		fmt.Println("car lot A is greater than B")
	}

	switch carLotA {
	case 15:
		fmt.Println("case value is 15")
	case 9, 11, 12:
		fmt.Println("case value is 9 or 11 or 12")
	default:
		fmt.Println("default")
	}
}
