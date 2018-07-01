package chapter2

import "fmt"

func init() {
	fmt.Println("=== Slices ===")
	var carTypes[3] string
	carTypes[0] = "Toyota"
	carTypes[1] = "Ford"
	carTypes[2] = "Nissan"
	fmt.Println(carTypes[1])

	carTypes2 := [3]string{"Toyota", "Ford", "Nissan"}
	fmt.Println(carTypes2[0])

	carTypesSlice := []string{"Toyota", "Ford", "Nissan"}
	fmt.Println(carTypesSlice[2])
	carTypesSlice = append(carTypesSlice, "Telsa")
	carTypesSlice = append(carTypesSlice, "Append")
	fmt.Println("carTypesSlice len = ", len(carTypesSlice))

	carTypesSliceMake := make([]string, 3)
	fmt.Println("carTypesSliceMake len = ", len(carTypesSliceMake))
	carTypesSliceMake[0] = "Toyota"
	carTypesSliceMake[1] = "Ford"
	carTypesSliceMake[2] = "Nissan"

	// carTypesSlice 배열에 2번째 부터 3(4-1)번째 요소 가져오기기
	carTypesSlice2 := carTypesSlice[2:4]
	fmt.Println("carTypesSlice2 len ", len(carTypesSlice2))
	fmt.Println("slice[2:4] = ", carTypesSlice2)
}

