package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	if resp, err := http.Get("http://google.com"); err == nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	} else {
		fmt.Println(err)
	}
}
