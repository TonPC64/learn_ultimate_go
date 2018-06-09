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
	myvalue := MyType{Value1: 1, Value2: "ss"}
	pointer := unsafe.Pointer(&myvalue)
	fmt.Printf("Addr: %v value1: %d value2: %s\n", pointer, myvalue.Value1, myvalue.Value2)

	changeMyValue(myvalue)
	fmt.Printf("Addr: %v value1: %d value2: %s\n", pointer, myvalue.Value1, myvalue.Value2)
}

func changeMyValue(myvalue MyType) {
	myvalue.Value1 = 20
	myvalue.Value2 = "ton"

	pointer := unsafe.Pointer(&myvalue)
	fmt.Printf("Addr: %v value1: %d value2: %s\n", pointer, myvalue.Value1, myvalue.Value2)

}
