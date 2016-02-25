package main

import (
	"fmt"
	"errors"
)

type PersonBase struct {
	Name string
	Age  int
}
type Boy struct {
	Person PersonBase
	Sex    string
}

func Add(addA int, addB int) (int, error){
	if addA==0 {
		return 0,errors.New("第一个参数不能为0")
	}
	result := addA + addB
	return result,nil
}
func test(slice []string)(){
	slice[2]="y"
	return
}
func main() {
	s := []string{"1","2","3"}
	s1:= s;
	for i,v:=range s {
		fmt.Println(i, v)
	}
	s1[2]="w";
	for i,v:=range s {
		fmt.Println(i, v)
	}
	test(s1)
	for i,v:=range s {
		fmt.Println(i, v)
	}
}

