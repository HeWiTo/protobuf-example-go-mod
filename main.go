package main

import (
	"fmt"

	simplepb "github.com/HeWiTo/protobuf-example-go-mod/src/simple"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "Simple Test",
		SampleList: []int32{1, 2, 4, 8},
	}

	fmt.Println(sm)
}
