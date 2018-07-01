package chapter2

import "fmt"

func init() {
	fmt.Println("=== Map ===")
	carDealerInventory := make(map[string] int)
	carDealerInventory["A1 Auto"] = 76
	fmt.Println("A1 Auto Inv = ", carDealerInventory["A1 Auto"])
	fmt.Println("len of carDealerInventory = " , len(carDealerInventory))

	carDealerInventory["Southest Auto Mall"] = 112
	fmt.Println("len of carDealerInventory = " , len(carDealerInventory))
	delete(carDealerInventory, "A1 Auto")
	fmt.Println("len of carDealerInventory = " , len(carDealerInventory))
	fmt.Println("A1 Auto Inv = ", carDealerInventory["A1 Auto"])
	fmt.Println("Southest Auto Mall Inv = ", carDealerInventory["Southest Auto Mall"])


}
