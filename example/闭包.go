package main

import "fmt"

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		//go func() {
		//	fmt.Println(v)
		//	done <- true
		//}()
		go func(v string) {
			fmt.Println(v)
			done <- true
		}(v)

	}

	// 在退出前等待所有Go程完成
	for _ = range values {
		<-done
	}
}
