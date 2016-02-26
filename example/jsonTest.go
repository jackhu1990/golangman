package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}
type Girl struct {
	Person Person
	Sex    bool
}

func main() {
	defer func() {
		fmt.Println("程序结束了")
	}()
	person := Person{"胡彦春", 18}
	gilr := Girl{person, true}
	b, err := json.Marshal(gilr)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(b))

}
