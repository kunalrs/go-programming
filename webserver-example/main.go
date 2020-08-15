package main

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Request in..")
	fmt.Fprint(res, "<h1>Hello, my name is Kunal Sagar</h1>")
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Listing on 4000...")
	http.ListenAndServe("localhost:4000", nil)
}
