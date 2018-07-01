package chapter2

import "fmt"


func createDealerFullName(s1 string, s2 string) string {
	return s1 + "of" + s2
}

func calculateInvUtil(remaining float32, start float32, dealer dealerShip) (percentSold float32, dealerName string) {
	dealerName = dealer.name + "sold"
	percentSold = 1 - remaining / start
	return
}



func init() {
	fmt.Println("=== Functions ===")
	var d1 = dealerShip{"A1 Auto", "Los Angeles"}
	var dealerName = createDealerFullName(d1.name, d1.city)
	fmt.Println(dealerName)

	var sold, name = calculateInvUtil(100, 175, d1)
	fmt.Println(name, sold)
}


