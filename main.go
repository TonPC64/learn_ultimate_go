package main

import (
	"fmt"
	"unsafe"
)

type MyType struct {
	Value1 int
	Value2 string
}

func main() {
	// myvalue := new(MyType)
	// myvalue := &MyType{Value1: 1, Value2: "ss"}

	myMap := make(map[string]string)
	myMap["ton"] = "ton"

	pointer := unsafe.Pointer(&myMap)
	fmt.Printf("Addr: %v value: %s\n", pointer, myMap["ton"])

	changeMyMap(&myMap)
	fmt.Printf("Addr: %v value: %s\n", pointer, myMap["ton"])
}

func changeMyMap(myMap *map[string]string) {
	(*myMap)["ton"] = "tonton"

	pointer := unsafe.Pointer(myMap)
	fmt.Printf("Addr: %v value: %s\n", pointer, (*myMap)["ton"])

}

func changeMyValue(myvalue *MyType) {
	myvalue.Value1 = 20
	myvalue.Value2 = "ton"

	pointer := unsafe.Pointer(myvalue)
	fmt.Printf("Addr: %v value1: %d value2: %s\n", pointer, myvalue.Value1, myvalue.Value2)

}
