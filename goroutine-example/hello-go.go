package main

import (
	"fmt"
	"time"
)

func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	go count()
	time.Sleep(2 * time.Millisecond)
	fmt.Println("Hello Go")
	time.Sleep(5 * time.Millisecond)
}
