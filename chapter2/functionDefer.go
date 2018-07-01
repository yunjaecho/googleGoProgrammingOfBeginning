package chapter2

import (
	"os"
	"fmt"
)

func init() {
	fmt.Println("=== Function Defer ====")
	var file = createFile("dealers.txt")
	// defer : init 함수가 종료되기 직전에 지연 시행 시킴
	defer closeFile(file)
	writeToFile(file, "A1 Auto")
}

func createFile(path string) *os.File {
	fmt.Println("create file")
	file, err := os.Create(path)
	if err != nil {
		// Stop (exit)
		panic(err)
	}
	return file
}


func writeToFile(file *os.File, dealerName string) {
	fmt.Println("writing to file")
	fmt.Fprintln(file, dealerName)
}

func closeFile(file *os.File) {
	fmt.Println("close File")
	file.Close()
}