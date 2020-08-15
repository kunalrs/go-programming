package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello")

var spanish bool

var greeting string

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use spanish language")
	flag.BoolVar(&spanish, "s", false, "Use spanish language")

	flag.StringVar(&greeting, "greet", "Hello", "Set greeting text")
	flag.StringVar(&greeting, "g", "Hello", "Set greeting text")
}

func main() {
	flag.Parse()

	if spanish {
		fmt.Printf("Holla %s!\n", *name)
	} else {
		fmt.Printf("%s %s!\n", greeting, *name)
	}
}
