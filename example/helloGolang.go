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
	fmt.Println("Hello Golang")

	result,err := Add(0,10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	if i := 10; i < 5 {
		fmt.Print("***********")
	}
	var person PersonBase
	var personP *PersonBase
	person.Name = "jackhu"
	person.Age = 27
	personP = new(PersonBase)
	personP.Name = "gaofei"
	personP.Age = 27

	arr :=[...] int {1,2,3,4,5}
	for i,v:= range arr{
		fmt.Println("index:", i, "value:", v)
	}
	m := map[int]string{1:"one", 2:"two"}
	for i,v:= range m{
		fmt.Println("index:", i, "value:", v)
	}
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

