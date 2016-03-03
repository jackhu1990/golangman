package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	secret string
}
type Girl struct {
	Person Person
	Sexy    bool
}

func main() {
	defer func() {
		fmt.Println("程序结束了")
	}()
	person := Person{"胡彦春", 18, "秘密是小写的,不会出现在json中,很有意思的特性"}
	fmt.Println(person.secret)
	gilr := Girl{person, true}
	b, err := json.Marshal(gilr)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(b))

}
