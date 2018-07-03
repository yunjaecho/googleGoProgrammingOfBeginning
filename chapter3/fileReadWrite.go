package chapter3

import (
	"io/ioutil"
	"fmt"
)

func init() {
	fmt.Println("=== File Read and Write ===")
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	//write
	stringContent := string(b) + "\n new stuff"
	fmt.Println("write to file : ", stringContent)
	error := ioutil.WriteFile("output.txt", []byte(stringContent), 0644)
	if error != nil {
		panic(error)
	}
}
