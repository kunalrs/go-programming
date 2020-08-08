package main

import (
	"fmt"
	"time"
)

func printCount(ch chan int) {
	num := 0
	for num >= 0 {
		num = <-ch
		fmt.Print(num, " ")
	}
}

func main() {
	ch := make(chan int)

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	go printCount(ch)
	fmt.Println("Called Print and waiting to check...")
	time.Sleep(3 * time.Second)
	for _, v := range a {
		ch <- v
	}
	time.Sleep(1 * time.Millisecond)
	fmt.Println("\nMain ended")
}
