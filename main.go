package main

import (
	"fmt"
	"net/http"

	"github.com/el-zacharoo/go-101/handler"
	"github.com/el-zacharoo/go-101/print"
)

func main() {
	http.HandleFunc("/person", handler.PersonHandler)
	http.HandleFunc("/people", handler.PeopleHandler)
	http.HandleFunc("/print", print.Message)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Print(err)
	}
}
