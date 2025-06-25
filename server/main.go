package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})

	err := http.ListenAndServe(":6379", nil)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
