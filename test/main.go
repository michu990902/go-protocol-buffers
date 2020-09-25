package main

import (
	"fmt"
	"test/src/simple"
)

func main() {
	doSimple()
}

func doSimple() *simplepb.SimpleMessage{
	sm := simplepb.SimpleMessage{
		Id:       12345,
		IsSimple: true,
		Name: "test",
		SampleList: []int32{1,2,3,4},
	}

	fmt.Println(sm)
	return &sm
}
