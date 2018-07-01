package chapter2

import "fmt"

func init() {
	fmt.Println("=== Function Variadic ===")
	var d1 = "A1 Auto"
	var d2 = "Discount AUto"
	var d3 = "Riverside Automart"
	var dealers = []string {d1, d2, d3}
	printDealers(dealers...)
}

/**
  input parameter : dealers 가변 변수
 */
func printDealers(dealers ... string) {
	// 배열요소는 첫번째 인덱스, 두번째 배열요소 값
	for _, dealerName := range dealers {
		fmt.Println(dealerName)
	}
}