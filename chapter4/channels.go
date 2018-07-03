package chapter4

import (
	"strconv"
	"fmt"
	"time"
)

var frameId = 0
var frameName = ""

func init() {
	fmt.Println("=== Channel ... ===")

	framesToCreate := 4
	frameInfoChan := make(chan string)

	for i:=0; i < framesToCreate; i++ {
		go assembleFrame(frameInfoChan)
		go addBody(frameInfoChan)
		go addInterior(frameInfoChan)
		time.Sleep(time.Microsecond * 1000)
	}
}

func assembleFrame(frameInfoChan chan string) {
	frameId++
	frameName = "Frame ID" + strconv.Itoa(frameId)
	fmt.Println("Frame assembly complete", frameName, "Proceed to body")
	frameInfoChan <- frameName
	time.Sleep(time.Microsecond * 5)
}

func addBody(frameInfoChan chan string) {
	body := <- frameInfoChan
	fmt.Println("Add body to ", body, "and proceed to interior")
	frameInfoChan <- frameName
	time.Sleep(time.Microsecond * 5)
}

func addInterior(frameInfoChan chan string) {
	interior := <- frameInfoChan
	fmt.Println("Add body to ", interior, "and proceed to paint")
	time.Sleep(time.Microsecond * 5)
}

