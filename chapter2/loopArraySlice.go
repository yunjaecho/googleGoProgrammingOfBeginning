package chapter2

import "fmt"

func init() {
	fmt.Println("=== Loop Array Slice")
	var carTypes[3]string
	carTypes[0] = "Toyota"
	carTypes[1] = "Ford"
	carTypes[2] = "Nissan"

	i := 0
	for i < len(carTypes) {
		fmt.Println(carTypes[i])
		i++
	}

	fmt.Println("\n for loop verbose")
	for j :=0; j < len(carTypes); j++ {
		fmt.Println(carTypes[j])
	}

	fmt.Println("\n range loop")
	for k, value := range carTypes {
		fmt.Println("k= ", k, "value=", value)
	}

	carDealerInventory := make(map[string]int)
	carDealerInventory["A1 Auto"] = 76
	carDealerInventory["Southwest Auto Mall"] = 112

	carDealerInventory2 := make(map[int]string)
	carDealerInventory2[76] = "A1 Auto"
	carDealerInventory2[112] = "Southwest Auto Mall"

	fmt.Println("\n first dictionary loop")
	for m := 0; m < len(carDealerInventory2); m++ {
		fmt.Println(carDealerInventory2[m])
	}
	fmt.Println(carDealerInventory2[76])

	for key, value := range carDealerInventory {
		fmt.Println("key = ", key, " & value = ", value)
	}

 }