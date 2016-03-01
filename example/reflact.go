package main

import (
	"fmt"
	"reflect"
)

func main() {
	x:=3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))
}