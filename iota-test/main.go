package main

import "fmt"

const (
	first = iota + 1
	second
	third
)

func main() {
	fmt.Println(first, second, third)
}
