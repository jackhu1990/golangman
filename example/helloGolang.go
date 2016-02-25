package main

import "fmt"

type PersonBase struct {
	Name string
	Age  int
}
type Boy struct {
	Person PersonBase
	Sex    string
}

func main() {
	fmt.Println("Hello Golang")
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

}
func Add(addA int, addB int) (result int){
	result = addA + addB
}
