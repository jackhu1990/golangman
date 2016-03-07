package main

import "fmt"

type Handle int64
func (h Handle) Show1(i int64) int64{
	h = 1
	return i + int64(h)
}
func (h *Handle) Show2(i int64) int64{
	*h = 2
	return i + int64(*h)
}

func main() {
	var hand Handle = 0
	fmt.Println(hand.Show1(100))
	fmt.Println(hand)

	fmt.Println(hand.Show2(100))
	fmt.Println(hand)
}
