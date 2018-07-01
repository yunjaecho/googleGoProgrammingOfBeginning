package chapter2

import (
	"fmt"
	"strconv"
)

/**
Type Cast : strconv.Atoi
 */
func init() {
	fmt.Println("=== Type and Cast ===")

	var mystring string = "1"
	var x = 10
	var f float32 = 2.0
	fmt.Println(mystring , x, f)
	number, err := strconv.Atoi(mystring)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(number)
	x = number
	fmt.Println(x)
	fmt.Printf("%T\n", f)
	fmt.Printf("%v\n", f)
}