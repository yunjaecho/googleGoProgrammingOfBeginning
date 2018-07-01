package chapter2

import "fmt"

var shippingDays = 30
var shippingDayPtr = new(int)


func init() {
	fmt.Println("=== Point ===")
	shippingDaysAjustments(shippingDays)
	fmt.Println("after shippingDaysAjustments call ", shippingDays)

	shippingDaysAjustmentsPtr(&shippingDays)
	fmt.Println("after shippingDaysAjustmentsPtr call ", shippingDays)
	fmt.Println("after shippingDaysAjustmentsPtr call address ", &shippingDays)

	shipper := shipper{}
	shipper.id = 400
	shipper.name = "Pacific Shippers"
	shipperUpdates(&shipper)
	fmt.Println("shipper.id after shiiperUpdates call ", shipper.id)
	fmt.Println("shipper.name after shiiperUpdates call ", shipper.name)

	*shippingDayPtr = 55
	shippingDaysAjustmentsPtr(shippingDayPtr)
	fmt.Println("shippingDayPtr after shippingDaysAjustmentsPtr call ", shippingDayPtr)
}


type shipper struct {
	name string
	id int
}

func shippingDaysAjustments(shippingDays int) {
	shippingDays += 10
}

func shippingDaysAjustmentsPtr(shippingDays* int) {
	*shippingDays += 10
}

func shipperUpdates(s* shipper) {
	s.id += 10
	s.name += " Inc"
}