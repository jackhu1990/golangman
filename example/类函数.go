package main

import "fmt"

type Cat struct {
	Name    string
	Age     int32
	Address string
}

func (cat *Cat) Grow() {
	cat.Age++
}
func main() {
	myCat := Cat{"Little C", 2, "In the house"}
	myCat.Grow()
	fmt.Printf("%v", myCat)
}
