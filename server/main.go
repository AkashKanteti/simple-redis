package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Hello World")
	})

	err := http.ListenAndServe("localhost:6379", nil)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
