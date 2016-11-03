package main

import "fmt"

func main(){
	var l int64
	l = (int64)(len([]byte("Golang中文社区——这里是多余的")))
	fmt.Print(l)
}

//utf-8
//42 =  6 + 12*3