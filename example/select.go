package main

import "fmt"

func main() {
	ch4 := make(chan int, 1)
	ch4 <- 1
	for i := 0; i < 4; i++ {
		select {
		case e, ok := <-ch4:
			if !ok {
				fmt.Println("End.")
				return
			}
			fmt.Println(e)

		default:
			fmt.Println("No Data!")
			close(ch4)
		}
	}
}